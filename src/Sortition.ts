import { RandInt } from "./utils/rand";
import { bn254 } from "@noble/curves/bn254"; // also known as alt_bn128
import { Field } from "@noble/curves/abstract/modular"; // also known as alt_bn128
import { ProjPointType, AffinePoint } from "ccxt/js/src/static_dependencies/noble-curves/abstract/weierstrass";
import { encodePacked } from "web3-utils";
import { toBN } from "./protocol/utils/voting-utils";
import { sha256 } from "ethers";

const p = BigInt("21888242871839275222246405745257275088696311157297823662689037894645226208583");

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
    const sk = RandInt(31);
    const pk = bn254.ProjectivePoint.BASE.multiply(sk);
    const key: SortitionKey = { sk: sk, pk: pk };

    return key;
}

export function VerifiableRandomness(key: SortitionKey, seed: bigint, replicate: bigint): Proof {
    let toHash: string = encodePacked(
        { value: toBN(seed.toString()), type: "uint256" },
        { value: toBN(replicate.toString()), type: "uint256" }
    );

    console.log(toHash);

    const h = g1HashToPoint(toHash);
    console.log(h);
    console.log("hhh", key.sk);
    const gamma = bn254.ProjectivePoint.BASE.multiply(key.sk);

    const k = RandInt(31);
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
    console.log(bn254.CURVE.p, p);
    const c = BigInt(sha256(toHash)) % bn254.CURVE.n;
    console.log(c);
    const s = (k - c * key.sk) % bn254.CURVE.n;
    const proof: Proof = { gamma: gamma, c: c, s: s };

    return proof;
}

function g1YFromX(x: bigint) {
    const ySquare = (x * x * x + BigInt(3)) % p;
    const fp = Field(p);

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
    console.log(h);
    let x = h % bn254.CURVE.p;
    console.log(x);
    while (true) {
        const point = g1YFromX(x);
        if (point != null) {
            return point;
        }
        x += BigInt(1);
    }
}
