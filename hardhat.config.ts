/* eslint-disable @typescript-eslint/no-unused-vars */
import '@nomicfoundation/hardhat-chai-matchers'
import '@nomiclabs/hardhat-truffle5'
import '@nomiclabs/hardhat-web3'
import '@nomicfoundation/hardhat-web3-v4'
import 'solidity-coverage'

import dotenv from 'dotenv'
import type { HardhatUserConfig } from 'hardhat/config'
import { task } from 'hardhat/config'
import type { HardhatNetworkAccountUserConfig } from 'hardhat/types/config'

import { PATHS } from './deployment/config'
import { deployContracts } from './deployment/scripts/deploy-contracts'
import { runAdminDaemon } from './deployment/scripts/run-admin-daemon'
import { loadNetworkParameters } from './deployment/utils'
import loadTestAccounts from './hardhat.utils'

dotenv.config()

task('deploy-contracts', `Deploy contracts to the network`).setAction(
    async (_args, hre, _runSuper) => {
        await deployContracts(hre)
    }
)

task('run-admin-daemon', `Does admin tasks`).setAction(
    async (_args, _hre, _runSuper) => {
        const parameters = loadNetworkParameters(PATHS.configPath)
        await runAdminDaemon(parameters)
    }
)

const accounts: HardhatNetworkAccountUserConfig[] = loadTestAccounts()
const privateKeys = accounts.map(
    (x: HardhatNetworkAccountUserConfig) => x.privateKey
)

const config: HardhatUserConfig = {
    solidity: {
        version: '0.8.18',
        settings: {
            evmVersion: 'london',
            optimizer: {
                enabled: true,
                runs: 200,
            },
        },
    },
    mocha: {
        timeout: 100000000,
    },
    defaultNetwork: 'hardhat',
    networks: {
        hardhat: {
            accounts,
            blockGasLimit: 8000000,
            /**
             * Normally each Truffle smart contract interaction that modifies state results
             * in a transaction mined in a new block with a +1s block timestamp.
             * This is problematic because we need perform multiple smart contract actions
             * in the same price epoch, and the block timestamps end up not fitting into an epoch duration,
             * causing test failures.
             *
             * Enabling consecutive blocks with the same timestamp is not perfect,
             * but it alleviates this problem. A better solution would be manual mining and packing
             * multiple e.g. setup transactions into a single block with a controlled timestamp, but that
             * would make test code more complex and seems to be not very well supported by Truffle.
             */
            allowBlocksWithSameTimestamp: true,
            mining: {
                auto: false,
                interval: 5000,
            },
        },
        scdev: {
            url: 'http://127.0.0.1:9650/ext/bc/C/rpc',
            timeout: 40000,
            accounts: privateKeys,
        },
        staging: {
            url: 'http://127.0.0.1:9650/ext/bc/C/rpc',
            timeout: 40000,
            accounts: privateKeys,
        },
        songbird: {
            url: 'https://songbird-api.flare.network/ext/C/rpc',
            timeout: 40000,
            accounts: privateKeys,
        },
        flare: {
            url: 'https://flare-api.flare.network/ext/C/rpc',
            timeout: 40000,
            accounts: privateKeys,
        },
        coston: {
            url: 'https://coston-api.flare.network/ext/C/rpc',
            timeout: 40000,
            accounts: privateKeys,
        },
        coston2: {
            url: 'https://coston2-api.flare.network/ext/C/rpc',
            timeout: 40000,
            accounts: privateKeys,
        },
        local: {
            url: 'http://127.0.0.1:8545',
            chainId: 31337,
        },
        docker: {
            url: 'http://chain:8545',
            chainId: 31337,
        },
    },
    paths: {
        sources: './contracts',
        tests: './test',
        cache: './cache',
        artifacts: './artifacts',
    },
}

export default config
