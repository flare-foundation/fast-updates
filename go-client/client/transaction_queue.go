package client

import (
	"context"
	"fast-updates-client/config"
	"fast-updates-client/logger"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
)

const queueLen = 50

// Default EIP-1559 fee parameters, mirroring flare-system-client. Applied when
// the corresponding config value is unset/<=0.
const (
	defaultBaseFeeCapMultiplier int64 = 4
	defaultTipMultiplier        int64 = 2
	defaultMinPriorityFeeWei    int64 = 100_000_000_000  // 100 Gwei
	defaultMaxPriorityFeeWei    int64 = 5000_000_000_000 // 5000 Gwei
)

type TransactionQueue struct {
	InputChan         chan ComputationRequest
	ErrChan           chan *ErrorRequest
	chainClient       *ethclient.Client
	numInQueue        int
	transactionConfig config.TransactionsConfig
	accounts          []*Account
	toTrash           bool
	CurrentBlockNum   uint64
	Stop              bool
	sync.Mutex
}

type ComputationRequest struct {
	Function    func(*bind.TransactOpts) error
	BeforeBlock uint64
	AfterBlock  uint64
}

type ErrorRequest struct {
	Err     error
	Request ComputationRequest
}

func NewTransactionQueue(client *ethclient.Client, accounts []*Account, txConfig config.TransactionsConfig) *TransactionQueue {
	txQueue := TransactionQueue{
		InputChan:         make(chan ComputationRequest, queueLen),
		ErrChan:           make(chan *ErrorRequest, queueLen),
		chainClient:       client,
		accounts:          accounts,
		transactionConfig: txConfig,
	}

	return &txQueue
}

func (txQueue *TransactionQueue) Run() {
	for i := 0; i < len(txQueue.accounts); i++ {
		go txQueue.QueueExecution(i)
	}
	go txQueue.ErrorHandler()
	go txQueue.CurrentBlockNumUpdater()
}

func (txQueue *TransactionQueue) QueueExecution(i int) {
	for {
		if txQueue.Stop {
			return
		}
		compReq := <-txQueue.InputChan
		txQueue.Lock()
		txQueue.numInQueue++
		txQueue.Unlock()

		if txQueue.toTrash {
			logger.Info("ignoring submission")
			txQueue.ErrChan <- nil
			continue
		}

		if compReq.AfterBlock != 0 {
			txQueue.WaitForBlockWithCheck(compReq.AfterBlock)
		}

		// again check if the transaction should be ignored
		if txQueue.toTrash {
			logger.Info("ignoring submission")
			txQueue.ErrChan <- nil
			continue
		}

		if compReq.BeforeBlock != 0 {
			if txQueue.CurrentBlockNum > compReq.BeforeBlock {
				err := fmt.Errorf("skipping submission, too late")
				txQueue.ErrChan <- &ErrorRequest{err, compReq}
				continue
			}
		}
		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(config.CallTimeoutMillisDefault)*time.Millisecond)
		nonce, err := txQueue.chainClient.NonceAt(ctx, txQueue.accounts[i].Address, nil)
		cancelFunc()
		if err != nil {
			txQueue.ErrChan <- &ErrorRequest{err, compReq}
			continue
		}
		opts, err := bind.NewKeyedTransactorWithChainID(txQueue.accounts[i].PrivateKey, txQueue.transactionConfig.ChainId)
		if err != nil {
			txQueue.ErrChan <- &ErrorRequest{fmt.Errorf("CreateFastUpdatesClient: NewKeyedTransactorWithChainID: %w", err), compReq}
			continue
		}
		opts.Value = big.NewInt(int64(txQueue.transactionConfig.Value)) // in wei
		opts.GasLimit = uint64(txQueue.transactionConfig.GasLimit)      // in units
		opts.Nonce = big.NewInt(int64(nonce))

		ctx, cancelFunc = context.WithTimeout(context.Background(), time.Duration(config.CallTimeoutMillisDefault)*time.Millisecond)
		err = txQueue.setGasOptions(ctx, opts)
		cancelFunc()
		if err != nil {
			txQueue.ErrChan <- &ErrorRequest{err, compReq}
			continue
		}

		err = compReq.Function(opts)
		if err != nil {
			txQueue.ErrChan <- &ErrorRequest{err, compReq}
		} else {
			txQueue.ErrChan <- nil
		}

	}
}

