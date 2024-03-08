import type Web3 from 'web3'
import type { ContractOptions } from 'web3'
import type { Web3Account } from 'web3-eth-accounts'
import type { Transaction, TransactionReceipt } from 'web3-types'

import type { FastUpdateIncentiveManager } from '../../typechain-web3/contracts/fastUpdates/implementation/FastUpdateIncentiveManager'
import type { FastUpdater } from '../../typechain-web3/contracts/fastUpdates/implementation/FastUpdater'
import type {
    FlareSystemManager,
    VoterRegistry,
} from '../../typechain-web3/contracts/fastUpdates/interface/mocks'
import type { FlareSystemMock } from '../../typechain-web3/contracts/fastUpdates/test'
import type { NonPayableTransactionObject } from '../../typechain-web3/types'
import { getOrCreateLogger, loadContract, sleepFor } from '../utils'
import type {
    FastUpdatesContractAddresses,
    NetworkParameters,
    Proof,
    SortitionKey,
} from '../utils'
import type { BareSignature } from '../utils'

interface TypeChainContracts {
    readonly voterRegistry: VoterRegistry
    readonly flareSystemManager: FlareSystemManager
    readonly fastUpdater: FastUpdater
    readonly fastUpdateIncentiveManager: FastUpdateIncentiveManager
    readonly flareSystemMock: FlareSystemMock
}

type SortitionCredential = [string, [string, string], string, string]

type FastUpdate = [
    string,
    [string, [string, string], string, string],
    [string[], string],
    [number, string, string],
]

type SortitionPolicy = [string, string, number]

export class Web3Provider {
    private readonly logger = getOrCreateLogger(Web3Provider.name)

    private constructor(
        readonly web3: Web3,
        readonly contracts: TypeChainContracts,
        private readonly config: NetworkParameters,
        readonly account: Web3Account
    ) {
        this.account = account
        this.logger.info(
            `Creating Web3Provider instance for ${this.account.address}`
        )
    }

    /**
     * Creates a new instance of Web3Provider.
     * Separated from the constructor to allow for async initialization.
     *
     * @param web3 - The Web3 instance.
     * @param contractAddresses - The contract addresses.
     * @param params - The network parameters.
     * @param account - The Web3 account.
     * @returns A Promise that resolves to a new instance of Web3Provider.
     */
    public static async create(
        web3: Web3,
        contractAddresses: FastUpdatesContractAddresses,
        params: NetworkParameters,
        account: Web3Account
    ): Promise<Web3Provider> {
        const contracts = {
            voterRegistry: await loadContract<VoterRegistry>(
                web3,
                contractAddresses.voterRegistry,
                'VoterRegistry'
            ),
            flareSystemManager: await loadContract<FlareSystemManager>(
                web3,
                contractAddresses.flareSystemManager,
                'FlareSystemManager'
            ),
            fastUpdater: await loadContract<FastUpdater>(
                web3,
                contractAddresses.fastUpdater,
                'FastUpdater'
            ),
            flareSystemMock: await loadContract<FlareSystemMock>(
                web3,
                contractAddresses.flareSystemManager,
                'FlareSystemMock'
            ),
            fastUpdateIncentiveManager:
                await loadContract<FastUpdateIncentiveManager>(
                    web3,
                    contractAddresses.fastUpdateIncentiveManager,
                    'FastUpdateIncentiveManager'
                ),
        }
        const provider = new Web3Provider(web3, contracts, params, account)
        return provider
    }

    /**
     * Signs and finalizes a transaction.
     *
     * @param label - The label for the transaction.
     * @param toAddress - The address to send the transaction to.
     * @param fnToEncode - The function to encode in the transaction.
     * @param value - The value to send with the transaction.
     * @param from - The account to send the transaction from.
     * @param addToNonce - The value to add to the transaction nonce.
     * @returns A promise that resolves to the transaction receipt.
     * @throws If the dry-run of the transaction fails or if the state changes during the transaction.
     */
    private async signAndFinalize(
        label: string,
        toAddress: string,
        fnToEncode: NonPayableTransactionObject<void>,
        value: string = '0',
        from: Web3Account = this.account,
        addToNonce: number = 0
    ): Promise<TransactionReceipt> {
        // Calculate tx parameters
        const [gasPrice, currentNonce] = await Promise.all([
            this.web3.eth.getGasPrice(),
            this.getNonce(from),
        ])
        const tx: Transaction = {
            from: from.address,
            to: toAddress,
            gas: this.config.gasLimit,
            gasPrice: gasPrice,
            data: fnToEncode.encodeABI(),
            value: value,
            nonce: Number(currentNonce) + addToNonce,
        }

        // Send tx and wait for confirmation
        this.logger.debug(`Sending tx ${label} (nonce:${tx.nonce})`)
        const signedTx = await from.signTransaction(tx)
        const txReceipt: TransactionReceipt = await this.web3.eth
            .sendSignedTransaction(signedTx.rawTransaction)
            .on('confirmation', (confirmation) => {
                this.logger.debug(
                    `Confirmed tx ${label} (confirmations:${confirmation.confirmations}, nonce:${tx.nonce})`
                )
            })
            .on('error', (error) => {
                this.logger.error(`Error in tx ${label}: ${error.message}`)
            })

        return txReceipt
    }

