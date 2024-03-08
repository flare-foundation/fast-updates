import type { Web3Account } from 'web3-eth-accounts'

import { EPOCH_LEN, FEEDS, PATHS, WEIGHT } from '../deployment/config'

import { ExamplePriceFeedProvider } from './providers/ExamplePriceFeedProvider'
import { FastUpdatesProvider } from './providers/FastUpdatesProvider'
import { Web3Provider } from './providers/Web3Provider'
import {
    createWeb3Instance,
    getOrCreateLogger,
    loadFastUpdatesContractAddresses,
    loadNetworkParameters,
    loadProviderAccount,
} from './utils'

/**
 * The main function that serves as the entry point of the script.
 * It extracts the provider id from the command line arguments,
 * loads the FTSO parameters and contracts, initializes the data provider,
 * registers as a voter for the next epoch, and runs the price voter.
 */
async function main(provider?: number): Promise<void> {
    // Extract the provider id from the command line arguments
    let providerId: number

    if (!provider) providerId = +(process.argv[2] as unknown as number)
    else providerId = provider

    if (!providerId) throw Error('Must provide a data provider id.')
    if (providerId <= 0) throw Error('Data provider id must be greater than 0.')

    const logger = getOrCreateLogger(
        'fast-update-provider-' + providerId,
        'fast-update-provider-' + providerId
    )

    // Load the FTSO parameters and contracts from the config file
    const networkParams = loadNetworkParameters(PATHS.configPath)
    const fastUpdatesContractAddresses = loadFastUpdatesContractAddresses(
        PATHS.contractsPath
    )
    const web3 = createWeb3Instance(networkParams.rpcUrl)
    logger.info(
        `Initializing fast updates provider ${providerId}, connecting to ${networkParams.rpcUrl}`
    )
    let providerAccount: Web3Account
    if (process.env['DATA_PROVIDER_VOTING_KEY']) {
        logger.info(`Using environment variable for private key`)
        const pk = process.env['DATA_PROVIDER_VOTING_KEY']
        providerAccount = {
            address: web3.eth.accounts.privateKeyToAccount(pk).address,
            privateKey: pk,
        } as Web3Account
    } else {
        logger.info(`Loading private key from file: ${PATHS.accountsPath}`)
        providerAccount = loadProviderAccount(
            web3,
            providerId,
            PATHS.accountsPath
        )
    }

    // Create the data provider instance
    const web3Provider = await Web3Provider.create(
        web3,
        fastUpdatesContractAddresses,
        networkParams,
        providerAccount
    )

    const precision = Number(
        await web3Provider.contracts.fastUpdateIncentiveManager.methods
            .getPrecision()
            .call()
    )
    const genesisBlockSec = await web3Provider.getGenesisTimeSec()
    const priceFeedProvider = new ExamplePriceFeedProvider(
        FEEDS.length,
        precision,
        genesisBlockSec
    )

    // Attach the price feed provider and data provider into a fast updates provider
    const fastUpdatesProvider = new FastUpdatesProvider(
        web3Provider,
        priceFeedProvider,
        EPOCH_LEN,
        WEIGHT,
        providerAccount.address,
        providerAccount.privateKey,
        logger
    )

    // Initialize the fast updates provider (will only start at new reward epoch)
    await fastUpdatesProvider.init()

    // Run the fast updates provider
    await fastUpdatesProvider.run()

    process.exit(0)
}

main().catch((error) => {
    console.error(error)
    process.exit(1)
})
