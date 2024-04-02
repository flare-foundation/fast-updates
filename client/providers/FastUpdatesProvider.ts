import type { BytesLike } from 'ethers'
import { sha256 } from 'ethers'
import type { TransactionReceipt } from 'web3-types'
import type { Logger } from 'winston'

import { FEEDS } from '../../deployment/config'
import type { PriceDeltas, Proof } from '../utils/'
import {
    calculateRandomness,
    generateSortitionKey,
    generateVerifiableRandomnessProof,
} from '../utils/'
import { signMessage } from '../utils/'

import type { ExamplePriceFeedProvider } from './ExamplePriceFeedProvider'
import type { Web3Provider } from './Web3Provider'

/**
 * Represents a provider for fast updates.
 */
export class FastUpdatesProvider {
    private readonly key = generateSortitionKey()
    private lastRegisteredEpoch: number

    constructor(
        private readonly provider: Web3Provider,
        private readonly priceFeedProvider: ExamplePriceFeedProvider,
        private readonly epochLen: number,
        private readonly weight: number,
        private readonly address: string,
        private readonly privateKey: string,
        private readonly logger: Logger
    ) {
        this.lastRegisteredEpoch = -1
        this.logger.info(
            `FastUpdatesProvider initialized at ${provider.account.address}`
        )
    }

    /**
     * Initializes the FastUpdatesProvider.
     * Waits for a new reward epoch and registers as a voter for the next epoch.
     * @returns A Promise that resolves once the initialization is complete.
     */
    public async init(): Promise<void> {
        const nextEpoch = await this.provider.delayAfterRewardEpochEdge(
            this.epochLen
        )
        const txReceipt = await this.registerAsVoter(nextEpoch, this.weight)
        if (Number(txReceipt.status) === 1) {
            this.logger.info(
                `Block ${txReceipt.blockNumber}, Voter registration for epoch ${nextEpoch} successful`
            )
        } else {
            this.logger.error(
                `Block ${txReceipt.blockNumber}, Voter registration for epoch ${nextEpoch} failed`
            )
        }
    }

    /**
     * Registers the current user as a voter for a specific epoch with a given weight.
     * @param epoch - The epoch number to register for.
     * @param weight - The weight of the voter.
     * @param addToNonce - An optional value to add to the nonce.
     * @returns A promise that resolves to the transaction receipt.
     */
    private async registerAsVoter(
        epoch: number,
        weight: number,
        addToNonce?: number
    ): Promise<TransactionReceipt> {
        const txReceipt = await this.provider.registerAsAVoter(
            epoch,
            this.key,
            weight,
            this.address,
            addToNonce
        )
        this.lastRegisteredEpoch = epoch
        this.logger.info(
            `Gas consumed by registerAsVoter ${txReceipt.gasUsed.toString()}`
        )
        return txReceipt
    }

    /**
     * Submits updates to the provider.
     *
     * @param proof - The proof object.
     * @param replicate - The replicate value.
     * @param deltas - The deltas array.
     * @param submissionBlockNum - The submission block number.
     * @param addToNonce - Optional parameter to add to the nonce.
     * @returns A promise that resolves to the transaction receipt.
     */
    private async submitUpdates(
        proof: Proof,
        replicate: string,
        deltas: string,
        submissionBlockNum: string,
        addToNonce?: number
    ): Promise<TransactionReceipt> {
        const msg = this.provider.web3.eth.abi.encodeParameters(
            [
                'uint256',
                'uint256',
                'uint256',
                'uint256',
                'uint256',
                'uint256',
                'bytes',
            ],
            [
                submissionBlockNum,
                replicate,
                proof.gamma.x.toString(),
                proof.gamma.y.toString(),
                proof.c.toString(),
                proof.s.toString(),
                deltas,
            ]
        )

        const signature = signMessage(
            this.provider.web3,
            sha256(msg as BytesLike),
            this.privateKey
        )

        return this.provider.submitUpdates(
            proof,
            replicate,
            deltas,
            submissionBlockNum,
            signature,
            addToNonce
        )
    }

    /**
     * Retrieves the weight from the provider for the current address.
     * @returns A Promise that resolves to a string representing the weight.
     */
    private async getWeight(): Promise<string> {
        return await this.provider.getWeight(this.provider.account.address)
    }