    /**
     * Registers a voter with the specified epoch, sortition key and weight.
     * @param epoch - The epoch number.
     * @param key - The sortition key.
     * @param weight - The weight of the voter.
     * @param addToNonce - Optional value to add to the nonce.
     * @returns A promise that resolves to the transaction receipt.
     */
    public async registerAsAVoter(
        epoch: number,
        key: SortitionKey,
        weight: number,
        address: string,
        addToNonce?: number
    ): Promise<TransactionReceipt> {
        const x =
            '0x' +
            '0'.repeat(64 - key.pk.x.toString(16).length) +
            key.pk.x.toString(16)
        const y =
            '0x' +
            '0'.repeat(64 - key.pk.y.toString(16).length) +
            key.pk.y.toString(16)

        const policy: SortitionPolicy = [x, y, weight]
        const methodCall =
            this.contracts.flareSystemMock.methods.registerAsVoter(
                epoch,
                address,
                policy
            )
        return this.signAndFinalize(
            'RegisterAsAVoter',
            (this.contracts.voterRegistry['options'] as ContractOptions)
                .address as string,
            methodCall,
            undefined,
            undefined,
            addToNonce
        )
    }

    /**
     * Submits updates to the fast updater contract.
     *
     * @param proof - The proof object containing the necessary cryptographic proofs.
     * @param replicate - The number of times to replicate the proof.
     * @param deltas - The deltas to be submitted.
     * @param submissionBlockNum - The block number of the submission.
     * @param addToNonce - Optional parameter to add to the nonce.
     * @returns A promise that resolves to the transaction receipt.
     */
    public async submitUpdates(
        proof: Proof,
        replicate: string,
        deltas: [string[], string],
        submissionBlockNum: string,
        signature: BareSignature,
        addToNonce?: number
    ): Promise<TransactionReceipt> {
        const sortitionCredential: SortitionCredential = [
            replicate,
            [proof.gamma.x.toString(), proof.gamma.y.toString()],
            proof.c.toString(),
            proof.s.toString(),
        ]
        const newFastUpdate: FastUpdate = [
            submissionBlockNum,
            sortitionCredential,
            deltas,
            [signature.v, signature.r, signature.s],
        ]

        const methodCall =
            this.contracts.fastUpdater.methods.submitUpdates(newFastUpdate)

        return this.signAndFinalize(
            'submitUpdates',
            (this.contracts.fastUpdater['options'] as ContractOptions)
                .address as string,
            methodCall,
            undefined,
            undefined,
            addToNonce
        )
    }

    /**
     * Advances the incentive by calling the 'advance' method of the fastUpdateIncentiveManager contract.
     * To be called by the daemon.
     *
     * @param addToNonce - Optional parameter to add to the nonce value.
     * @returns A promise that resolves to a TransactionReceipt object.
     */
    public async advanceIncentive(
        addToNonce?: number
    ): Promise<TransactionReceipt> {
        const methodCall =
            this.contracts.fastUpdateIncentiveManager.methods.advance()
        return this.signAndFinalize(
            'advance',
            (
                this.contracts.fastUpdateIncentiveManager[
                    'options'
                ] as ContractOptions
            ).address as string,
            methodCall,
            undefined,
            undefined,
            addToNonce
        )
    }

    /**
     * Calls the 'freeSubmitted' method of the 'fastUpdater' contract.
     * To be called by the daemon.
     *
     * @param addToNonce - Optional parameter to add to the nonce value.
     * @returns A promise that resolves to a TransactionReceipt object.
     */
    public async freeSubmitted(
        addToNonce?: number
    ): Promise<TransactionReceipt> {
        const methodCall = this.contracts.fastUpdater.methods.freeSubmitted()
        return this.signAndFinalize(
            'freeSubmitted',
            (this.contracts.fastUpdater['options'] as ContractOptions)
                .address as string,
            methodCall,
            undefined,
            undefined,
            addToNonce
        )
    }

