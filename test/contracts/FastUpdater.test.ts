import path from 'path'

import type { BytesLike } from 'ethers'
import { sha256 } from 'ethers'
import type { Web3Account } from 'web3-eth-accounts'
import { encodePacked } from 'web3-utils'

import { signMessage } from '../../client/utils'
import {
    generateSortitionKey,
    generateVerifiableRandomnessProof,
} from '../../client/utils'
import type { Proof, SortitionKey } from '../../client/utils'
import { PATHS } from '../../deployment/config'
import {
    RangeFPA,
    SampleFPA,
    loadProviderAccounts,
    randomInt,
} from '../../deployment/utils'
import type {
    FastUpdateIncentiveManagerContract,
    FastUpdateIncentiveManagerInstance,
} from '../../typechain-truffle/contracts/fastUpdates/implementation/FastUpdateIncentiveManager'
import type {
    FastUpdaterContract,
    FastUpdaterInstance,
} from '../../typechain-truffle/contracts/fastUpdates/implementation/FastUpdater'
import type {
    FlareSystemMockContract,
    FlareSystemMockInstance,
} from '../../typechain-truffle/contracts/fastUpdates/test/FlareSystemMock'

const FastUpdater = artifacts.require('FastUpdater') as FastUpdaterContract
const FastUpdateIncentiveManager = artifacts.require(
    'FastUpdateIncentiveManager'
) as FastUpdateIncentiveManagerContract
const FlareSystemMock = artifacts.require(
    'FlareSystemMock'
) as FlareSystemMockContract

let TEST_REWARD_EPOCH: bigint

const EPOCH_LEN = 1000 as const
const NUM_ACCOUNTS = 3 as const
const VOTER_WEIGHT = 1000 as const
const SUBMISSION_WINDOW = 10 as const

const DURATION = 8 as const
const BASE_SAMPLE_SIZE = SampleFPA(2)
const BASE_RANGE = RangeFPA(2 ** -5)
const SAMPLE_INCREASE_LIMIT = SampleFPA(5)
const SCALE = 1 + BASE_RANGE / BASE_SAMPLE_SIZE
const RANGE_INCREASE_PRICE = 5 as const

const ANCHOR_PRICES = [1000, 10000]
const NUM_FEEDS = ANCHOR_PRICES.length

const ZEROS64 = '0x' + '0'.repeat(64)
const ZEROS52 = '0x' + '0'.repeat(52)

