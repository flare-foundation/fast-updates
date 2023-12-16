import BN from "bn.js";
import {
    FastUpdatersInstance,
    FastUpdaterInstance,
    VoterRegistryInstance,
    FastUpdateIncentiveManagerInstance,
} from "../typechain-truffle";
import { getTestFile } from "../test-utils/utils/constants";
import { KeyGen, VerifiableRandomness, SortitionKey, Proof } from "../src/Sortition";
import { toBN } from "../src/protocol/utils/voting-utils";
import { loadAccounts } from "../deployment/tasks/common";
import { Account } from "web3-core";

const FastUpdaters = artifacts.require("FastUpdaters");
const FastUpdater = artifacts.require("FastUpdater");
const FastUpdateIncentiveManager = artifacts.require("FastUpdateIncentiveManager");
const VoterRegistry = artifacts.require("VoterRegistry");

const NUM_ACCOUNTS = 10;
const ANCHOR_PRICES = [100, 1000, 10000, 100000, 1000000, 10000000];
const NUM_FEEDS = ANCHOR_PRICES.length;
const TEST_EPOCH = 1;
const VOTER_WEIGHT = 1000;

contract(`FastUpdater.sol; ${getTestFile(__filename)}`, async () => {
    let fastUpdaters: FastUpdatersInstance;
    let fastUpdater: FastUpdaterInstance;
    let fastUpdateIncentiveManager: FastUpdateIncentiveManagerInstance;
    let voterRegistry: VoterRegistryInstance;
    let accounts: Account[];
    let credentials: (bigint | bigint[])[][];
    let keys: SortitionKey[];
    const weights: number[] = Array();
    before(async () => {
        accounts = loadAccounts(web3);
        const governance = accounts[0];

        voterRegistry = await VoterRegistry.new();

        fastUpdaters = await FastUpdaters.new(voterRegistry.address);
        fastUpdateIncentiveManager = await FastUpdateIncentiveManager.new(governance.address, 100, 1, 1, 1, 8);
        // fastUpdateIncentiveManager.setBase(100, 100, 100, [1], [1]);

        for (let i = 1; i <= NUM_ACCOUNTS; i++) {
            await voterRegistry.registerAsAVoter(TEST_EPOCH, toBN(VOTER_WEIGHT), { from: accounts[i].address });
        }
        const seedBN = await fastUpdaters.getBaseSeed.call();
        const seed = BigInt(seedBN.toString());
        keys = new Array<SortitionKey>(NUM_ACCOUNTS);
        credentials = new Array(NUM_ACCOUNTS);
        for (let i = 0; i < NUM_ACCOUNTS; i++) {
            const key: SortitionKey = KeyGen();
            keys[i] = key;
            const replicate = BigInt(0); // Registration doesn't allow cherry-picking a replicate
            const proof: Proof = VerifiableRandomness(key, seed, replicate);
            const pubKey = [key.pk.x, key.pk.y];
            const sortitionCredential = [replicate, [proof.gamma.x, proof.gamma.y], proof.c, proof.s];
            const newProvider = [pubKey, sortitionCredential];
            credentials[i] = sortitionCredential;
            await fastUpdaters.registerNewProvider(newProvider, { from: accounts[i + 1].address });
        }

        fastUpdater = await FastUpdater.new(
            fastUpdaters.address, 
            fastUpdateIncentiveManager.address,
            ANCHOR_PRICES,
            10,
            TEST_EPOCH
        );
    });

    it("should submit updates", async () => {
        let submissionBlockNum = (await web3.eth.getBlockNumber()) + 1;

        for (let i = 0; i < NUM_ACCOUNTS; i++) {
            const provider = await fastUpdater.activeProviders.call(accounts[i + 1].address);
            weights[i] = Number(provider[1].toString());
        }

        const feeds: number[] = new Array();
        for (let i = 0; i < NUM_FEEDS; i++) {
            feeds.push(i);
        }
        const startingPricesBN: BN[] = await fastUpdater.fetchCurrentPrices(feeds);
        const startingPrices: bigint[] = new Array();
        console.log("Starting prices");
        for (let i = 0; i < NUM_FEEDS; i++) {
            startingPrices[i] = BigInt(startingPricesBN[i].toString());
            console.log(BigInt(startingPrices[i]));
        }
        let breakVar = false;
        console.log();
        console.log("Blocks")
        while (!breakVar) {
            const sortitionRound = await fastUpdater.getSortitionRound(submissionBlockNum);
            console.log(submissionBlockNum, sortitionRound.seed.toString());
            for (let i = 0; i < NUM_ACCOUNTS; i++) {
                for (let rep = 0; rep < weights[i]; rep++) {
                    const replicate = BigInt(rep);
                    const proof: Proof = VerifiableRandomness(keys[i], sortitionRound.seed, replicate);
                    const sortitionCredential = [replicate, [proof.gamma.x, proof.gamma.y], proof.c, proof.s];

                    if (proof.gamma.x < sortitionRound.cutoff) {
                        console.log("submitting +-0-0+ client", i, "with rep ", rep);
                        const delta1 = "0x7310000000000000000000000000000000000000000000000000000000000000";
                        const delta2 = "0x0000000000000000000000000000000000000000000000000000";
                        const deltas = [[delta1, delta1, delta1, delta1, delta1, delta1, delta1], delta2];
                        const newFastUpdate = [submissionBlockNum, sortitionCredential, deltas];

                        await fastUpdater.submitUpdates(newFastUpdate, { from: accounts[i + 1].address });
                        breakVar = true;
                        break;
                    }
                }
                if (breakVar) break;
            }
            await fastUpdater.finalizeBlock(false, TEST_EPOCH);
            submissionBlockNum = (await web3.eth.getBlockNumber()) + 1;
            // console.log(submissionBlockNum);
        }
        let pricesBN: BN[] = await fastUpdater.fetchCurrentPrices(feeds);
        const prices: bigint[] = new Array();
        console.log("Middle prices");
        for (let i = 0; i < NUM_FEEDS; i++) {
            prices[i] = BigInt(pricesBN[i].toString());
            console.log(BigInt(prices[i]));
        }
        breakVar = false;
        while (!breakVar) {
            const sortitionRound = await fastUpdater.getSortitionRound(submissionBlockNum);
            console.log(submissionBlockNum, sortitionRound.seed.toString());

            for (let i = 0; i < NUM_ACCOUNTS; i++) {
                for (let rep = 0; rep < weights[i]; rep++) {
                    const replicate = BigInt(rep);
                    const proof: Proof = VerifiableRandomness(keys[i], sortitionRound.seed, replicate);
                    const sortitionCredential = [replicate, [proof.gamma.x, proof.gamma.y], proof.c, proof.s];

                    if (proof.gamma.x < sortitionRound.cutoff) {
                        console.log("submitting -+0+0- client", i, "with rep ", rep);
                        const delta1 = "0xd130000000000000000000000000000000000000000000000000000000000000";
                        const delta2 = "0x0000000000000000000000000000000000000000000000000000";
                        const deltas = [[delta1, delta1, delta1, delta1, delta1, delta1, delta1], delta2];
                        const newFastUpdate = [submissionBlockNum, sortitionCredential, deltas];

                        await fastUpdater.submitUpdates(newFastUpdate, { from: accounts[i + 1].address });
                        breakVar = true;
                        break;
                    }
                }
                if (breakVar) break;
            }
            await fastUpdater.finalizeBlock(true, TEST_EPOCH);
            submissionBlockNum = (await web3.eth.getBlockNumber()) + 1;
            // console.log(submissionBlockNum);
        }

        pricesBN = await fastUpdater.fetchCurrentPrices(feeds);
        console.log("End prices");
        for (let i = 0; i < NUM_FEEDS; i++) {
            prices[i] = BigInt(pricesBN[i].toString());
            console.log(BigInt(prices[i]));
        }
        for (let i = 0; i < NUM_FEEDS; i++) {
            expect(Number(prices[i])).to.be.greaterThanOrEqual(Number(startingPrices[i]) * 0.99);
            expect(Number(prices[i])).to.be.lessThanOrEqual(Number(startingPrices[i]) * 1.01);
        }
    });
});
