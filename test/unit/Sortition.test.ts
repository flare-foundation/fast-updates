import { TestSortitionContractInstance } from "../../typechain-truffle";
import { getTestFile } from "../../test-utils/utils/constants";
import { RandInt } from "../../src/utils/rand";
import { KeyGen, VerifiableRandomness, SortitionKey, Proof } from "../../src/Sortition";
import { bn254 } from "@noble/curves/bn254";

const SortitionContract = artifacts.require("TestSortitionContract");

contract(`Sortition.sol; ${getTestFile(__filename)}`, async accounts => {
    let sortition: TestSortitionContractInstance;
    before(async () => {
        const governance = accounts[0];
        sortition = await SortitionContract.new(governance);
    });

    it("should generate a verifiable randomness", async () => {
        const key: SortitionKey = KeyGen();
        const seed = RandInt(bn254.CURVE.n);
        const blockNum = BigInt(await web3.eth.getBlockNumber());
        const replicate = RandInt(bn254.CURVE.n);
        const proof: Proof = VerifiableRandomness(key, seed, blockNum, replicate);
        const pubKey = [key.pk.x, key.pk.y];
        const sortitionCredential = [replicate, [proof.gamma.x, proof.gamma.y], proof.c, proof.s];

        const sortitionState = [seed, blockNum, 0, 0, pubKey];

        const check = await sortition.testVerifySortitionProof(sortitionState, sortitionCredential);

        expect(check).to.equal(true);
    });
    it("should correctly accept or reject the randomness", async () => {
        const key: SortitionKey = KeyGen();
        const scoreCutoff = BigInt(2) ** BigInt(256 - 8);
        for (;;) {
            const seed: bigint = RandInt(bn254.CURVE.n);
            const replicate = RandInt(bn254.CURVE.n);
            const blockNum = BigInt(await web3.eth.getBlockNumber());
            const weight = replicate + BigInt(1);

            const proof: Proof = VerifiableRandomness(key, seed, blockNum, replicate);
            const sortitionRound = [true, seed, scoreCutoff];
            const pubKey = [key.pk.x, key.pk.y];
            const sortitionCredential = [replicate, [proof.gamma.x, proof.gamma.y], proof.c, proof.s];

            const sortitionState = [seed, await web3.eth.getBlockNumber(), scoreCutoff, weight, pubKey];

            const check = await sortition.testVerifySortitionCredential(sortitionState, sortitionCredential);

            if (proof.gamma.x > scoreCutoff) {
                expect(check).to.equal(false);
            } else {
                expect(check).to.equal(true);
                break;
            }
        }
    });
});
