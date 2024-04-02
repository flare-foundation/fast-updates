import path from 'path'

import { bn254 } from '@noble/curves/bn254'
import { expect } from 'chai'

import { randomInt } from '../../client/utils'
import type {
    TestBn256Contract,
    TestBn256Instance,
} from '../../typechain-truffle/contracts/fastUpdates/test/TestBn256'

const TestBn256 = artifacts.require('TestBn256') as TestBn256Contract

contract(
    `Bn256.sol; ${path.relative(path.resolve(), __filename)}`,
    (accounts) => {
        let bn256Instance: TestBn256Instance
        before(async () => {
            const governance = accounts[0]
            if (!governance) throw new Error('No governance account')
            bn256Instance = await TestBn256.new(
                governance as Truffle.TransactionDetails
            )
        })

        it('should add two points', async () => {
            const r1 = randomInt(bn254.CURVE.n)
            const r2 = randomInt(bn254.CURVE.n)
            const a = bn254.ProjectivePoint.BASE.multiply(r1)
            const b = bn254.ProjectivePoint.BASE.multiply(r2)

            const c = await bn256Instance.publicG1Add(
                {
                    x: a.x.toString(),
                    y: a.y.toString(),
                },
                {
                    x: b.x.toString(),
                    y: b.y.toString(),
                }
            )

            const cCheck = a.add(b)
            expect(c.x.toString()).to.equal(cCheck.x.toString())
            expect(c.y.toString()).to.equal(cCheck.y.toString())
        })

        it('should multiply a point with a scalar', async () => {
            const r1 = randomInt(bn254.CURVE.n)
            const r2 = randomInt(bn254.CURVE.n)
            const a = bn254.ProjectivePoint.BASE.multiply(r1)

            const c = await bn256Instance.publicG1ScalarMultiply(
                { x: a.x.toString(), y: a.y.toString() },
                r2.toString()
            )

            const cCheck = a.multiply(r2)
            expect(c.x.toString()).to.equal(cCheck.x.toString())
            expect(c.y.toString()).to.equal(cCheck.y.toString())
        })
    }
)
