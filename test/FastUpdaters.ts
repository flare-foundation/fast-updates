import { expect } from "chai";
import { FastUpdatersInstance, VoterRegistryInstance } from "../typechain-truffle";
import { getTestFile } from "../test-utils/utils/constants";
import { KeyGen, VerifiableRandomness, SortitionKey, Proof } from "../src/Sortition";
import { toBN } from "../src/protocol/utils/voting-utils";
import { loadAccounts } from "../deployment/tasks/common";
import { Account } from "web3-core";

const FastUpdaters = artifacts.require("FastUpdaters");
const VoterRegistry = artifacts.require("VoterRegistry");

const NUM_ACCOUNTS = 10;
const TEST_EPOCH = 1;
const VOTER_WEIGHT = 1000;

contract(`FastUpdaters.sol; ${getTestFile(__filename)}`, async () => {
    let fastUpdaters: FastUpdatersInstance;
    let voterRegistry: VoterRegistryInstance;
    let accounts: Account[];
    let seed: bigint;
    before(async () => {
        accounts = loadAccounts(web3);
        // const governance = accounts[0];
        fastUpdaters = await FastUpdaters.new();
        voterRegistry = await VoterRegistry.new();
        await fastUpdaters.setVoterRegistry(voterRegistry.address);
        for (let i = 1; i <= NUM_ACCOUNTS; i++) {
            await voterRegistry.registerAsAVoter(TEST_EPOCH, toBN(VOTER_WEIGHT), { from: accounts[i].address });
        }
        const seedBN = await fastUpdaters.getBaseSeed.call();
        seed = BigInt(seedBN.toString());
    });

    it("should register a new provider", async () => {
        const keys = new Array<SortitionKey>(NUM_ACCOUNTS);
        const proofs = new Array<Proof>(NUM_ACCOUNTS);
        for (let i = 0; i < NUM_ACCOUNTS; i++) {
            const key: SortitionKey = KeyGen();
            keys[i] = key;
            const replicate = BigInt(0); // Registration doesn't allow cherry-picking a replicate
            const proof: Proof = VerifiableRandomness(key, seed, replicate);
            proofs[i] = proof;
            const pubKey = [key.pk.x, key.pk.y];
            const sortitionCredential = [replicate, [proof.gamma.x, proof.gamma.y], proof.c, proof.s];
            const newProvider = [pubKey, sortitionCredential];
            await fastUpdaters.registerNewProvider(newProvider, { from: accounts[i + 1].address });
        }

        const nextData = await fastUpdaters.nextProviderRegistry.call(TEST_EPOCH);
        for (let i = 0; i < NUM_ACCOUNTS; i++) {
            const providerAddress = nextData[1][i];
            expect(providerAddress).to.equal(accounts[i + 1].address);
            const providerKey = nextData[2][i];
            expect(providerKey[0]).to.equal(keys[i].pk.x);
            expect(providerKey[1]).to.equal(keys[i].pk.y);
            const providerWeight = parseInt(nextData[3][i]);
            const expectedWeight = Math.floor((VOTER_WEIGHT << 12) / (VOTER_WEIGHT * NUM_ACCOUNTS));
            expect(providerWeight).to.equal(expectedWeight);
        }
    });
});