    /**
     * Tries to submit updates based on the provided parameters.
     *
     * @param myWeight - The weight of the updates to be submitted.
     * @param blockNum - The block number for which the updates are being submitted.
     * @param seed - The seed value used for generating randomness.
     * @returns A Promise that resolves to void.
     */
    private async tryToSubmitUpdates(
        deltas: PriceDeltas,
        myWeight: number,
        blockNum: bigint,
        seed: string
    ): Promise<void> {
        let addToNonce = 0
        const blockNumStr = blockNum.toString()
        const cutoff = BigInt(await this.provider.getCurrentScoreCutoff())

        for (let replicate = 0; replicate < myWeight; replicate++) {
            const replicateStr = replicate.toString()
            const r: bigint = calculateRandomness(
                this.key,
                seed,
                blockNumStr,
                replicateStr
            )

            if (r < cutoff) {
                const proof: Proof = generateVerifiableRandomnessProof(
                    this.key,
                    seed,
                    blockNumStr,
                    replicateStr
                )
                this.logger.info(
                    `Block: ${blockNum}, Rep: ${replicate}, Update: ${deltas[1]}`
                )
                await this.submitUpdates(
                    proof,
                    replicateStr,
                    deltas[0],
                    blockNumStr,
                    addToNonce
                )
                    .then((receipt: TransactionReceipt) => {
                        this.logger.info(
                            `Block: ${receipt.blockNumber}, Update successful, ${deltas[1]}`
                        )
                    })
                    .catch((error: unknown) => {
                        this.logger.error(
                            `Block: ${blockNum}, Failed to submit updates ${error as string}`
                        )
                    })
                addToNonce++
                break
            }
        }
    }

    /**
     * Re-registers the provider for the next epoch.
     *
     * @param epoch - The current epoch number.
     * @returns A promise that resolves to a boolean indicating whether the re-registration was successful.
     */
    private async reRegister(epoch: number): Promise<boolean> {
        const txReceipt = await this.registerAsVoter(epoch + 1, this.weight)

        if (Number(txReceipt.status) === 1) {
            this.logger.info(
                `Epoch: ${epoch + 1} (Block ${txReceipt.blockNumber}), Registration successful`
            )
        } else {
            this.logger.error(`Epoch: ${epoch + 1}, Registration failed`)
        }
        return Number(txReceipt.status) === 1
    }

    /**
     * Runs the FastUpdatesProvider.
     * This method waits for a new epoch, performs necessary operations within each epoch,
     * and waits for the next block to continue the process.
     * @returns A Promise that resolves to void.
     */
    public async run(): Promise<void> {
        this.logger.info('Waiting for a new epoch...')
        await this.provider.waitForNewEpoch(this.epochLen)

        let currentWeight: string = ''
        let currentBaseSeed: bigint = 0n

        for (;;) {
            const blockNum = Number(await this.provider.getBlockNumber())
            const epoch = Math.floor((blockNum + 1) / this.epochLen)

            // Within the last 4 blocks of the epoch, re-register for the next epoch
            if (blockNum % this.epochLen >= this.epochLen - 4) {
                if (epoch + 1 > this.lastRegisteredEpoch) {
                    let registered = false
                    try {
                        registered = await this.reRegister(epoch)
                    } catch (error: unknown) {
                        this.logger.error(
                            `Block: ${blockNum}, Failed to re-register ${error as string}`
                        )
                    }
                    if (registered) {
                        // If successful, wait for the next epoch
                        await this.provider.waitForNewEpoch(this.epochLen)
                    } else {
                        // If failed, wait for the next block to try again
                        await this.provider.waitForBlock(blockNum + 1)
                        continue
                    }
                }
            }

            // Get the weight and the seed for the current epoch
            if (blockNum % this.epochLen === 0) {
                ;[currentWeight, currentBaseSeed] = await Promise.all([
                    this.getWeight(),
                    this.provider.getBaseSeed(),
                ])
                this.logger.info(`Epoch: ${epoch}, Weight: ${currentWeight}`)
            }

            // Fetch the on-chain and off-chain prices
            const onChainPrices: number[] = (
                await this.provider.fetchCurrentPrices(Array.from(FEEDS))
            ).map((x) => Number(x))
            const offChainPrices = this.priceFeedProvider.getCurrentPrices(
                Array.from(FEEDS)
            )

            this.logger.info(
                `blockNumber: ${blockNum}, onChainPrices: ${onChainPrices.join(', ')}, offChainPrices: ${offChainPrices.join(', ')}`
            )
            // Compare the on-chain and off-chain prices
            const deltas: PriceDeltas =
                this.priceFeedProvider.getFastUpdateDeltas(
                    onChainPrices,
                    offChainPrices
                )
            // Don't submit updates if there are no changes
            const rep = deltas[1]
            if (!rep.includes('+') && !rep.includes('-')) {
                this.logger.debug(`No updates for block ${blockNum}`)
            } else {
                try {
                    await this.tryToSubmitUpdates(
                        deltas,
                        Number(currentWeight),
                        BigInt(blockNum),
                        currentBaseSeed.toString()
                    )
                } catch (error: unknown) {
                    this.logger.error(
                        `Block: ${blockNum}, Failed to submit update ${error as string}`
                    )
                }
            }

            await this.provider.waitForBlock(blockNum + 1)
        }
    }
}
