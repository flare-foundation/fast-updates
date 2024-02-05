import { readFileSync } from "fs";

import { HardhatRuntimeEnvironment } from "hardhat/types";
import { EPOCH_LEN, FEEDS, OUTPUT_FILE, loadAccounts } from "./common";
import { FTSOParameters } from "../config/FTSOParameters";
import { getLogger } from "../../src/utils/logger";
import { Web3Provider } from "../../src/providers/Web3Provider";
import { getWeb3 } from "../../src/utils/web3";

const logger = getLogger("admin-daemon");

function loadContracts() {
    return JSON.parse(readFileSync(OUTPUT_FILE).toString());
}

export async function runAdminDaemon(hre: HardhatRuntimeEnvironment, parameters: FTSOParameters) {
    const web3 = getWeb3(parameters.rpcUrl.toString());

    const contractAddresses = loadContracts();

    let privateKey: string;
    if (process.env.DATA_PROVIDER_VOTING_KEY != undefined) {
        privateKey = process.env.DATA_PROVIDER_VOTING_KEY;
    } else {
        const accounts = loadAccounts(web3);
        privateKey = accounts[0].privateKey;
    }
    const web3Provider = await Web3Provider.create(contractAddresses, web3, parameters, privateKey);

    for (;;) {
        try {
            const prices: string[] = await web3Provider.fetchCurrentPrices(FEEDS);
            const blockNum = await web3Provider.getBlockNumber();
            logger.info(`prices in block ${blockNum}: ${prices}`);

            web3Provider.advanceIncentive().catch(error => {
                logger.error(`failed to advance incentive ${error}`);
            });
            web3Provider.freeSubmitted(1).catch(error => {
                logger.error(`failed to free submitted ${error}`);
            });

            await web3Provider.waitForBlock(blockNum + 1);
        } catch (e) {
            logger.error(e);
        }
    }
}
