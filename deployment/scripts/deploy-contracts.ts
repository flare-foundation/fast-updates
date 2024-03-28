import { writeFileSync } from 'fs'

import type { HardhatRuntimeEnvironment } from 'hardhat/types'

import type { FastUpdateIncentiveManagerContract } from '../../typechain-truffle/contracts/fastUpdates/implementation/FastUpdateIncentiveManager'
import type { FastUpdaterContract } from '../../typechain-truffle/contracts/fastUpdates/implementation/FastUpdater'
import type { FlareSystemMockContract } from '../../typechain-truffle/contracts/fastUpdates/test/FlareSystemMock'
import {
    ANCHOR_PRICES,
    BACKLOG_LEN,
    BASE_RANGE,
    BASE_SAMPLE_SIZE,
    DURATION,
    EPOCH_LEN,
    PATHS,
    RANGE_INCREASE_PRICE,
    SAMPLE_INCREASE_LIMIT,
    SUBMISSION_WINDOW,
} from '../config'
import type { FastUpdatesContractAddresses } from '../utils'
import {
    getOrCreateLogger,
    loadProviderAccount,
    randomInt,
    syncTimeToNow,
} from '../utils'

const logger = getOrCreateLogger('deploy-contracts')

/**
 * Deploys the contracts required for fast updates.
 * @param hre - The Hardhat runtime environment.
 * @returns A promise that resolves to the addresses of the deployed contracts.
 */
export async function deployContracts(
    hre: HardhatRuntimeEnvironment
): Promise<FastUpdatesContractAddresses> {
    await syncTimeToNow(hre)
    const artifacts = hre.artifacts

    const flareSystemMock = await (
        artifacts.require('FlareSystemMock') as FlareSystemMockContract
    ).new(randomInt(2n ** 256n - 1n).toString(), EPOCH_LEN)
    logger.info(
        `Deployed contract VoterRegistry/FlareSystemManager at ${flareSystemMock.address}`
    )

    const governance = loadProviderAccount(web3, 0, PATHS.accountsPath)
    const fastUpdateIncentiveManager = await (
        artifacts.require(
            'FastUpdateIncentiveManager'
        ) as FastUpdateIncentiveManagerContract
    ).new(
        governance.address,
        governance.address,
        BASE_SAMPLE_SIZE,
        BASE_RANGE,
        SAMPLE_INCREASE_LIMIT,
        RANGE_INCREASE_PRICE,
        DURATION
    )
    logger.info(
        `Deployed contract FastUpdateIncentiveManager at ${fastUpdateIncentiveManager.address}`
    )

    const fastUpdater = await (
        artifacts.require('FastUpdater') as FastUpdaterContract
    ).new(
        governance.address,
        flareSystemMock.address,
        flareSystemMock.address,
        fastUpdateIncentiveManager.address,
        ANCHOR_PRICES,
        SUBMISSION_WINDOW,
        BACKLOG_LEN
    )
    logger.info(`Deployed contract FastUpdater at ${fastUpdater.address}`)

    const deployedContracts: FastUpdatesContractAddresses = {
        voterRegistry: flareSystemMock.address,
        flareSystemManager: flareSystemMock.address,
        fastUpdateIncentiveManager: fastUpdateIncentiveManager.address,
        fastUpdater: fastUpdater.address,
    }

    writeDeployedContracts(deployedContracts)
    logger.info(`Finished deploying all contracts`)
    return deployedContracts
}

function writeDeployedContracts(deployed: FastUpdatesContractAddresses): void {
    writeFileSync(PATHS.contractsPath, JSON.stringify(deployed, null, 2))
    logger.info(`Contract addresses written to: ${PATHS.contractsPath}`)
}
