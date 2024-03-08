import { expect } from 'chai'
import { describe, it } from 'mocha'
import Web3 from 'web3'

import {
    loadNetworkParameters,
    loadProviderAccounts,
} from '../../client/utils/loader'
import { PATHS } from '../../deployment/config'

describe('Loader Utils', () => {
    describe('loadNetworkParameters', () => {
        it('should load FTSO parameters from a file', () => {
            const params = loadNetworkParameters(PATHS.configPath)

            expect(params).to.have.property('rpcUrl')
            expect(params).to.have.property('gasLimit')
            expect(params).to.have.property('gasPriceMultiplier')
        })
    })

    describe('loadProviderAccounts', () => {
        it('should load provider accounts from a file and convert them to Web3 accounts', () => {
            const web3 = new Web3()
            const accounts = loadProviderAccounts(web3, PATHS.accountsPath)

            expect(accounts).to.be.an('array')
            expect(accounts).to.have.lengthOf.at.least(1)
            expect(accounts[0]).to.have.property('address')
            expect(accounts[0]).to.have.property('privateKey')
        })
    })
})
