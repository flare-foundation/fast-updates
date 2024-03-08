import path from 'path'

import type { Web3Account } from 'web3-eth-accounts'

import { PATHS } from '../../deployment/config'
import { loadProviderAccounts } from '../../deployment/utils'
import type {
    FastUpdateIncentiveManagerContract,
    FastUpdateIncentiveManagerInstance,
} from '../../typechain-truffle/contracts/fastUpdates/implementation/FastUpdateIncentiveManager'

const FastUpdateIncentiveManager = artifacts.require(
    'FastUpdateIncentiveManager'
) as FastUpdateIncentiveManagerContract

const BASE_SAMPLE_SIZE = 5 * 2 ** 8 // 2^8 since scaled for 2^(-8) for fixed precision arithmetic
const BASE_RANGE = 2 * 2 ** 8
const SAMPLE_INCREASE_LIMIT = 5 * 2 ** 8
const RANGE_INCREASE_PRICE = 5
const DURATION = 8

contract(
    `FastUpdateIncentiveManager.sol; ${path.relative(path.resolve(), __filename)}`,
    () => {
        let fastUpdateIncentiveManager: FastUpdateIncentiveManagerInstance
        let accounts: Web3Account[]

        before(async () => {
            accounts = loadProviderAccounts(web3, PATHS.accountsPath)
            const governance = accounts[0]
            if (!governance) throw new Error('Governance account not found')

            fastUpdateIncentiveManager = await FastUpdateIncentiveManager.new(
                governance.address,
                governance.address,
                BASE_SAMPLE_SIZE,
                BASE_RANGE,
                SAMPLE_INCREASE_LIMIT,
                RANGE_INCREASE_PRICE,
                DURATION
            )
        })

        it('should get expected sample size', async () => {
            const sampleSize =
                await fastUpdateIncentiveManager.getExpectedSampleSize()
            expect(sampleSize).to.equal(BASE_SAMPLE_SIZE)
        })

        it('should get range', async () => {
            const range = await fastUpdateIncentiveManager.getRange()
            expect(range).to.equal(BASE_RANGE)
        })

        it('should get precision', async () => {
            const precision = await fastUpdateIncentiveManager.getPrecision()
            // precision scaled for 2^(-15)
            expect(precision).to.equal(
                Math.floor((BASE_RANGE / BASE_SAMPLE_SIZE) * 2 ** 15)
            )
        })

        it('should get scale', async () => {
            const scale = await fastUpdateIncentiveManager.getScale()
            expect(scale).to.equal(
                Math.floor(2 ** 15 + (BASE_RANGE / BASE_SAMPLE_SIZE) * 2 ** 15)
            )
        })

        it('should offer incentive', async () => {
            const rangeIncrease = BASE_RANGE
            const rangeLimit = 4 * 2 ** 8
            const offer = {
                rangeIncrease: rangeIncrease.toString(),
                rangeLimit: rangeLimit.toString(),
            }
            if (!accounts[1]) throw new Error('Account not found')
            await fastUpdateIncentiveManager.offerIncentive(offer, {
                from: accounts[1].address,
                value: '100000',
            })

            const newRange = (
                await fastUpdateIncentiveManager.getRange()
            ).toNumber()
            expect(newRange).to.equal(BASE_RANGE * 2)

            const newSampleSize = (
                await fastUpdateIncentiveManager.getExpectedSampleSize()
            ).toNumber()
            expect(newSampleSize).to.equal(BASE_SAMPLE_SIZE * 2 - 1)

            const precision = await fastUpdateIncentiveManager.getPrecision()
            expect(precision).to.equal(
                Math.floor((newRange / newSampleSize) * 2 ** 15)
            )

            const scale = await fastUpdateIncentiveManager.getScale()
            expect(scale).to.equal(
                Math.floor(2 ** 15 + (newRange / newSampleSize) * 2 ** 15)
            )
        })

        it('should change incentive duration', async () => {
            const incentiveDuration =
                await fastUpdateIncentiveManager.getIncentiveDuration()
            expect(incentiveDuration.toString()).to.equal(DURATION.toString())

            if (!accounts[0]) throw new Error('Account not found')
            await fastUpdateIncentiveManager.setIncentiveDuration('10', {
                from: accounts[0].address,
            })

            const newIncentiveDuration =
                await fastUpdateIncentiveManager.getIncentiveDuration()

            expect(newIncentiveDuration.toString()).to.equal('10')
        })
    }
)
