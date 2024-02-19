import { readFileSync } from "fs";
import { Web3Provider } from "../../src/providers/Web3Provider";
import { loadFTSOParameters } from "../config/FTSOParameters";
import { FEEDS, EPOCH_LEN, ContractAddresses, OUTPUT_FILE, loadAccounts } from "../tasks/common";
import { getLogger, setGlobalLogFile } from "../../src/utils/logger";
import { getWeb3 } from "../../src/utils/web3";
import { PriceVoter } from "../../src/PriceVoter";
import { PriceFeedProvider } from "../../src/price-feeds/PriceFeedProvider";

const WEIGHT = 1000;
const logger = getLogger("price-voter");

async function main() {
    const myId = +process.argv[2];
    if (!myId) throw Error("Must provide a data provider id.");
    if (myId <= 0) throw Error("Data provider id must be greater than 0.");

    setGlobalLogFile(`price-voter-${myId}`);

    const parameters = loadFTSOParameters();
    const web3 = getWeb3(parameters.rpcUrl.toString());

    const contractAddresses = loadContracts();
    logger.info(`Initializing data provider ${myId}, connecting to ${parameters.rpcUrl}`);

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
    // console.log(`contract addresses`, contractAddresses);
    const web3Provider = await Web3Provider.create(contractAddresses, web3, parameters, privateKey);

    const priceFeedProvider = new PriceFeedProvider(FEEDS.length);

    const priceVoter = new PriceVoter(web3Provider, priceFeedProvider, address, EPOCH_LEN, WEIGHT, privateKey);

    await priceVoter.run();

    process.exit(0);
}

function loadContracts(): ContractAddresses {
    const parsed = JSON.parse(readFileSync(OUTPUT_FILE).toString());
    if (Object.entries(parsed).length == 0) throw Error(`No contract addresses found in ${OUTPUT_FILE}`);
    return parsed;
}

main().catch(e => {
    logger.error(`Price voter error, exiting ${e}`);
    getLogger("price-voter").error(e);
    process.exit(1);
});
