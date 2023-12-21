import { readFileSync } from "fs";
import { Account } from "web3-core";

import Web3 from "web3";

export interface ContractAddresses {
    voterRegistry: string;
    fastUpdaters: string;
    fastUpdateIncentiveManager: string;
    fastUpdater: string;
}

export const OUTPUT_FILE = "./deployed-contracts.json";
export const TEST_ACCOUNT_FILE = "./deployment/config/test-11-accounts.json";

export function loadAccounts(web3: Web3): Account[] {
    return JSON.parse(readFileSync(TEST_ACCOUNT_FILE).toString()).map((x: any) =>
        web3.eth.accounts.privateKeyToAccount(x.privateKey)
    );
}
