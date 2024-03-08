import path from 'path'

import { bn254 } from '@noble/curves/bn254'

import {
    generateSortitionKey,
    generateVerifiableRandomnessProof,
    randomInt,
} from '../../client/utils'
import type { Proof, SortitionKey } from '../../client/utils'
import type {
    TestSortitionContractContract,
    TestSortitionContractInstance,
} from '../../typechain-truffle/contracts/fastUpdates/test/TestSortition.sol/TestSortitionContract'

const SortitionContract = artifacts.require(
    'TestSortitionContract'
) as TestSortitionContractContract

contract(
    `Sortition.sol; ${path.relative(path.resolve(), __filename)}`,
    (accounts) => {
        let sortition: TestSortitionContractInstance
        before(async () => {
            const governance = accounts[0]
            if (!governance) {
                throw new Error('No governance account')
            }
            sortition = await SortitionContract.new(
                governance as Truffle.TransactionDetails
            )
        })

        it('should generate a verifiable randomness', async () => {
            const key: SortitionKey = generateSortitionKey()
            const seed = randomInt(bn254.CURVE.n).toString()
            const blockNum = (await web3.eth.getBlockNumber()).toString()
            const replicate = randomInt(bn254.CURVE.n).toString()
            const proof: Proof = generateVerifiableRandomnessProof(
                key,
                seed,
                blockNum,
                replicate
            )
            const pubKey = { x: key.pk.x.toString(), y: key.pk.y.toString() }
            const sortitionCredential = {
                replicate: replicate,
                gamma: {
                    x: proof.gamma.x.toString(),
                    y: proof.gamma.y.toString(),
                },
                c: proof.c.toString(),
                s: proof.s.toString(),
            }
            const sortitionState = {
                baseSeed: seed,
                blockNumber: blockNum,
                scoreCutoff: 0,
                weight: 0,
                pubKey: pubKey,
            }

            const check = await sortition.testVerifySortitionProof(
                sortitionState,
                sortitionCredential
            )

            expect(check).to.equal(true)
        })
        it('should correctly accept or reject the randomness', async () => {
            const key: SortitionKey = generateSortitionKey()
            const scoreCutoff = 2n ** 248n
            for (;;) {
                const seed = randomInt(bn254.CURVE.n).toString()
                const replicate = randomInt(bn254.CURVE.n)
                const blockNum = (await web3.eth.getBlockNumber()).toString()
                const weight = replicate + 1n

                const proof: Proof = generateVerifiableRandomnessProof(
                    key,
                    seed,
                    blockNum,
                    replicate.toString()
                )
                const pubKey = {
                    x: key.pk.x.toString(),
                    y: key.pk.y.toString(),
                }
                const sortitionCredential = {
                    replicate: replicate.toString(),
                    gamma: {
                        x: proof.gamma.x.toString(),
                        y: proof.gamma.y.toString(),
                    },
                    c: proof.c.toString(),
                    s: proof.s.toString(),
                }
                const sortitionState = {
                    baseSeed: seed,
                    blockNumber: blockNum,
                    scoreCutoff: scoreCutoff.toString(),
                    weight: weight.toString(),
                    pubKey: pubKey,
                }

                const check = await sortition.testVerifySortitionCredential(
                    sortitionState,
                    sortitionCredential
                )

                if (proof.gamma.x > scoreCutoff) {
                    expect(check).to.equal(false)
                } else {
                    expect(check).to.equal(true)
                    break
                }
            }
        })
    }
)
