import BN from "bn.js";
import { FastUpdaterInstance, FlareSystemMockInstance, FastUpdateIncentiveManagerInstance } from "../typechain-truffle";
import { getTestFile } from "../test-utils/utils/constants";
import { KeyGen, VerifiableRandomness, SortitionKey, Proof } from "../src/Sortition";
import { RandInt } from "../src/utils/rand";
import { loadAccounts } from "../deployment/tasks/common";
import { Account } from "web3-core";
import { RangeFPA, SampleFPA } from "../src/utils/fixed-point-arithmetics";

const FastUpdater = artifacts.require("FastUpdater");
const FastUpdateIncentiveManager = artifacts.require("FastUpdateIncentiveManager");
const FlareSystemMock = artifacts.require("FlareSystemMock");

const NUM_ACCOUNTS = 10;
const ANCHOR_PRICES = [1000, 10000, 100000, 1000000, 10000000, 100000000];
const NUM_FEEDS = ANCHOR_PRICES.length;
let TEST_EPOCH: bigint;
const VOTER_WEIGHT = 1000;
const SUBMISSION_WINDOW = 10;
const BASE_SAMPLE_SIZE = SampleFPA(2);
const BASE_RANGE = RangeFPA(2 ** -5);
const SAMPLE_INCREASE_LIMIT = SampleFPA(5);
const SCALE = 1 + BASE_RANGE / BASE_SAMPLE_SIZE;
const RANGE_INCREASE_PRICE = 5;
const DURATION = 8;
const ZEROS64 = "0x" + "0".repeat(64);
const ZEROS52 = "0x" + "0".repeat(52);
const EPOCH_LEN = 1000;

