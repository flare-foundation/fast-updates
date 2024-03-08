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

        // Bit length tests

        it('should have 16 bit Scale values', async () => {
            const x = randomInt(2n ** 16n - 1n)
            const c1 = await fpaInstance.identityScaleTest(x.toString())

            expect(x).to.equal(c1)

            let c2
            try {
                c2 = await fpaInstance.identityScaleTest((2n ** 16n).toString())
            } catch (error) {
                expect(error).to.be.not.empty
            }
            expect(c2).to.be.undefined
        })
        it('should have 16 bit Precision values', async () => {
            const x = randomInt(2n ** 16n - 1n)
            const c1 = await fpaInstance.identityPrecisionTest(x.toString())

            expect(x).to.equal(c1)

            let c2
            try {
                c2 = await fpaInstance.identityPrecisionTest(
                    (2n ** 16n).toString()
                )
            } catch (error) {
                expect(error).to.be.not.empty
            }
            expect(c2).to.be.undefined
        })
        it('should have 16 bit SampleSize values', async () => {
            const x = randomInt(2n ** 16n - 1n)
            const c1 = await fpaInstance.identitySampleSizeTest(x.toString())

            expect(x).to.equal(c1)

            let c2
            try {
                c2 = await fpaInstance.identityPrecisionTest(
                    (2n ** 16n).toString()
                )
            } catch (error) {
                expect(error).to.be.not.empty
            }
            expect(c2).to.be.undefined
        })
        it('should have 16 bit Range values', async () => {
            const x = randomInt(2n ** 16n - 1n)
            const c1 = await fpaInstance.identitySampleSizeTest(x.toString())

            expect(x).to.equal(c1)

            let c2
            try {
                c2 = await fpaInstance.identitySampleSizeTest(
                    (2n ** 16n).toString()
                )
            } catch (error) {
                expect(error).to.be.not.empty
            }
            expect(c2).to.be.undefined
        })
        it('should have 32 bit Price values', async () => {
            const x = randomInt(2n ** 32n - 1n)
            const c1 = await fpaInstance.identityPriceTest(x.toString())

            expect(x).to.equal(c1)

            let c2
            try {
                c2 = await fpaInstance.identityPriceTest((2n ** 32n).toString())
            } catch (error) {
                expect(error).to.be.not.empty
            }
            expect(c2).to.be.undefined
        })
        it('should have 16 bit Fractional values', async () => {
            const x = randomInt(2n ** 16n - 1n)
            const c1 = await fpaInstance.identityFractionalTest(x.toString())

            expect(x).to.equal(c1)

            let c2
            try {
                c2 = await fpaInstance.identityFractionalTest(
                    (2n ** 16n).toString()
                )
            } catch (error) {
                expect(error).to.be.not.empty
            }
            expect(c2).to.be.undefined
        })
        it('should have 240 bit Fee values', async () => {
            const x = randomInt(2n ** 240n - 1n)
            const c1 = await fpaInstance.identityFeeTest(x.toString())

            expect(x).to.equal(c1)

            let c2
            try {
                c2 = await fpaInstance.identityFeeTest((2n ** 240n).toString())
            } catch (error) {
                expect(error).to.be.not.empty
            }
            expect(c2).to.be.undefined
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

        it('should multiply and divide Scale values', async () => {
            // Power is 13 here to prevent the product from overflowing
            const xI = Math.floor(2 ** 15 + Math.random() * 2 ** 13)
            const yI = Math.floor(2 ** 15 + Math.random() * 2 ** 13)

            const x = xI / 2 ** 15
            const y = yI / 2 ** 15

            const xy1 = Math.floor(x * y * 2 ** 15) / 2 ** 15
            const xy2 = Math.floor((x / y) * 2 ** 15) / 2 ** 15

            const c = await fpaInstance.mulScaleTest(xI, yI)

            // TODO: These should not be compared as numbers, but as bigints
            expect(c[0].toNumber() / 2 ** 15).to.equal(xy1)
            expect(c[1].toNumber() / 2 ** 15).to.equal(xy2)
        })
        it('should multiply Price and Scale values', async () => {
            const x = Math.floor(Math.random() * 2 ** 31)
            const yI = Math.floor(2 ** 15 + Math.random() * 2 ** 15)
            const y = yI / 2 ** 15

            const c = await fpaInstance.mulPriceScaleTest(x, yI)

            expect(c).to.equal(Math.floor(x * y))
        })
        it('should multiply Fee and Range values', async () => {
            const x = Math.floor(Math.random() * 2 ** 32)
            const yI = Math.floor(Math.random() * 2 ** 16)
            const y = yI / 2 ** 8

            const c = await fpaInstance.mulFeeRangeTest(x, yI)

            expect(c).to.equal(Math.floor(x * y))
        })
        it('should multiply Fractional and Fee values', async () => {
            const xI = Math.floor(Math.random() * 2 ** 16)
            const x = xI / 2 ** 16
            const y = Math.floor(Math.random() * 2 ** 32)

            const c = await fpaInstance.mulFractionalFeeTest(xI, y)

            expect(c).to.equal(Math.floor(x * y))
        })
        it('should multiply Fractional and SampleSize values', async () => {
            const xI = Math.floor(Math.random() * 2 ** 16)
            const x = xI / 2 ** 16
            const yI = Math.floor(Math.random() * 2 ** 16)
            const y = yI / 2 ** 8

            const xy = Math.floor(x * y * 2 ** 8) / 2 ** 8

            const c = await fpaInstance.mulFractionalSampleSizeTest(xI, yI)

            expect(c.toNumber() / 2 ** 8).to.equal(xy)
        })
        it('should divide Range values', async () => {
            const yI = Math.floor(Math.random() * 2 ** 16)
            const y = yI / 2 ** 8
            const xI = Math.floor(Math.random() * yI)
            const x = xI / 2 ** 8

            const xy = Math.floor((x / y) * 2 ** 16) / 2 ** 16

            const c = await fpaInstance.divRangeTest(xI, yI)

            expect(c.toNumber() / 2 ** 16).to.equal(xy)
        })
        it('should divide Fee values', async () => {
            const y = Math.floor(Math.random() * 2 ** 32)
            const x = Math.floor(Math.random() * y)

            const xy = Math.floor((x / y) * 2 ** 16) / 2 ** 16

            const c = await fpaInstance.divFeeTest(x, y)

            expect(c.toNumber() / 2 ** 16).to.equal(xy)
        })
        it('should divide Range and SampleSize values', async () => {
            const yI = Math.floor(Math.random() * 2 ** 16)
            const y = yI / 2 ** 8
            const xI = Math.floor(Math.random() * y)
            const x = xI / 2 ** 8

            const xy = Math.floor((x / y) * 2 ** 15) / 2 ** 15

            const c = await fpaInstance.divRangeSampleSizeTest(xI, yI)

            expect(c.toNumber() / 2 ** 15).to.equal(xy)
        })
        it('should divide Price and Scale values', async () => {
            const x = Math.floor(Math.random() * 2 ** 32)
            const yI = Math.floor(2 ** 15 + Math.random() * 2 ** 15)
            const y = yI / 2 ** 15

            const c = await fpaInstance.divPriceScaleTest(x, yI)

            expect(c).to.equal(Math.floor(x / y))
        })

        // Comparison and conversion tests

        it('should convert Precision to Scale', async () => {
            const xI = Math.floor(Math.random() * 2 ** 15)
            const x = xI / 2 ** 15

            const c = await fpaInstance.scaleWithPrecisionTest(xI)

            expect(c.toNumber() / 2 ** 15).to.equal(x + 1)
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