contract(
    `FastUpdater.sol; ${path.relative(path.resolve(), __filename)}`,
    () => {
        let fastUpdater: FastUpdaterInstance
        let fastUpdateIncentiveManager: FastUpdateIncentiveManagerInstance
        let flareSystemMock: FlareSystemMockInstance
        let accounts: Web3Account[]
        let sortitionKeys: SortitionKey[]
        const weights: number[] = []

        before(async () => {
            accounts = loadProviderAccounts(web3, PATHS.accountsPath)
            const governance = accounts[0]
            if (!governance) {
                throw new Error('No governance account')
            }

            flareSystemMock = await FlareSystemMock.new(
                randomInt(2n ** 256n - 1n).toString(),
                EPOCH_LEN
            )
            fastUpdateIncentiveManager = await FastUpdateIncentiveManager.new(
                governance.address,
                governance.address,
                BASE_SAMPLE_SIZE,
                BASE_RANGE,
                SAMPLE_INCREASE_LIMIT,
                RANGE_INCREASE_PRICE,
                DURATION
            )

            TEST_REWARD_EPOCH = BigInt(
                (await flareSystemMock.getCurrentRewardEpochId()).toString()
            )

            sortitionKeys = new Array<SortitionKey>(NUM_ACCOUNTS)
            for (let i = 0; i < NUM_ACCOUNTS; i++) {
                const key: SortitionKey = generateSortitionKey()
                sortitionKeys[i] = key
                const x = '0x' + web3.utils.padLeft(key.pk.x.toString(16), 64)
                const y = '0x' + web3.utils.padLeft(key.pk.y.toString(16), 64)
                const policy = {
                    pk_1: x,
                    pk_2: y,
                    weight: VOTER_WEIGHT,
                }
                await flareSystemMock.registerAsVoter(
                    TEST_REWARD_EPOCH.toString(),
                    accounts[i + 1]?.address ?? '',
                    policy
                )
            }

            // Create local instance of Fast Updater contract
            fastUpdater = await FastUpdater.new(
                governance.address,
                flareSystemMock.address,
                flareSystemMock.address,
                fastUpdateIncentiveManager.address,
                ANCHOR_PRICES,
                SUBMISSION_WINDOW
            )
        })

        it('should submit updates', async () => {
            let submissionBlockNum

            for (let i = 0; i < NUM_ACCOUNTS; i++) {
                const weight = await fastUpdater.currentSortitionWeight(
                    (accounts[i + 1] as Web3Account).address
                )
                weights[i] = weight.toNumber()
                expect(weights[i]).to.equal(Math.floor(4096 / NUM_ACCOUNTS))
            }

            // Fetch current feed prices from the contract
            const feeds: number[] = []
            for (let i = 0; i < NUM_FEEDS; i++) {
                feeds.push(i)
            }
            const startingPrices: number[] = (
                await fastUpdater.fetchCurrentPrices(feeds)
            ).map((x: BN) => x.toNumber())

            // Make price updates to the contract
            const feed = '+-0-0+'
            let delta = '0x731' + '0'.repeat(61)
            let numSubmitted = 0
            for (;;) {
                submissionBlockNum = (
                    await web3.eth.getBlockNumber()
                ).toString()
                const scoreCutoff = BigInt(
                    (await fastUpdater.currentScoreCutoff()).toString()
                )
                const baseSeed = (
                    await flareSystemMock.getCurrentRandom()
                ).toString()
                for (let i = 0; i < NUM_ACCOUNTS; i++) {
                    for (let rep = 0; rep < (weights[i] ?? 0); rep++) {
                        const repStr = rep.toString()
                        const proof: Proof = generateVerifiableRandomnessProof(
                            sortitionKeys[i] as SortitionKey,
                            baseSeed,
                            submissionBlockNum,
                            repStr
                        )

                        const sortitionCredential = {
                            replicate: repStr,
                            gamma: {
                                x: proof.gamma.x.toString(),
                                y: proof.gamma.y.toString(),
                            },
                            c: proof.c.toString(),
                            s: proof.s.toString(),
                        }

                        if (proof.gamma.x < scoreCutoff) {
                            const deltas = {
                                mainParts: [
                                    delta,
                                    ZEROS64,
                                    ZEROS64,
                                    ZEROS64,
                                    ZEROS64,
                                    ZEROS64,
                                    ZEROS64,
                                ],
                                tailPart: ZEROS52,
                            }
                            const msg = encodePacked(
                                {
                                    value: submissionBlockNum.toString(),
                                    type: 'uint256',
                                },
                                { value: repStr, type: 'uint256' },
                                {
                                    value: proof.gamma.x.toString(),
                                    type: 'uint256',
                                },
                                {
                                    value: proof.gamma.y.toString(),
                                    type: 'uint256',
                                },
                                { value: proof.c.toString(), type: 'uint256' },
                                { value: proof.s.toString(), type: 'uint256' },
                                {
                                    value: deltas.mainParts[0] || '',
                                    type: 'bytes32',
                                },
                                { value: ZEROS64, type: 'bytes32' },
                                { value: ZEROS64, type: 'bytes32' },
                                { value: ZEROS64, type: 'bytes32' },
                                { value: ZEROS64, type: 'bytes32' },
                                { value: ZEROS64, type: 'bytes32' },
                                { value: ZEROS64, type: 'bytes32' },
                                { value: ZEROS52, type: 'bytes32' }
                            )
                            const signature = signMessage(
                                web3,
                                sha256(msg as BytesLike),
                                (accounts[i + 1] as Web3Account).privateKey
                            )
                            const newFastUpdate = {
                                sortitionBlock: submissionBlockNum,
                                sortitionCredential: sortitionCredential,
                                deltas: deltas,
                                signature: signature,
                            }

                            // Submit updates to the contract
                            await fastUpdater.submitUpdates(newFastUpdate, {
                                from: accounts[i + 1]?.address ?? '',
                            })
                            numSubmitted++
                        }
                    }
                }
                if (numSubmitted > 0) break
                await fastUpdater.freeSubmitted({
                    from: accounts[0]?.address ?? '',
                })
            }

            // See effect of price updates made
            const prices1: number[] = await fastUpdater
                .fetchCurrentPrices(feeds)
                .then((prices) => {
                    return prices.map((x: BN) => x.toNumber())
                })
            for (let i = 0; i < NUM_FEEDS; i++) {
                let sign = 0
                if (feed[i] == '+') {
                    sign = 1
                }
                if (feed[i] == '-') {
                    sign = -1
                }
                expect(prices1[i]).to.be.greaterThanOrEqual(
                    (SCALE ** sign) ** numSubmitted *
                        (startingPrices[i] as number) *
                        0.99
                )
                expect(prices1[i]).to.be.lessThanOrEqual(
                    (SCALE ** sign) ** numSubmitted *
                        (startingPrices[i] as number) *
                        1.01
                )
            }

            delta = '0xd13' + '0'.repeat(61)
            let breakVar = false
            while (!breakVar) {
                submissionBlockNum = String(await web3.eth.getBlockNumber())

                const scoreCutoff = BigInt(
                    (await fastUpdater.currentScoreCutoff()).toString()
                )
                const baseSeed = (
                    await flareSystemMock.getCurrentRandom()
                ).toString()

                for (let i = 0; i < NUM_ACCOUNTS; i++) {
                    for (let rep = 0; rep < (weights[i] ?? 0); rep++) {
                        const repStr = rep.toString()
                        const proof: Proof = generateVerifiableRandomnessProof(
                            sortitionKeys[i] as SortitionKey,
                            baseSeed,
                            submissionBlockNum,
                            repStr
                        )
                        const sortitionCredential = {
                            replicate: repStr,
                            gamma: {
                                x: proof.gamma.x.toString(),
                                y: proof.gamma.y.toString(),
                            },
                            c: proof.c.toString(),
                            s: proof.s.toString(),
                        }

                        if (proof.gamma.x < scoreCutoff) {
                            const deltas = {
                                mainParts: [
                                    delta,
                                    ZEROS64,
                                    ZEROS64,
                                    ZEROS64,
                                    ZEROS64,
                                    ZEROS64,
                                    ZEROS64,
                                ],
                                tailPart: ZEROS52,
                            }
                            const msg = encodePacked(
                                {
                                    value: submissionBlockNum.toString(),
                                    type: 'uint256',
                                },
                                { value: repStr, type: 'uint256' },
                                {
                                    value: proof.gamma.x.toString(),
                                    type: 'uint256',
                                },
                                {
                                    value: proof.gamma.y.toString(),
                                    type: 'uint256',
                                },
                                { value: proof.c.toString(), type: 'uint256' },
                                { value: proof.s.toString(), type: 'uint256' },
                                {
                                    value: deltas.mainParts[0] || '',
                                    type: 'bytes32',
                                },
                                { value: ZEROS64, type: 'bytes32' },
                                { value: ZEROS64, type: 'bytes32' },
                                { value: ZEROS64, type: 'bytes32' },
                                { value: ZEROS64, type: 'bytes32' },
                                { value: ZEROS64, type: 'bytes32' },
                                { value: ZEROS64, type: 'bytes32' },
                                { value: ZEROS52, type: 'bytes32' }
                            )
                            const signature = signMessage(
                                web3,
                                sha256(msg as BytesLike),
                                (accounts[i + 1] as Web3Account).privateKey
                            )
                            const newFastUpdate = {
                                sortitionBlock: submissionBlockNum,
                                sortitionCredential: sortitionCredential,
                                deltas: deltas,
                                signature: signature,
                            }

                            await fastUpdater.submitUpdates(newFastUpdate, {
                                from: accounts[i + 1]?.address ?? '',
                            })
                            numSubmitted--
                            if (numSubmitted == 0) {
                                breakVar = true
                                break
                            }
                        }
                    }
                    if (breakVar) break
                }

                await fastUpdater.freeSubmitted({
                    from: accounts[0]?.address ?? '',
                })
            }

            const prices2: number[] = await fastUpdater
                .fetchCurrentPrices(feeds)
                .then((prices) => {
                    return prices.map((x: BN) => x.toNumber())
                })
            for (let i = 0; i < NUM_FEEDS; i++) {
                expect(prices2[i]).to.be.greaterThanOrEqual(
                    (startingPrices[i] as number) * 0.99
                )
                expect(prices2[i]).to.be.lessThanOrEqual(
                    (startingPrices[i] as number) * 1.01
                )
            }
        })
    }
)