    /**
     * Retrieves the sortition weight for a given address.
     * @param address - The address for which to retrieve the sortition weight.
     * @returns A promise that resolves to the sortition weight as a string.
     */
    public async getWeight(address: string): Promise<string> {
        const weight = await this.contracts.fastUpdater.methods
            .currentSortitionWeight(address)
            .call()
        return weight
    }

    /**
     * Fetches the current prices for the given feeds.
     * @param feeds An array of feed numbers.
     * @returns A promise that resolves to an array of strings representing the current prices.
     */
    public async fetchCurrentPrices(feeds: number[]): Promise<string[]> {
        return await this.contracts.fastUpdater.methods
            .fetchCurrentPrices(feeds)
            .call()
    }

    /**
     * Retrieves the base seed from the Flare System Manager contract.
     * @returns A promise that resolves to a bigint representing the base seed.
     */
    public async getBaseSeed(): Promise<bigint> {
        return BigInt(
            await this.contracts.flareSystemManager.methods
                .getCurrentRandom()
                .call()
        )
    }

    /**
     * Retrieves the current score cutoff from the fastUpdater contract.
     * @returns A Promise that resolves to a string representing the current score cutoff.
     */
    async getCurrentScoreCutoff(): Promise<string> {
        return await this.contracts.fastUpdater.methods
            .currentScoreCutoff()
            .call()
    }

    /**
     * Retrieves the current block number from the Ethereum network.
     * @returns A promise that resolves to a bigint representing the block number.
     */
    public async getBlockNumber(): Promise<bigint> {
        return await this.web3.eth.getBlockNumber()
    }

    /**
     * Retrieves the nonce for the specified Web3 account.
     * @param account The Web3 account.
     * @returns A promise that resolves to the nonce as a bigint.
     */
    private async getNonce(account: Web3Account): Promise<bigint> {
        return await this.web3.eth.getTransactionCount(account.address)
    }

    /**
     * Retrieves the genesis time (UNIX epoch) of the chain in seconds.
     * @returns A promise that resolves to the genesis time in seconds.
     */
    public async getGenesisTimeSec(): Promise<number> {
        // The timestamp conversion requires first converting to string, then to Number
        // Since the default JS library uses Number for timestamps, this conversion is safe
        return Number(
            (await this.web3.eth.getBlock('earliest')).timestamp.toString()
        )
    }

    /**
     * Waits for a new epoch to begin.
     * @param epochLen The length of each epoch.
     * @returns A Promise that resolves when the new epoch begins.
     */
    public async waitForNewEpoch(epochLen: number): Promise<void> {
        const blockNum = Number(await this.getBlockNumber())
        const newEpochBlockNum =
            (Math.floor(blockNum / epochLen) + 1) * epochLen

        await this.waitForBlock(newEpochBlockNum)
    }

    /**
     * Waits for a specific block number to be mined.
     * @param minedBlockNum The block number to wait for.
     * @returns A promise that resolves when the specified block number is mined.
     */
    public async waitForBlock(minedBlockNum: number): Promise<void> {
        for (;;) {
            const blockNum = await this.getBlockNumber()
            if (blockNum >= minedBlockNum) {
                return
            }
            await sleepFor(1000)
        }
    }

    /**
     * Delays execution until the reward epoch edge is reached.
     * If the reward epoch is just changing, it waits for the specified number of blocks.
     * @param epochLen The length of each reward epoch.
     * @param delayBlocks The number of blocks to delay after the reward epoch edge. Default is 3.
     * @returns The next epoch number.
     */
    public async delayAfterRewardEpochEdge(
        epochLen: number,
        delayBlocks: number = 3
    ): Promise<number> {
        let currentBlock = Number(await this.getBlockNumber())
        // If the reward epoch is just changing, wait
        if (
            currentBlock % epochLen == 0 ||
            (currentBlock + 1) % epochLen == 0 ||
            (currentBlock + 2) % epochLen == 0
        ) {
            await this.waitForBlock(currentBlock + delayBlocks)
            currentBlock = Number(await this.getBlockNumber())
        }

        const nextEpoch = Math.floor((currentBlock + 1) / epochLen) + 1
        return nextEpoch
    }
}
