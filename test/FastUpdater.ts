import BN from "bn.js";
import chai, { expect } from "chai";
import chaiBN from "chai-bn";
import {
    FastUpdatersInstance,
    FastUpdaterInstance,
    VoterRegistryInstance,
    FastUpdateIncentiveManagerInstance,
} from "../typechain-truffle";
import { getTestFile } from "../test-utils/utils/constants";
import { KeyGen, VerifiableRandomness, SortitionKey, Proof } from "../src/Sortition";
import { RandInt } from "../src/utils/rand";
import { bn254 } from "@noble/curves/bn254";
import { toBN } from "../src/protocol/utils/voting-utils";
import { loadAccounts } from "../deployment/tasks/common";
import { Account } from "web3-core";
import { int } from "hardhat/internal/core/params/argumentTypes";
// import { getBlockNumber } from "../src/providers/TruffleProvider";
// import { getBlockNumber } from "../src/TruffleProvider";

const FastUpdaters = artifacts.require("FastUpdaters");
const FastUpdater = artifacts.require("FastUpdater");
const FastUpdateIncentiveManager = artifacts.require("FastUpdateIncentiveManager");
const VoterRegistry = artifacts.require("VoterRegistry");

const NUM_ACCOUNTS = 10;
const ANCHOR_PRICES = [10, 100, 1000, 10000, 100000, 1000000];
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
    before(async () => {
        accounts = loadAccounts(web3);
        // const governance = accounts[0];
        fastUpdaters = await FastUpdaters.new();
        fastUpdater = await FastUpdater.new();
        fastUpdateIncentiveManager = await FastUpdateIncentiveManager.new();
        fastUpdateIncentiveManager.setBase(100, 100, 100, [1], [1]);

        voterRegistry = await VoterRegistry.new();
        await fastUpdaters.setVoterRegistry(voterRegistry.address);
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
            // proofs[i] = proof;
            const pubKey = [key.pk.x, key.pk.y];
            const sortitionCredential = [replicate, [proof.gamma.x, proof.gamma.y], proof.c, proof.s];
            const newProvider = [pubKey, sortitionCredential];
            credentials[i] = sortitionCredential;
            await fastUpdaters.registerNewProvider(newProvider, { from: accounts[i + 1].address });
        }

        // const nextData = await fastUpdaters.nextProviderRegistry.call(TEST_EPOCH);
        // for (let i = 0; i < NUM_ACCOUNTS; i++) {
        //     const providerAddress = nextData[1][i];
        //     expect(providerAddress).to.equal(accounts[i + 1].address);
        //     const providerKey = nextData[2][i];
        //     expect(providerKey[0]).to.equal(keys[i].pk.x);
        //     expect(providerKey[1]).to.equal(keys[i].pk.y);
        //     const providerWeight = parseInt(nextData[3][i]);
        //     const expectedWeight = Math.floor((VOTER_WEIGHT << 12) / (VOTER_WEIGHT * NUM_ACCOUNTS));
        //     expect(providerWeight).to.equal(expectedWeight);
        // }
    });

    it("should submit updates", async () => {
        await fastUpdater.setFastUpdaters(fastUpdaters.address);
        await fastUpdater.setFastUpdateIncentiveManager(fastUpdateIncentiveManager.address);
        await fastUpdater.setSubmissionWindow(10);

        // const extendedAnchorPrices: number[] = Array();
        // for (let i = ANCHOR_PRICES.length; i < 1000; i++) {
        //     extendedAnchorPrices.push(i);
        // }
        // const anchorPrices = ANCHOR_PRICES.concat(extendedAnchorPrices);
        await fastUpdater.setAnchorPrices(ANCHOR_PRICES);

        await fastUpdater.prepareForNewBlock(true, TEST_EPOCH);
        const submissionBlockNum = (await web3.eth.getBlockNumber()) + 1;

        const feeds: number[] = new Array();
        for (let i = 0; i < NUM_FEEDS; i++) {
            feeds.push(i);
        }
        const startingPrices: BN[] = await fastUpdater.fetchCurrentPrices(feeds);
        console.log("Starting prices");
        for (let i = 0; i < NUM_FEEDS; i++) {
            console.log(BigInt(startingPrices[i].toString()));
        }

        const sortitionRound = await fastUpdater.getSortitionRound(submissionBlockNum);
        // console.log("sortition round", sortitionRound[0].toString(), sortitionRound[1].toString());

        // struct FastUpdates {
        //     uint sortitionBlock;
        //     SortitionCredential sortitionCredential;
        //     Deltas deltas;
        // }
        // console.log(credentials[0]);
        for (let i = 0; i < NUM_ACCOUNTS; i++) {
            for (let rep = 0; rep < 409; rep++) {
                const replicate = BigInt(rep);
                const proof: Proof = VerifiableRandomness(keys[i], sortitionRound.seed, replicate);
                const sortitionCredential = [replicate, [proof.gamma.x, proof.gamma.y], proof.c, proof.s];

                if (proof.gamma.x < sortitionRound.cutoff) {
                    console.log("success", i, proof.gamma.x, sortitionRound.cutoff.toString());
                    const delta1 = "0x4330000000000000000000000000000000000000000000000000000000000000";
                    const delta2 = "0x0000000000000000000000000000000000000000000000000000";
                    const deltas = [[delta1, delta1, delta1, delta1, delta1, delta1, delta1], delta2];
                    const newFastUpdate = [submissionBlockNum, sortitionCredential, deltas];
                    // console.log(newFastUpdate);

                    await fastUpdater.submitUpdates(newFastUpdate, { from: accounts[i + 1].address });
                } else {
                    // console.log("fail", i, proof.gamma.x, sortitionRound.cutoff.toString());
                }
            }
        }
        const endPrices: BN[] = await fastUpdater.fetchCurrentPrices(feeds);
        console.log("End prices");
        for (let i = 0; i < NUM_FEEDS; i++) {
            console.log(BigInt(endPrices[i].toString()));
        }
    });
});
