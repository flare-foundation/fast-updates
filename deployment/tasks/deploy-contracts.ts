import { writeFileSync } from "fs";
import { ContractAddresses, OUTPUT_FILE, loadAccounts, EPOCH_LEN } from "./common";
import { HardhatRuntimeEnvironment } from "hardhat/types";
import { getLogger } from "../../src/utils/logger";
import { syncTimeToNow } from "../../test-utils/utils/test-helpers";
import { RandInt } from "../../src/utils/rand";
import { RangeFPA, SampleFPA } from "../../src/utils/fixed-point-arithmetics";

const logger = getLogger("deploy-contracts");

// fast updater parameters
const ANCHOR_PRICES = [100, 1000, 10000, 100000, 1000000, 10000000];
const SUBMISSION_WINDOW = 10;
// incentive manager parameters
const BASE_SAMPLE_SIZE = SampleFPA(2);
const BASE_RANGE = RangeFPA(2 ** -5);
const SAMPLE_INCREASE_LIMIT = SampleFPA(5);
const RANGE_INCREASE_PRICE = 5;
const DURATION = 8;

export async function deployContracts(hre: HardhatRuntimeEnvironment) {
    await syncTimeToNow(hre);
    const accounts = loadAccounts(web3);
    const privateKey = accounts[0].privateKey;
    const artifacts = hre.artifacts;
    const governance = web3.eth.accounts.privateKeyToAccount(privateKey);

    // const voterRegistry = await artifacts.require("VoterRegistry").new(governance.address);
    const flareSystemMock = await artifacts.require("FlareSystemMock").new(RandInt(2n ** 256n - 1n), EPOCH_LEN);
    const fastUpdateIncentiveManager = await artifacts
        .require("FastUpdateIncentiveManager")
        .new(
            governance.address,
            governance.address,
            BASE_SAMPLE_SIZE,
            BASE_RANGE,
            SAMPLE_INCREASE_LIMIT,
            RANGE_INCREASE_PRICE,
            DURATION
        );

    const fastUpdater = await artifacts
        .require("FastUpdater")
        .new(
            governance.address,
            flareSystemMock.address,
            flareSystemMock.address,
            fastUpdateIncentiveManager.address,
            ANCHOR_PRICES,
            SUBMISSION_WINDOW
        );

    const deployed = <ContractAddresses>{
        voterRegistry: flareSystemMock.address,
        flareSystemManager: flareSystemMock.address,
        fastUpdateIncentiveManager: fastUpdateIncentiveManager.address,
        fastUpdater: fastUpdater.address,
    };

    outputAddresses(deployed);

    logger.info("Deployed all contracts");
    return deployed;
}

function outputAddresses(deployed: ContractAddresses) {
    const contents = JSON.stringify(deployed, null, 2);
    writeFileSync(OUTPUT_FILE, contents);
    logger.info(`Contract addresses written to ${OUTPUT_FILE}:\n${contents}`);
}