contract(`FastUpdater.sol; ${getTestFile(__filename)}`, async () => {
    let fastUpdater: FastUpdaterInstance;
    let fastUpdateIncentiveManager: FastUpdateIncentiveManagerInstance;
    let flareSystemMock: FlareSystemMockInstance;
    let accounts: Account[];
    let keys: SortitionKey[];
    const weights: number[] = [];
    before(async () => {
        accounts = loadAccounts(web3);
        const governance = accounts[0];

        flareSystemMock = await FlareSystemMock.new(RandInt(2n ** 256n - 1n), EPOCH_LEN);
        fastUpdateIncentiveManager = await FastUpdateIncentiveManager.new(
            governance.address,
            BASE_SAMPLE_SIZE,
            BASE_RANGE,
            SAMPLE_INCREASE_LIMIT,
            RANGE_INCREASE_PRICE,
            DURATION
        );

        TEST_EPOCH = await flareSystemMock.getCurrentRewardEpochId();

        keys = new Array<SortitionKey>(NUM_ACCOUNTS);
        for (let i = 0; i < NUM_ACCOUNTS; i++) {
            const key: SortitionKey = KeyGen();
            keys[i] = key;
            const x = "0x" + web3.utils.padLeft(key.pk.x.toString(16), 64);
            const y = "0x" + web3.utils.padLeft(key.pk.y.toString(16), 64);
            const newProvider = [x, y, BigInt(VOTER_WEIGHT)];
            await flareSystemMock.registerAsVoter(TEST_EPOCH, newProvider, { from: accounts[i + 1].address });
        }

        fastUpdater = await FastUpdater.new(
            flareSystemMock.address,
            flareSystemMock.address,
            fastUpdateIncentiveManager.address,
            ANCHOR_PRICES,
            SUBMISSION_WINDOW
        );
    });

    it("should submit updates", async () => {
        let submissionBlockNum;

        for (let i = 0; i < NUM_ACCOUNTS; i++) {
            const weight = await fastUpdater.currentSortitionWeight(accounts[i + 1].address);
            weights[i] = Number(weight.toString());
            expect(weights[i]).to.equal(Math.floor(4096 / NUM_ACCOUNTS));
        }

        const feeds: number[] = [];
        for (let i = 0; i < NUM_FEEDS; i++) {
            feeds.push(i);
        }
        const startingPricesBN: BN[] = await fastUpdater.fetchCurrentPrices(feeds);
        const startingPrices: bigint[] = [];
        // console.log("Starting prices");
        for (let i = 0; i < NUM_FEEDS; i++) {
            startingPrices[i] = BigInt(startingPricesBN[i].toString());
            // console.log(BigInt(startingPrices[i]));
        }

        const feed = "+-0-0+";
        let delta = "0x731" + "0".repeat(61);
        let numSubmitted = 0;
        for (;;) {
            submissionBlockNum = await web3.eth.getBlockNumber();

            const scoreCutoff = await fastUpdater.currentScoreCutoff();
            const baseSeed = await flareSystemMock.getCurrentRandom();
            // console.log(submissionBlockNum, baseSeed.toString());
            for (let i = 0; i < NUM_ACCOUNTS; i++) {
                for (let rep = 0; rep < weights[i]; rep++) {
                    const replicate = BigInt(rep);
                    const proof: Proof = VerifiableRandomness(keys[i], baseSeed, BigInt(submissionBlockNum), replicate);
                    const sortitionCredential = [replicate, [proof.gamma.x, proof.gamma.y], proof.c, proof.s];

                    if (proof.gamma.x < scoreCutoff) {
                        // console.log("submitting", feed, "client", i, "with rep", rep);

                        const deltas = [[delta, ZEROS64, ZEROS64, ZEROS64, ZEROS64, ZEROS64, ZEROS64], ZEROS52];
                        const newFastUpdate = [submissionBlockNum, sortitionCredential, deltas];

                        await fastUpdater.submitUpdates(newFastUpdate, { from: accounts[i + 1].address });
                        let caughtError = false;
                        try {
                            // test if submitting again gives error
                            await fastUpdater.submitUpdates(newFastUpdate, { from: accounts[i + 1].address });
                        } catch (e) {
                            expect(e).to.be.not.empty;
                            caughtError = true;
                        }
                        expect(caughtError).to.equal(true);
                        numSubmitted++;
                    }
                }
            }
            if (numSubmitted > 0) break;
            await fastUpdater.freeSubmitted({ from: accounts[0].address });
        }
        let pricesBN: BN[] = await fastUpdater.fetchCurrentPrices.call(feeds);
        const prices: bigint[] = [];
        // console.log("Middle prices");
        for (let i = 0; i < NUM_FEEDS; i++) {
            prices[i] = BigInt(pricesBN[i].toString());
            let sign = 0;
            if (feed[i] == "+") {
                sign = 1;
            }
            if (feed[i] == "-") {
                sign = -1;
            }
            expect(Number(prices[i])).to.be.greaterThanOrEqual(
                (SCALE ** sign) ** numSubmitted * Number(startingPrices[i]) * 0.99
            );
            expect(Number(prices[i])).to.be.lessThanOrEqual(
                (SCALE ** sign) ** numSubmitted * Number(startingPrices[i]) * 1.01
            );
            // console.log(BigInt(prices[i]));
        }

        delta = "0xd13" + "0".repeat(61);
        let breakVar = false;
        while (!breakVar) {
            submissionBlockNum = await web3.eth.getBlockNumber();
            // console.log(submissionBlockNum);

            const scoreCutoff = await fastUpdater.currentScoreCutoff();
            const baseSeed = await flareSystemMock.getCurrentRandom();
            // console.log(submissionBlockNum, baseSeed.toString());

            for (let i = 0; i < NUM_ACCOUNTS; i++) {
                for (let rep = 0; rep < weights[i]; rep++) {
                    const replicate = BigInt(rep);
                    const proof: Proof = VerifiableRandomness(keys[i], baseSeed, BigInt(submissionBlockNum), replicate);
                    const sortitionCredential = [replicate, [proof.gamma.x, proof.gamma.y], proof.c, proof.s];

                    if (proof.gamma.x < scoreCutoff) {
                        // console.log("submitting -+0+0- client", i, "with rep", rep);
                        const deltas = [[delta, ZEROS64, ZEROS64, ZEROS64, ZEROS64, ZEROS64, ZEROS64], ZEROS52];
                        const newFastUpdate = [submissionBlockNum, sortitionCredential, deltas];

                        await fastUpdater.submitUpdates(newFastUpdate, { from: accounts[i + 1].address });
                        numSubmitted--;
                        if (numSubmitted == 0) {
                            breakVar = true;
                            break;
                        }
                    }
                }
                if (breakVar) break;
            }

            await fastUpdater.freeSubmitted({ from: accounts[0].address });
        }

        pricesBN = await fastUpdater.fetchCurrentPrices.call(feeds);
        // console.log("End prices");
        for (let i = 0; i < NUM_FEEDS; i++) {
            prices[i] = BigInt(pricesBN[i].toString());
            // console.log(BigInt(prices[i]));
        }
        for (let i = 0; i < NUM_FEEDS; i++) {
            expect(Number(prices[i])).to.be.greaterThanOrEqual(Number(startingPrices[i]) * 0.99);
            expect(Number(prices[i])).to.be.lessThanOrEqual(Number(startingPrices[i]) * 1.01);
        }
    });
});
