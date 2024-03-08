import { expect } from 'chai'
import { encodePacked } from 'web3-utils'

import {
    calculateRandomness,
    g1HashToPoint,
    g1compress,
    generateSortitionKey,
    generateVerifiableRandomnessProof,
    randomInt,
} from '../../client/utils/sortition'

describe('Sortition', () => {
    describe('generateSortitionKey', () => {
        it('should generate a sortition key with sk and pk properties', () => {
            const key = generateSortitionKey()

            expect(key).to.have.property('sk').that.is.a('bigint')
            expect(key).to.have.property('pk')
        })
    })

    describe('calculateRandomness', () => {
        it('should calculate the randomness value based on the provided parameters', () => {
            const key = generateSortitionKey()
            const baseSeed = '123'
            const blockNum = '456'
            const replicate = '789'

            const randomness = calculateRandomness(
                key,
                baseSeed,
                blockNum,
                replicate
            )

            expect(randomness).to.be.a('bigint')
        })
    })

    describe('generateVerifiableRandomnessProof', () => {
        it('should generate a verifiable randomness proof', () => {
            const key = generateSortitionKey()
            const baseSeed = '123'
            const blockNum = '456'
            const replicate = '789'

            const proof = generateVerifiableRandomnessProof(
                key,
                baseSeed,
                blockNum,
                replicate
            )

            expect(proof).to.have.property('gamma')
            expect(proof).to.have.property('c').that.is.a('bigint')
            expect(proof).to.have.property('s').that.is.a('bigint')
        })
    })

    describe('g1compress', () => {
        it('should compress a ProjPointType<bigint> into a string representation', () => {
            const baseSeed = 123n
            const blockNum = 456n
            const replicate = 789n

            const msg: string =
                encodePacked(
                    { value: baseSeed.toString(), type: 'uint256' },
                    { value: blockNum.toString(), type: 'uint256' },
                    { value: replicate.toString(), type: 'uint256' }
                ) ?? ''

            const projPoint = g1HashToPoint(msg)
            const compressed = g1compress(projPoint)
            expect(compressed).to.equal(
                '0x8aa3dd625c54203caad154a15acb97a7b32616ecfb25ec1f75dbaff55f1482ff'
            )
        })
    })
})

describe('randomInt', () => {
    it('should return a random integer less than the given max', () => {
        const max = 100n
        const result = randomInt(max)

        expect(result).to.be.lessThan(max)
    })

    it('should return a random integer within the range of 0 to max - 1', () => {
        const max = 100n
        const result = randomInt(max)

        expect(result).to.be.greaterThanOrEqual(0n)
        expect(result).to.be.lessThan(max)
    })

    it('should return a random integer within the range of 0 to 2^length - 1', () => {
        const max = 2n ** 32n
        const result = randomInt(max)

        expect(result).to.be.greaterThanOrEqual(0n)
        expect(result).to.be.lessThan(max)
    })
})
