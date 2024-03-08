import { readFileSync } from 'fs'

import dotenv from 'dotenv'
import type { HardhatNetworkAccountUserConfig } from 'hardhat/types/config'

import { PATHS } from './deployment/config'

dotenv.config()

/**
 * Loads the test accounts for the application.
 *
 * @returns An array of HardhatNetworkAccountUserConfig objects representing the test accounts.
 */
export default function loadTestAccounts(): HardhatNetworkAccountUserConfig[] {
    const testAccounts: HardhatNetworkAccountUserConfig[] = []

    // Add deployer account as the first account
    if (process.env['DEPLOYER_PRIVATE_KEY']) {
        testAccounts.push({
            privateKey: process.env['DEPLOYER_PRIVATE_KEY'],
            balance: '100000000000000000000000000000000',
        })
    }
    // Add governance settings deployer account as the second account
    if (process.env['GOVERNANCE_SETTINGS_DEPLOYER_PRIVATE_KEY']) {
        testAccounts.push({
            privateKey: process.env['GOVERNANCE_SETTINGS_DEPLOYER_PRIVATE_KEY'],
            balance: '100000000000000000000000000000000',
        })
    }
    // Add list of accounts from json
    testAccounts.push(
        ...(
            JSON.parse(
                readFileSync(PATHS.accountsPath).toString()
            ) as HardhatNetworkAccountUserConfig[]
        )
            .slice(0, process.env['TENDERLY'] == 'true' ? 150 : 2000)
            .filter(
                (x: HardhatNetworkAccountUserConfig) =>
                    x.privateKey != process.env['DEPLOYER_PRIVATE_KEY']
            )
    )
    // Add genesis governance account as second last account
    if (process.env['GENESIS_GOVERNANCE_PRIVATE_KEY']) {
        testAccounts.push({
            privateKey: process.env['GENESIS_GOVERNANCE_PRIVATE_KEY'],
            balance: '100000000000000000000000000000000',
        })
    }
    // Add governance account as last account
    if (process.env['GOVERNANCE_PRIVATE_KEY']) {
        testAccounts.push({
            privateKey: process.env['GOVERNANCE_PRIVATE_KEY'],
            balance: '100000000000000000000000000000000',
        })
    }

    return testAccounts
}
