import { createConnection } from 'net'
import type { TcpNetConnectOpts } from 'net'
import { URL } from 'url'

import glob from 'glob'
import Web3 from 'web3'
import type { Logger } from 'winston'

export interface BareSignature {
    readonly v: number
    readonly r: string
    readonly s: string
}

/**
 * Creates and returns a Web3 instance with the specified RPC link.
 * @param rpcLink The RPC link to connect to.
 * @param logger (Optional) The logger to use for logging WebSocket connection close events.
 * @returns A Web3 instance.
 */
export function createWeb3Instance(rpcLink: string, logger?: Logger): Web3 {
    let provider
    const url = new URL(rpcLink)
    if (logger) {
        logger.info(`Attempting to connect to provider: ${url.href}`)
    }
    if (url.protocol === 'http:' || url.protocol === 'https:') {
        provider = new Web3.providers.HttpProvider(url.href)
        logger?.info(`Connected to HTTP provider: ${url.href}`)
    } else if (url.protocol === 'ws:' || url.protocol === 'wss:') {
        const netConnectOpts: TcpNetConnectOpts = {
            port: Number(url.port),
            keepAlive: true,
            keepAliveInitialDelay: 60000,
        }
        const socket = createConnection(netConnectOpts)
        provider = new Web3.providers.WebsocketProvider(url.href, socket)
        provider.on('open', () => {
            if (logger) {
                logger.info(`Connected to RPC provider: ${url.href}`)
            }
        })
        provider.on('close', () => {
            if (logger) {
                logger.error(`Closed RPC provider connection`)
            }
        })
        logger?.info(`Connected to WebSocket provider: ${url.href}`)
    }
    const web3 = new Web3(provider)
    return web3
}

/**
 * Loads a contract of type `ContractType` using the provided `web3` instance, `address`, and `name`.
 *
 * @param web3 The Web3 instance used to interact with the blockchain.
 * @param address The address of the contract.
 * @param name The name of the contract.
 * @returns A promise that resolves to the loaded contract of type `ContractType`.
 * @throws Error if the `address` is not provided.
 */
export async function loadContract<ContractType>(
    web3: Web3,
    address: string,
    name: string
): Promise<ContractType> {
    if (!address) throw Error(`Address for ${name} not provided`)
    const abiPath = await relativeContractABIPathForContractName(name)
    const { abi } = require(`../../artifacts/${abiPath}`) // eslint-disable-line @typescript-eslint/no-var-requires,@typescript-eslint/no-unsafe-assignment
    const contract = new web3.eth.Contract(abi, address)
    return contract as ContractType
}

/**
 * Retrieves the relative path of the contract ABI file for a given contract name.
 * @param name - The name of the contract.
 * @param artifactsRoot - The root directory where the contract artifacts are stored. Default is 'artifacts'.
 * @returns A promise that resolves to the relative path of the contract ABI file.
 * @throws If no files are found for the contract.
 */
async function relativeContractABIPathForContractName(
    name: string,
    artifactsRoot = 'artifacts'
): Promise<string> {
    return new Promise((resolve, reject) => {
        glob(
            `contracts/**/${name}.sol/${name}.json`,
            { cwd: artifactsRoot },
            (er: Error | null, files: string[] | null) => {
                if (er) {
                    reject(er)
                } else {
                    if (files && files.length === 1) {
                        resolve(files[0] as string)
                    } else {
                        reject(Error(`No files found for contract ${name}`))
                    }
                }
            }
        )
    })
}

/**
 * Signs a message using the provided private key and returns the signature.
 * @param web3 - The Web3 instance.
 * @param message - The message to be signed.
 * @param privateKey - The private key used for signing.
 * @returns The signature object containing the v, r, and s values.
 */
export function signMessage(
    web3: Web3,
    message: string,
    privateKey: string
): BareSignature {
    const signature = web3.eth.accounts.sign(message, privateKey)
    return {
        v: parseInt(signature.v, 16),
        r: signature.r,
        s: signature.s,
    }
}
