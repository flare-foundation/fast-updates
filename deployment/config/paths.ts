import path from 'path'

export type PathConfig = {
    readonly configPath: string
    readonly accountsPath: string
    readonly contractsPath: string
}

const BASE_PATH = path.join(path.resolve(), 'deployment', 'config')

const TEST_PATHS = {
    configPath: path.join(BASE_PATH, 'network-config', 'config-local.json'),
    accountsPath: path.join(
        BASE_PATH,
        'account-config',
        'test-11-accounts.json'
    ),
    contractsPath: path.join(
        BASE_PATH,
        'deployed-contracts',
        'deployed-contracts.json'
    ),
}

const pathConfig = (): PathConfig => {
    const network = process.env['NETWORK'] as
        | 'local-test'
        | 'docker'
        | 'from-env'
        | 'coston2'

    switch (network) {
        case 'local-test':
        case 'docker':
            return {
                configPath: path.join(
                    BASE_PATH,
                    'network-config',
                    'config-docker.json'
                ),
                accountsPath: path.join(
                    BASE_PATH,
                    'account-config',
                    'test-11-accounts.json'
                ),
                contractsPath: path.join(
                    BASE_PATH,
                    'deployed-contracts',
                    'deployed-contracts.json'
                ),
            }
        case 'coston2':
            return {
                configPath: path.join(
                    BASE_PATH,
                    'network-config',
                    'config-coston2.json'
                ),
                accountsPath: path.join(
                    BASE_PATH,
                    'account-config',
                    'test-11-accounts.json'
                ),
                contractsPath: path.join(
                    BASE_PATH,
                    'deployed-contracts',
                    'deployed-contracts.json'
                ),
            }
        case 'from-env':
            return {
                configPath: process.env['CONFIG_PATH'] ?? '',
                accountsPath: process.env['ACCOUNTS_PATH'] ?? '',
                contractsPath: process.env['CONTRACTS_PATH'] ?? '',
            }
        default:
            return TEST_PATHS
    }
}

export const PATHS = pathConfig()
