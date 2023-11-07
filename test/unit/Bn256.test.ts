import BN from "bn.js";
import chai, { expect } from "chai";
import chaiBN from "chai-bn";
import { TestBn256Instance } from "../../typechain-truffle";
import { getTestFile } from "../../test-utils/utils/constants";
import { bn254 } from "@noble/curves/bn254"; // also known as alt_bn128
import crypto from "crypto";

const TestBn256 = artifacts.require("TestBn256");
chai.use(chaiBN(BN));
const q = BigInt("21888242871839275222246405745257275088548364400416034343698204186575808495617");

contract(`Bn256.sol; ${getTestFile(__filename)}`, async accounts => {
  let altBn128: TestBn256Instance;
  before(async () => {
    const governance = accounts[0];
    altBn128 = await TestBn256.new(governance);
  });

  it("should add two points", async () => {
    const r1 = RandInt(31);
    const r2 = RandInt(31);
    const a = bn254.ProjectivePoint.BASE.multiply(r1);
    const b = bn254.ProjectivePoint.BASE.multiply(r2);

    const c = await altBn128.PublicG1Add([a.x, a.y], [b.x, b.y]);

    const cCheck = a.add(b);
    expect(c[0].toString()).to.equal(cCheck.x.toString());
    expect(c[1].toString()).to.equal(cCheck.y.toString());
  });

  it("should multiply a point with a scalar", async () => {
    const r1 = RandInt(31);
    const r2 = RandInt(31);
    const a = bn254.ProjectivePoint.BASE.multiply(r1);

    const c = await altBn128.PublicG1ScalarMultiply([a.x, a.y], r2);

    const cCheck = a.multiply(r2);
    expect(c[0].toString()).to.equal(cCheck.x.toString());
    expect(c[1].toString()).to.equal(cCheck.y.toString());
  });
});

export function RandInt(numBytes: number) {
  const randbytes = crypto.randomBytes(numBytes).toString("hex");
  const r = BigInt("0x" + randbytes);

  return r;
}
