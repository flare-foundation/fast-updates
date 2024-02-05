import { getWeb3 } from "../../src/utils/web3";
import { retry } from "../../src/utils/retry";
import { promisify } from "util";
import Web3 from "web3";
import { FTSOParameters } from "../config/FTSOParameters";

/**
 * This script is used to run a local simulation of the FTSO on the local hardhat network.
 * It deploys contracts and starts a cluster of data providers.
 */
export async function chainConfig(parameters: FTSOParameters) {
    const web3 = await retry(() => getWeb3(parameters.rpcUrl.toString()), 3, 1000);
    await setIntervalMining(web3, 5000);
}

/** Configures Hardhat to automatically mine blocks in the specified interval. */
export async function setIntervalMining(web3: Web3, interval: number = 1000) {
    await promisify((web3.currentProvider as any).send.bind(web3.currentProvider))({
        jsonrpc: "2.0",
        method: "evm_setAutomine",
        params: [false],
        id: new Date().getTime(),
    });

    await promisify((web3.currentProvider as any).send.bind(web3.currentProvider))({
        jsonrpc: "2.0",
        method: "evm_setIntervalMining",
        params: [interval],
        id: new Date().getTime(),
    });
}
