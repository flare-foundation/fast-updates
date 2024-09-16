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
	"github.com/ethereum/go-ethereum/ethclient"
)

const queueLen = 50

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
		ctx, cancelFunc = context.WithTimeout(context.Background(), time.Duration(config.CallTimeoutMillisDefault)*time.Millisecond)
		gasPrice, err := txQueue.chainClient.SuggestGasPrice(ctx)
		cancelFunc()
		if err != nil {
			txQueue.ErrChan <- &ErrorRequest{err, compReq}
			continue
		}
		opts, err := bind.NewKeyedTransactorWithChainID(txQueue.accounts[i].PrivateKey, txQueue.transactionConfig.ChainId)
		if err != nil {
			txQueue.ErrChan <- &ErrorRequest{fmt.Errorf("CreateFastUpdatesClient: NewKeyedTransactorWithChainID: %w", err), compReq}
		}
		opts.Value = big.NewInt(int64(txQueue.transactionConfig.Value)) // in wei
		opts.GasLimit = uint64(txQueue.transactionConfig.GasLimit)      // in units

		opts.Nonce = big.NewInt(int64(nonce))
		gasPriceFloat := new(big.Float).Mul(new(big.Float).SetInt(gasPrice), big.NewFloat(txQueue.transactionConfig.GasPriceMultiplier))
		opts.GasPrice, _ = gasPriceFloat.Int(nil)

		err = compReq.Function(opts)
		if err != nil {
			txQueue.ErrChan <- &ErrorRequest{err, compReq}
		} else {
			txQueue.ErrChan <- nil
		}

	}
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