// setGasOptions populates opts for an EIP-1559 (type 2) transaction. It derives a
// generous fee cap and tip from the next-block base fee so a submission can ride
// through a base-fee spike within the submission window without overpaying on
// normal blocks. GasPrice is left nil so go-ethereum's bind emits a DynamicFeeTx.
func (txQueue *TransactionQueue) setGasOptions(ctx context.Context, opts *bind.TransactOpts) error {
	cfg := txQueue.transactionConfig

	baseFee, err := txQueue.baseFee(ctx)
	if err != nil {
		return err
	}

	baseFeeMultiplier := cfg.BaseFeeMultiplier
	if baseFeeMultiplier <= 0 {
		baseFeeMultiplier = defaultBaseFeeCapMultiplier
	}
	tipMultiplier := cfg.MaxPriorityFeeMultiplier
	if tipMultiplier <= 0 {
		tipMultiplier = defaultTipMultiplier
	}
	minTip := cfg.MinimalMaxPriorityFee
	if minTip <= 0 {
		minTip = defaultMinPriorityFeeWei
	}
	maxTip := cfg.MaximalMaxPriorityFee
	if maxTip <= 0 {
		maxTip = defaultMaxPriorityFeeWei
	}

	gasTipCap := new(big.Int).Mul(baseFee, big.NewInt(tipMultiplier))
	// clamp the tip to [minTip, maxTip] so a low base fee still yields an
	// attractive tip and a spike does not balloon overpayment.
	if gasTipCap.Cmp(big.NewInt(minTip)) < 0 {
		gasTipCap = big.NewInt(minTip)
	} else if gasTipCap.Cmp(big.NewInt(maxTip)) > 0 {
		gasTipCap = big.NewInt(maxTip)
	}

	gasFeeCap := new(big.Int).Mul(baseFee, big.NewInt(baseFeeMultiplier))
	gasFeeCap.Add(gasFeeCap, gasTipCap)

	opts.GasTipCap = gasTipCap
	opts.GasFeeCap = gasFeeCap
	opts.GasPrice = nil // must be nil for bind to emit a DynamicFeeTx

	return nil
}

// baseFee returns the base fee per gas. It prefers eth_baseFee (the next-block
// base fee, available on Flare nodes), falling back to the latest block header's
// base fee on nodes that do not implement eth_baseFee (e.g. the ganache node used
// in integration tests).
func (txQueue *TransactionQueue) baseFee(ctx context.Context) (*big.Int, error) {
	var result hexutil.Big
	if err := txQueue.chainClient.Client().CallContext(ctx, &result, "eth_baseFee"); err == nil {
		return (*big.Int)(&result), nil
	}

	header, err := txQueue.chainClient.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("eth_baseFee unavailable and fetching latest header failed: %w", err)
	}
	if header.BaseFee == nil {
		return nil, fmt.Errorf("eth_baseFee unavailable and latest block has no base fee (pre-EIP-1559 chain?)")
	}
	return header.BaseFee, nil
}

func (txQueue *TransactionQueue) ErrorHandler() {
	// todo: define actions when error happens
	for {
		if txQueue.Stop {
			return
		}
		errReq := <-txQueue.ErrChan
		if errReq != nil {
			logger.Error("Error executing transaction: %s", errReq.Err)
		}
		txQueue.Lock()
		txQueue.numInQueue--
		txQueue.Unlock()
	}
}

func (txQueue *TransactionQueue) EmptyQueue() {
	txQueue.Lock()
	txQueue.toTrash = true
	txQueue.Unlock()
	for {
		if txQueue.numInQueue == 0 {
			txQueue.Lock()
			txQueue.toTrash = false
			txQueue.Unlock()
			return
		}
	}
}

func (txQueue *TransactionQueue) WaitToEmptyQueue() {
	// wait for all tx to come in the queue first
	time.Sleep(100 * time.Millisecond)

	for {
		if txQueue.numInQueue == 0 {
			return
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func (txQueue *TransactionQueue) WaitForBlockWithCheck(blockNum uint64) {
	for {
		if txQueue.toTrash {
			return
		}
		if txQueue.CurrentBlockNum < blockNum {
			time.Sleep(200 * time.Millisecond)
		} else {
			return
		}
	}
}

func (txQueue *TransactionQueue) CurrentBlockNumUpdater() {
	for {
		if txQueue.Stop {
			return
		}

		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(config.CallTimeoutMillisDefault)*time.Millisecond)
		currentBlockNum, err := txQueue.chainClient.BlockNumber(ctx)
		cancelFunc()
		if err != nil {
			logger.Error("failed obtaining current block number: %s", err)
			time.Sleep(200 * time.Millisecond)

			continue
		}
		if currentBlockNum > txQueue.CurrentBlockNum {
			txQueue.Lock()
			txQueue.CurrentBlockNum = currentBlockNum
			txQueue.Unlock()
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func (txQueue *TransactionQueue) StopQueue() {
	txQueue.Lock()
	txQueue.Stop = true
	txQueue.Unlock()
}
