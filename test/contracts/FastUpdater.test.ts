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
    RangeOrSampleFPA,
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
const NUM_ACCOUNTS = 5

const NUM_FEEDS = 256
const ANCHOR_PRICES = [5000, 10000, 20000, 30000, 40000, 50000, 60000, 70000]
for (let i = 8; i < NUM_FEEDS; i++) {
    ANCHOR_PRICES.push(i * 10000)
}

const VOTER_WEIGHT = 1000
const SUBMISSION_WINDOW = 10
const BASE_SAMPLE_SIZE = 16
const BASE_RANGE = 2**-5
const SAMPLE_INCREASE_LIMIT = 5
const SCALE = 1 + BASE_RANGE / BASE_SAMPLE_SIZE
const RANGE_INCREASE_PRICE = 16
const DURATION = 8
const EPOCH_LEN = 1000
const BACKLOG_LEN = 20

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
                RangeOrSampleFPA(BASE_SAMPLE_SIZE),
                RangeOrSampleFPA(BASE_RANGE),
                RangeOrSampleFPA(SAMPLE_INCREASE_LIMIT),
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
                SUBMISSION_WINDOW,
                BACKLOG_LEN
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

            // test with feeds of various length
            let feed = '+--+00--'.repeat(16)
            let deltas = '0x' + '7d0f'.repeat(16)
            const differentFeed = '-+0000++'.repeat(8) + '-+00'
            let differentDeltas = 'd005'.repeat(8) + 'd0'
            deltas += differentDeltas
            feed += differentFeed
            differentDeltas = '0x' + differentDeltas

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
                    submissionBlockNum = (
                        await web3.eth.getBlockNumber()
                    ).toString()

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
                            let update = deltas
                            if (numSubmitted == 1) {
                                // use a different update with different length for this test
                                update = differentDeltas
                            }

                            const toHash = web3.eth.abi.encodeParameters(
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
                                    repStr,
                                    proof.gamma.x.toString(),
                                    proof.gamma.y.toString(),
                                    proof.c.toString(),
                                    proof.s.toString(),
                                    update,
                                ]
                            )
                            const signature = signMessage(
                                web3,
                                sha256(toHash as BytesLike),
                                (accounts[i + 1] as Web3Account).privateKey
                            )
                            const newFastUpdate = {
                                sortitionBlock: submissionBlockNum,
                                sortitionCredential: sortitionCredential,
                                deltas: update,
                                signature: signature,
                            }

                            // Submit updates to the contract
                            const tx = await fastUpdater.submitUpdates(
                                newFastUpdate,
                                {
                                    from: accounts[i + 1]?.address ?? '',
                                }
                            )
                            // console.log('cost', tx.receipt.gasUsed)

                            // let caughtError = false
                            // try {
                            //     // test if submitting again gives error
                            //     console.log('gere2')

                            //     const tx = await fastUpdater.submitUpdates(
                            //         newFastUpdate,
                            //         {
                            //             from: accounts[i + 1]?.address ?? '',
                            //         }
                            //     )
                            //     console.log('gere3', tx)
                            // } catch (e) {
                            //     expect(e).to.be.not.empty
                            //     caughtError = true
                            // }
                            // console.log('gere')
                            // expect(caughtError).to.equal(true)

                            numSubmitted++
                            if (numSubmitted >= 2) break
                        }
                    }
                    await fastUpdater.freeSubmitted({
                        from: accounts[0]?.address ?? '',
                    })
                    if (numSubmitted >= 2) break
                }
                if (numSubmitted > 0) break
            }

            // See effect of price updates made
            let pricesBN: BN[] = await fastUpdater.fetchCurrentPrices(feeds)
            const prices: number[] = []
            for (let i = 0; i < NUM_FEEDS; i++) {
                prices[i] = pricesBN[i]!.toNumber()
                let newPrice = startingPrices[i]!
                for (let j = 0; j < numSubmitted; j++) {
                    let delta = feed[i]
                    if (j == 1) {
                        delta = differentFeed[i]
                    }

                    if (delta == '+') {
                        newPrice *= SCALE
                    }
                    if (delta == '-') {
                        newPrice /= SCALE
                    }
                    newPrice = Math.floor(newPrice)
                }

                expect(prices[i]).to.be.equal(newPrice)
            }

            console.log('applying deltas')
            const tx = await fastUpdater.applySubmitted({
                from: accounts[0]!.address,
            })
            // console.log('cost2', tx.receipt.gasUsed)

            pricesBN = await fastUpdater.fetchCurrentPrices.call(feeds)
            for (let i = 0; i < NUM_FEEDS; i++) {
                prices[i] = pricesBN[i]!.toNumber()
                let newPrice = startingPrices[i]!
                for (let j = 0; j < numSubmitted; j++) {
                    let delta = feed[i]
                    if (j == 1) {
                        delta = differentFeed[i]
                    }

                    if (delta == '+') {
                        newPrice *= SCALE
                    }
                    if (delta == '-') {
                        newPrice /= SCALE
                    }
                    newPrice = Math.floor(newPrice)
                }
                expect(prices[i]).to.be.equal(newPrice)
            }
        })
    }
)
