import { readFileSync } from 'fs'

import type Web3 from 'web3'
import type { Web3Account } from 'web3-eth-accounts'

import { getOrCreateLogger } from './logger'

export type NetworkParameters = {
    readonly rpcUrl: string
    readonly gasLimit: string
    readonly gasPriceMultiplier: number
}

type ProviderAccount = {
    readonly privateKey: string
    readonly balance: string
}

export type FastUpdatesContractAddresses = {
    readonly voterRegistry: string
    readonly flareSystemManager: string
    readonly fastUpdateIncentiveManager: string
    readonly fastUpdater: string
}

const logger = getOrCreateLogger('loader')

/**
 * Loads FTSO parameters from a file.
 *
 * @param filepath - The path to the file containing the parameters.
 * @returns The loaded FTSO parameters.
 */
export function loadNetworkParameters(filepath: string): NetworkParameters {
    logger.info(`Loading FTSO parameters from ${filepath}`)
    const params = JSON.parse(
        readFileSync(filepath).toString()
    ) as NetworkParameters

    logger.info(
        `Loaded FTSO parameters: gasLimit=${params.gasLimit}, gasPriceMultiplier=${params.gasPriceMultiplier}, rpcUrl=${params.rpcUrl}`
    )
    return params
}

/**
 * Loads the contract addresses from the specified file path.
 * @param filepath - The path to the file containing the contract addresses.
 * @returns The parsed contract addresses.
 * @throws Error if no contract addresses are found in the file.
 */
export function loadFastUpdatesContractAddresses(
    filepath: string
): FastUpdatesContractAddresses {
    const parsed = JSON.parse(
        readFileSync(filepath).toString()
    ) as FastUpdatesContractAddresses
    if (Object.entries(parsed).length == 0)
        throw Error(`No contract addresses found in ${filepath}`)
    return parsed
}

/**
 * Loads provider accounts from a file and converts them to Web3 accounts.
 * @param filepath - The path to the file containing the provider accounts.
 * @param web3 - The Web3 instance.
 * @returns An array of converted provider accounts.
 */
export function loadProviderAccounts(
    web3: Web3,
    filepath: string
): Web3Account[] {
    return (
        JSON.parse(readFileSync(filepath).toString()) as ProviderAccount[]
    ).map((x: ProviderAccount) =>
        web3.eth.accounts.privateKeyToAccount(x.privateKey)
    )
}

/**
 * Loads the provider account based on the given provider ID.
 *
 * @param web3 - The Web3 instance.
 * @param providerId - The ID of the provider.
 * @param filepath - The optional filepath to load provider accounts from.
 * @returns The loaded account.
 * @throws Error if no account is found for the given provider ID.
 */
export function loadProviderAccount(
    web3: Web3,
    providerId: number,
    filepath: string
): Web3Account {
    const accounts = loadProviderAccounts(web3, filepath)
    const account = accounts[providerId]
    if (!account) {
        throw new Error(`No account found for provider id ${providerId}`)
    }
    return account
}
