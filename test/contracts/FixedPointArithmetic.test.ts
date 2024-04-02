import path from 'path'

import { expect } from 'chai'

import { randomInt } from '../../deployment/utils'
import type {
    TestFixedPointArithmeticContract,
    TestFixedPointArithmeticInstance,
} from '../../typechain-truffle/contracts/fastUpdates/test/TestFixedPointArithmetic'

const TestFixedPointArithmetic = artifacts.require(
    'TestFixedPointArithmetic'
) as TestFixedPointArithmeticContract

contract(
    `FixedPointArithmetic.sol; ${path.relative(path.resolve(), __filename)}`,
    (accounts) => {
        let fpaInstance: TestFixedPointArithmeticInstance
        before(async () => {
            const governance = accounts[0]
            if (!governance) {
                throw new Error('No governance account')
            }
            fpaInstance = await TestFixedPointArithmetic.new(
                governance as Truffle.TransactionDetails
            )
        })

        // Arithmetic identity tests

        it('should have one as additive one', async () => {
            const x = randomInt(2n ** 16n - 1n)
            const c = await fpaInstance.oneTest(x.toString())

            expect(c[0]).to.equal(x)
            expect(c[1]).to.equal(x)
        })
        it('should have zeroS as additive zero', async () => {
            const x = randomInt(2n ** 16n - 1n)
            const c = await fpaInstance.zeroSTest(x.toString())

            expect(c[0]).to.equal(x)
            expect(c[1]).to.equal(x)
        })
        it('should have zeroR as additive zero', async () => {
            const x = randomInt(2n ** 16n - 1n)
            const c = await fpaInstance.zeroRTest(x.toString())

            expect(c[0]).to.equal(x)
            expect(c[1]).to.equal(x)
        })

        // Addition/subtraction tests

        it('should add and subtract SampleSize values', async () => {
            let x = randomInt(2n ** 16n - 1n)
            let y = randomInt(2n ** 16n - 1n - x)
            if (x < y) {
                const z = x
                x = y
                y = z
            }

            const c = await fpaInstance.addSampleSizeTest(
                x.toString(),
                y.toString()
            )

            expect(c[0]).to.equal(x + y)
            expect(c[1]).to.equal(x - y)
        })

        it('should add and subtract Range values', async () => {
            let x = randomInt(2n ** 16n - 1n)
            let y = randomInt(2n ** 16n - 1n - x)
            if (x < y) {
                const z = x
                x = y
                y = z
            }

            const c = await fpaInstance.addRangeTest(x.toString(), y.toString())

            expect(c[0]).to.equal(x + y)
            expect(c[1]).to.equal(x - y)
        })

        it('should add and subtract Fee values', async () => {
            let x = randomInt(2n ** 240n - 1n)
            let y = randomInt(2n ** 240n - 1n - x)
            if (x < y) {
                const z = x
                x = y
                y = z
            }

            const c = await fpaInstance.addFeeTest(x.toString(), y.toString())

            expect(c[0]).to.equal(x + y)
            expect(c[1]).to.equal(x - y)
        })

        // Multiplication/division tests

        it('should multiply Scale values', async () => {
            // 47 is the largest exponent for a Number with a whole number of bytes
            // 45 is small enough to prevent overflow.
            const xS =
                '0x' +
                Math.floor(2 ** 47 + Math.random() * 2 ** 45).toString(16) +
                '0'.repeat(20)
            const yS =
                '0x' +
                Math.floor(2 ** 47 + Math.random() * 2 ** 45).toString(16) +
                '0'.repeat(20)

            const c = await fpaInstance.mulScaleTest(xS, yS)

            const x = BigInt(xS)
            const y = BigInt(yS)

            // TODO: These should not be compared as numbers, but as bigints
            expect(c.toString(16)).to.equal(((x * y) >> 127n).toString(16))
        })
        it('should multiply Price and Scale values', async () => {
            const xN = Math.floor(Math.random() * 2 ** 32)
            const yS =
                '0x' +
                Math.floor(2 ** 47 + Math.random() * 2 ** 45).toString(16) +
                '0'.repeat(20)

            const c = await fpaInstance.mulPriceScaleTest(xN, yS)

            const x = BigInt(xN)
            const y = BigInt(yS)

            expect(c.toString(16)).to.equal(((x * y) >> 127n).toString(16))
        })
        it('should multiply Fee and Range values', async () => {
            const xS =
                '0x' +
                Math.floor(Math.random() * 2 ** 39).toString(16) +
                '0'.repeat(20)
            const yS =
                '0x' +
                Math.floor(Math.random() * 2 ** 47).toString(16) +
                '0'.repeat(20)

            const c = await fpaInstance.mulFeeRangeTest(xS, yS)

            const x = BigInt(xS)
            const y = BigInt(yS)

            expect(c.toString(16)).to.equal(((x * y) >> 120n).toString(16))
        })
        it('should multiply Fractional and Fee values', async () => {
            const xS =
                '0x' +
                Math.floor(Math.random() * 2 ** 47).toString(16) +
                '0'.repeat(20)
            const yN = Math.floor(Math.random() * 2 ** 32)

            const c = await fpaInstance.mulFractionalFeeTest(xS, yN)

            const x = BigInt(xS)
            const y = BigInt(yN)

            expect(c.toString(16)).to.equal(((x * y) >> 128n).toString(16))
        })
        it('should multiply Fractional and SampleSize values', async () => {
            const xS =
                '0x' +
                Math.floor(Math.random() * 2 ** 47).toString(16) +
                '0'.repeat(20)
            const yS =
                '0x' +
                Math.floor(Math.random() * 2 ** 47).toString(16) +
                '0'.repeat(20)

            const c = await fpaInstance.mulFractionalSampleSizeTest(xS, yS)

            const x = BigInt(xS)
            const y = BigInt(yS)

            expect(c.toString(16)).to.equal(((x * y) >> 128n).toString(16))
        })
        it('should divide Range values', async () => {
            const y0 = Math.floor(Math.random() * 2 ** 47)
            const yS = '0x' + y0.toString(16) + '0'.repeat(20)
            const x0 = Math.floor(Math.random() * y0)
            const xS = '0x' + x0.toString(16) + '0'.repeat(20)

            const c = await fpaInstance.divRangeTest(xS, yS)

            const x = BigInt(xS)
            const y = BigInt(yS)

            expect(c.toString(16)).to.equal(((x << 128n) / y).toString(16))
        })
        it('should divide Fee values', async () => {
            const y0 = Math.floor(Math.random() * 2 ** 47)
            const yS = '0x' + y0.toString(16) + '0'.repeat(20)
            const x0 = Math.floor(Math.random() * y0)
            const xS = '0x' + x0.toString(16) + '0'.repeat(20)

            const c = await fpaInstance.divRangeTest(xS, yS)

            const x = BigInt(xS)
            const y = BigInt(yS)

            expect(c.toString(16)).to.equal(((x << 128n) / y).toString(16))
        })
        it('should divide Range and SampleSize values', async () => {
            const y0 = Math.floor(Math.random() * 2 ** 47)
            const yS = '0x' + y0.toString(16) + '0'.repeat(20)
            const x0 = Math.floor(Math.random() * y0)
            const xS = '0x' + x0.toString(16) + '0'.repeat(20)

            const c = await fpaInstance.divRangeTest(xS, yS)

            const x = BigInt(xS)
            const y = BigInt(yS)

            expect(c.toString(16)).to.equal(((x << 128n) / y).toString(16))
        })
        it('should divide Price and Scale values', async () => {
            const xN = Math.floor(Math.random() * 2 ** 32)
            const yS =
                '0x' +
                Math.floor(Math.random() * xN).toString(16) +
                '0'.repeat(20)

            const c = await fpaInstance.mulPriceScaleTest(xN, yS)

            const x = BigInt(xN)
            const y = BigInt(yS)

            expect(c.toString(16)).to.equal(((x * y) >> 127n).toString(16))
        })

        // Comparison and conversion tests

        it('should convert Precision to Scale', async () => {
            const xS =
                '0x' +
                Math.floor(Math.random() * 2 ** 46).toString(16) +
                '0'.repeat(20)

            const c = await fpaInstance.scaleWithPrecisionTest(xS)

            const x = BigInt(xS)

            expect(c.toString(16)).to.equal((x + (1n << 127n)).toString(16))
        })
        it('should compare Range values', async () => {
            const xI = Math.floor(Math.random() * 2 ** 16)
            const x = xI / 2 ** 8
            const yI = Math.floor(Math.random() * 2 ** 16)
            const y = yI / 2 ** 8

            const c = await fpaInstance.lessThanRangeTest(xI, yI)

            expect(c).to.equal(x < y)
        })
        it('should compare Fee values', async () => {
            const x = Math.floor(Math.random() * 2 ** 32)
            const y = Math.floor(Math.random() * 2 ** 32)

            const c = await fpaInstance.lessThanFeeTest(x, y)

            expect(c).to.equal(x < y)
        })
        it('should compare Range and SampleSize values', async () => {
            const xI = Math.floor(Math.random() * 2 ** 16)
            const x = xI / 2 ** 8
            const yI = Math.floor(Math.random() * 2 ** 16)
            const y = yI / 2 ** 8

            const c = await fpaInstance.lessThanRangeSampleSizeTest(xI, yI)

            expect(c).to.equal(x < y)
        })
    }
)
