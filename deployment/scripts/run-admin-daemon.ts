import { readFileSync } from 'fs'

import type { Web3Account } from 'web3-eth-accounts'
import type { TransactionReceipt } from 'web3-types'

import { ExamplePriceFeedProvider } from '../../client/providers/ExamplePriceFeedProvider'
import { Web3Provider } from '../../client/providers/Web3Provider'
import { FEEDS, PATHS } from '../config'
import type { FastUpdatesContractAddresses, NetworkParameters } from '../utils'
import {
    createWeb3Instance,
    getOrCreateLogger,
    loadProviderAccount,
} from '../utils'

/**
 * Runs the admin daemon.
 *
 * @param parameters - The network parameters.
 * @returns A Promise that resolves to void.
 */
export async function runAdminDaemon(
    parameters: NetworkParameters
): Promise<void> {
    const logger = getOrCreateLogger('admin-daemon', 'admin-daemon')

    const web3 = createWeb3Instance(parameters.rpcUrl, logger)
    const contractAddresses = JSON.parse(
        readFileSync(PATHS.contractsPath).toString()
    ) as FastUpdatesContractAddresses
    let account: Web3Account
    if (process.env['DATA_PROVIDER_VOTING_KEY'] != undefined) {
        const privateKey = process.env['DATA_PROVIDER_VOTING_KEY']
        account = web3.eth.accounts.privateKeyToAccount(privateKey)
    } else {
        account = loadProviderAccount(web3, 0, PATHS.accountsPath)
    }
    const web3Provider = await Web3Provider.create(
        web3,
        contractAddresses,
        parameters,
        account
    )

    const genesisBlockSec = await web3Provider.getGenesisTimeSec()
    const priceFeedProvider = new ExamplePriceFeedProvider(
        FEEDS.length,
        0,
        genesisBlockSec
    )

    for (;;) {
        try {
            const onChainPrices: string[] =
                await web3Provider.fetchCurrentPrices(Array.from(FEEDS))
            const offChainPrices = priceFeedProvider.getCurrentPrices(
                Array.from(FEEDS)
            )
            const blockNum = Number(await web3.eth.getBlockNumber())

            logger.info(
                `Feeds ${FEEDS.length}, Block: ${blockNum}, onChainPrices: ${onChainPrices.join(', ')}, offChainPrices: ${offChainPrices.join(', ')}`
            )

            const promises = []
            promises.push(
                web3Provider
                    .advanceIncentive()
                    .then((receipt: TransactionReceipt) => {
                        logger.info(
                            `AdvanceIncentive successful at block ${receipt.blockNumber}`
                        )
                    })
                    .catch((error: unknown) => {
                        logger.error(`(advanceIncentive), ${error as string}`)
                    })
            )
            promises.push(
                web3Provider
                    .freeSubmitted(1)
                    .then((receipt: TransactionReceipt) => {
                        logger.info(
                            `FreeSubmitted successful at block ${receipt.blockNumber}`
                        )
                    })
                    .catch((error: unknown) => {
                        logger.error(`(freeSubmitted), ${error as string}`)
                    })
            )
            promises.push(
                web3Provider
                    .applySubmitted(2)
                    .then((receipt: TransactionReceipt) => {
                        logger.info(
                            `ApplySubmitted successful at block ${receipt.blockNumber}`
                        )
                    })
                    .catch((error: unknown) => {
                        logger.error(`(applySubmitted), ${error as string}`)
                    })
            )
            promises.push(web3Provider.waitForBlock(blockNum + 1))
            await Promise.all(promises)
        } catch (e) {
            logger.error(e)
        }
    }
}
