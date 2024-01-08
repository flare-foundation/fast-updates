import { readFileSync } from "fs";
import { FTSOClient } from "../../src/FTSOClient";
import { Web3Provider } from "../../src/providers/Web3Provider";
import { loadFTSOParameters } from "../config/FTSOParameters";
import { ContractAddresses, OUTPUT_FILE, loadAccounts } from "../tasks/common";
import { getLogger, setGlobalLogFile } from "../../src/utils/logger";
import { getWeb3 } from "../../src/utils/web3";
import { PriceVoter } from "../../src/PriceVoter";
import { PriceFeedProvider } from "../../src/price-feeds/PriceFeedProvider";

// const TEST_EPOCH = 1;
const EPOCH_LEN = 10;
const WEIGHT = 1000;
const FEEDS = [0, 1, 2, 3, 4, 5];

async function main() {
    const myId = +process.argv[2];
    if (!myId) throw Error("Must provide a data provider id.");
    if (myId <= 0) throw Error("Data provider id must be greater than 0.");

    setGlobalLogFile(`price-voter-${myId}`);

    const parameters = loadFTSOParameters();
    const web3 = getWeb3(parameters.rpcUrl.toString());

    const contractAddresses = loadContracts();
    getLogger("price-voter").info(`Initializing data provider ${myId}, connecting to ${parameters.rpcUrl}`);

    let privateKey: string;
    let address: string;
    if (process.env.DATA_PROVIDER_VOTING_KEY != undefined) {
        privateKey = process.env.DATA_PROVIDER_VOTING_KEY;
        address = web3.eth.accounts.privateKeyToAccount(privateKey).address;
    } else {
        const accounts = loadAccounts(web3);
        privateKey = accounts[myId].privateKey;
        address = accounts[myId].address;
    }
    const web3Provider = await Web3Provider.create(contractAddresses, web3, parameters, privateKey);

    const priceFeedProvider = new PriceFeedProvider(6);

    const priceVoter = new PriceVoter(web3Provider, priceFeedProvider, address, EPOCH_LEN, WEIGHT);
    const currentBlock = await web3Provider.getBlockNumber();
    const nextEpoch = Math.floor((currentBlock + 1) / EPOCH_LEN) + 1;

    const receipt1 = priceVoter.registerAsVoter(nextEpoch, WEIGHT);
    const receipt2 = priceVoter.registerAsProvider(1);

    let receipt = await receipt1;
    console.log(
        "registered as a voter for epoch",
        nextEpoch,
        "in block",
        receipt.blockNumber,
        "status",
        receipt.status
    );
    receipt = await receipt2;
    console.log("registered as a provider of Fast Updates", "in block", receipt.blockNumber, "status", receipt.status);

    await priceVoter.run(FEEDS);

    process.exit(0);
}

function loadContracts(): ContractAddresses {
    const parsed = JSON.parse(readFileSync(OUTPUT_FILE).toString());
    if (Object.entries(parsed).length == 0) throw Error(`No contract addresses found in ${OUTPUT_FILE}`);
    return parsed;
}

main().catch(e => {
    console.error("Price voter error, exiting", e);
    getLogger("price-voter").error(e);
    process.exit(1);
});
