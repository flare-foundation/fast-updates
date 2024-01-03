import { FTSOParameters } from "../config/FTSOParameters";
import { writeFileSync } from "fs";
import { ContractAddresses, OUTPUT_FILE } from "./common";
import { HardhatRuntimeEnvironment } from "hardhat/types";
import { getLogger } from "../../src/utils/logger";
import { syncTimeToNow } from "../../test-utils/utils/test-helpers";

const logger = getLogger("deploy-contracts");

// TODO: extract constants to config
// fast updater parameters
const ANCHOR_PRICES = [100, 1000, 10000, 100000, 1000000, 10000000];
const FIRST_EPOCH = 1;
const SUBMISSION_WINDOW = 10;
// incentive manager parameters
const BASE_SAMPLE_SIZE = 5 * 2 ** 8; // 2^8 since scaled for 2^(-8) for fixed precision arithmetic
const BASE_RANGE = 2 * 2 ** 8;
const SAMPLE_INCREASE_LIMIT = 5 * 2 ** 8;
const RANGE_INCREASE_PRICE = 5;
const DURATION = 8;

export async function deployContracts(hre: HardhatRuntimeEnvironment, parameters: FTSOParameters) {
    await syncTimeToNow(hre);

    const artifacts = hre.artifacts;
    const governance = web3.eth.accounts.privateKeyToAccount(parameters.governancePrivateKey);

    const voterRegistry = await artifacts.require("VoterRegistry").new(governance.address);
    const fastUpdaters = await artifacts.require("FastUpdaters").new(voterRegistry.address);
    const fastUpdateIncentiveManager = await artifacts
        .require("FastUpdateIncentiveManager")
        .new(governance.address, BASE_SAMPLE_SIZE, BASE_RANGE, SAMPLE_INCREASE_LIMIT, RANGE_INCREASE_PRICE, DURATION);

    const fastUpdater = await artifacts
        .require("FastUpdater")
        .new(fastUpdaters.address, fastUpdateIncentiveManager.address, ANCHOR_PRICES, SUBMISSION_WINDOW, FIRST_EPOCH);

    const deployed = <ContractAddresses>{
        voterRegistry: voterRegistry.address,
        fastUpdaters: fastUpdaters.address,
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
