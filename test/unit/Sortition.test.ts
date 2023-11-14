import BN from "bn.js";
import chai, { expect } from "chai";
import chaiBN from "chai-bn";
import { SortitionContractInstance } from "../../typechain-truffle";
import { getTestFile } from "../../test-utils/utils/constants";
import { RandInt } from "../../src/utils/rand";
import { KeyGen, VerifiableRandomness, SortitionKey, Proof } from "../../src/Sortition";
import { bn254 } from "@noble/curves/bn254"; // also known as alt_bn128

const SortitionContract = artifacts.require("SortitionContract");
chai.use(chaiBN(BN));

contract(`Sortition.sol; ${getTestFile(__filename)}`, async accounts => {
    let sortition: SortitionContractInstance;
    before(async () => {
        const governance = accounts[0];
        sortition = await SortitionContract.new(governance);
    });

    it("should generate a verifiable randomness", async () => {
        const key: SortitionKey = KeyGen();
        const seed = RandInt(bn254.CURVE.n);
        const replicate = RandInt(bn254.CURVE.n);
        const proof: Proof = VerifiableRandomness(key, seed, replicate);
        const sortitionRound = [seed, 1000];

        const check = await sortition.VerifySortitionProof(
            sortitionRound,
            [key.pk.toAffine().x, key.pk.toAffine().y],
            [replicate, [proof.gamma.toAffine().x, proof.gamma.toAffine().y], proof.c, proof.s]
        );

        expect(check).to.equal(true);
    });
});
