import { RandInt } from "./utils/rand";
import { bn254 } from "@noble/curves/bn254"; // also known as alt_bn128
import { Field } from "@noble/curves/abstract/modular"; // also known as alt_bn128
import { ProjPointType, AffinePoint } from "@noble/curves/abstract/weierstrass";
import { encodePacked } from "web3-utils";
import { toBN } from "./protocol/utils/voting-utils";
import { sha256 } from "ethers";

export type SortitionKey = {
    sk: bigint;
    pk: ProjPointType<bigint>;
};

export type Proof = {
    gamma: ProjPointType<bigint>;
    c: bigint;
    s: bigint;
};

export function KeyGen(): SortitionKey {
    const sk = RandInt(bn254.CURVE.n);
    const pk = bn254.ProjectivePoint.BASE.multiply(sk);
    const key: SortitionKey = { sk: sk, pk: pk };

    return key;
}

export function VerifiableRandomness(key: SortitionKey, seed: bigint, replicate: bigint): Proof {
    let toHash: string = encodePacked(
        { value: toBN(seed.toString()), type: "uint256" },
        { value: toBN(replicate.toString()), type: "uint256" }
    );

    const h = g1HashToPoint(toHash);
    const gamma = h.multiply(key.sk);

    const k = RandInt(bn254.CURVE.n);
    const gToK = bn254.ProjectivePoint.BASE.multiply(k);
    const hToK = h.multiply(k);
    toHash = encodePacked(
        { value: toBN(bn254.ProjectivePoint.BASE.x.toString()), type: "uint256" },
        { value: toBN(bn254.ProjectivePoint.BASE.y.toString()), type: "uint256" },
        { value: toBN(h.toAffine().x.toString()), type: "uint256" },
        { value: toBN(h.toAffine().y.toString()), type: "uint256" },
        { value: toBN(key.pk.toAffine().x.toString()), type: "uint256" },
        { value: toBN(key.pk.toAffine().y.toString()), type: "uint256" },
        { value: toBN(gamma.toAffine().x.toString()), type: "uint256" },
        { value: toBN(gamma.toAffine().y.toString()), type: "uint256" },
        { value: toBN(gToK.toAffine().x.toString()), type: "uint256" },
        { value: toBN(gToK.toAffine().y.toString()), type: "uint256" },
        { value: toBN(hToK.toAffine().x.toString()), type: "uint256" },
        { value: toBN(hToK.toAffine().y.toString()), type: "uint256" }
    );

    const c = BigInt(sha256(toHash)) % bn254.CURVE.n;
    const s = (((k - c * key.sk) % bn254.CURVE.n) + bn254.CURVE.n) % bn254.CURVE.n;
    const proof: Proof = { gamma: gamma, c: c, s: s };

    return proof;
}

function g1YFromX(x: bigint) {
    const ySquare = (x * x * x + BigInt(3)) % bn254.CURVE.p;
    const fp = Field(bn254.CURVE.p);

    try {
        const y = fp.sqrt(ySquare);

        const point1: AffinePoint<bigint> = { x: x, y: y };
        const point2 = bn254.ProjectivePoint.fromAffine(point1);
        return point2;
    } catch (e) {
        return null;
    }
}

function g1HashToPoint(m: string): ProjPointType<bigint> {
    const h = BigInt(sha256(m));
    let x = h % bn254.CURVE.p;
    while (true) {
        const point = g1YFromX(x);
        if (point != null) {
            return point;
        }
        x += BigInt(1);
    }
}
