import BN from "bn.js";
import chai, { expect } from "chai";
import chaiBN from "chai-bn";
import { TestBn256Instance } from "../../typechain-truffle";
import { getTestFile } from "../../test-utils/utils/constants";
import { bn254 } from "@noble/curves/bn254"; // also known as alt_bn128
import { RandInt } from "../../src/utils/rand";
import { g1compress } from "../../src/Sortition";

const TestBn256 = artifacts.require("TestBn256");
chai.use(chaiBN(BN));

contract(`Bn256.sol; ${getTestFile(__filename)}`, async accounts => {
    let bn256Instance: TestBn256Instance;
    before(async () => {
        const governance = accounts[0];
        bn256Instance = await TestBn256.new(governance);
    });

    it("should add two points", async () => {
        const r1 = RandInt(bn254.CURVE.n);
        const r2 = RandInt(bn254.CURVE.n);
        const a = bn254.ProjectivePoint.BASE.multiply(r1);
        const b = bn254.ProjectivePoint.BASE.multiply(r2);

        const c = await bn256Instance.publicG1Add([a.x, a.y], [b.x, b.y]);

        const cCheck = a.add(b);
        expect(c[0].toString()).to.equal(cCheck.x.toString());
        expect(c[1].toString()).to.equal(cCheck.y.toString());
    });

    it("should multiply a point with a scalar", async () => {
        const r1 = RandInt(bn254.CURVE.n);
        const r2 = RandInt(bn254.CURVE.n);
        const a = bn254.ProjectivePoint.BASE.multiply(r1);

        const c = await bn256Instance.publicG1ScalarMultiply([a.x, a.y], r2);

        const cCheck = a.multiply(r2);
        expect(c[0].toString()).to.equal(cCheck.x.toString());
        expect(c[1].toString()).to.equal(cCheck.y.toString());
    });
    it("should compress/decompress a point", async () => {
        const r1 = RandInt(bn254.CURVE.n);
        const a = bn254.ProjectivePoint.BASE.multiply(r1);

        const aCompressed = g1compress(a);
        const aCheck = await bn256Instance.publicG1Decompress(aCompressed);
        expect(a.x.toString()).to.equal(aCheck[0].toString());
        expect(a.y.toString()).to.equal(aCheck[1].toString());
    });
});
