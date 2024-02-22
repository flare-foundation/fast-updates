import BN from "bn.js";
import { FastUpdateIncentiveManagerInstance } from "../typechain-truffle";
import { getTestFile } from "../test-utils/utils/constants";
import { KeyGen, VerifiableRandomness, SortitionKey, Proof } from "../src/Sortition";
import { toBN } from "../src/protocol/utils/voting-utils";
import { loadAccounts } from "../deployment/tasks/common";
import { Account } from "web3-core";

const FastUpdateIncentiveManager = artifacts.require("FastUpdateIncentiveManager");

const BASE_SAMPLE_SIZE = 5 * 2 ** 8; // 2^8 since scaled for 2^(-8) for fixed precision arithmetic
const BASE_RANGE = 2 * 2 ** 8;
const SAMPLE_INCREASE_LIMIT = 5 * 2 ** 8;
const RANGE_INCREASE_PRICE = 5;
const DURATION = 8;

contract(`FastUpdateIncentiveManager.sol; ${getTestFile(__filename)}`, async () => {
    let fastUpdateIncentiveManager: FastUpdateIncentiveManagerInstance;
    let accounts: Account[];
    before(async () => {
        accounts = loadAccounts(web3);
        const governance = accounts[0];
        fastUpdateIncentiveManager = await FastUpdateIncentiveManager.new(
            governance.address,
            governance.address,
            BASE_SAMPLE_SIZE,
            BASE_RANGE,
            SAMPLE_INCREASE_LIMIT,
            RANGE_INCREASE_PRICE,
            DURATION
        );
    });

    it("should get expected sample size", async () => {
        const sampleSize = await fastUpdateIncentiveManager.getExpectedSampleSize();
        expect(sampleSize).to.equal(BASE_SAMPLE_SIZE);
    });
    it("should get range", async () => {
        const range = await fastUpdateIncentiveManager.getRange();
        expect(range).to.equal(BASE_RANGE);
    });
    it("should get precision", async () => {
        const precision = await fastUpdateIncentiveManager.getPrecision();
        // precision scaled for 2^(-15)
        expect(precision).to.equal(Math.floor((BASE_RANGE / BASE_SAMPLE_SIZE) * 2 ** 15));
    });
    it("should get scale", async () => {
        const scale = await fastUpdateIncentiveManager.getScale();
        expect(scale).to.equal(Math.floor(2 ** 15 + (BASE_RANGE / BASE_SAMPLE_SIZE) * 2 ** 15));
    });

    it("should offer incentive", async () => {
        const rangeIncrease = 2 * 2 ** 8;
        const rangeLimit = 4 * 2 ** 8;
        const offer = [rangeIncrease, rangeLimit];
        await fastUpdateIncentiveManager.offerIncentive(offer, {
            from: accounts[1].address,
            value: 100000,
        });
        // todo: contract needs to be finished

        // expect(param[0]).to.equal(BASE_SAMPLE_SIZE);
        // expect(param[1]).to.equal(Math.floor(2 ** 15 + (BASE_RANGE / BASE_SAMPLE_SIZE) * 2 ** 15));
    });
});
