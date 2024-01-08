// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import "hardhat/console.sol";
import {Bn256} from "./Bn256.sol";

// Encoding of EC point when space is premium
struct ECPoint {
    uint256 x;
    bool signed;
}

struct SortitionRound {
    bool present;
    uint seed;
    uint scoreCutoff;
}

function sortitionRound(uint seed, uint scoreCutoff) pure returns (SortitionRound memory) {
    return SortitionRound(true, seed, scoreCutoff);
}

function g1SignedPointToG1Point(ECPoint memory ecPt) view returns (Bn256.G1Point memory pt) {
    pt.x = ecPt.x;
    pt.y = Bn256.g1YFromX(ecPt.x);
    if (ecPt.signed) {
        pt.y = Bn256.p - pt.y;
    }
}

struct SortitionCredential {
    uint256 replicate;
    Bn256.G1Point gamma;
    uint256 c;
    uint256 s;
}

function verifySortitionCredential(
    SortitionRound memory round,
    Bn256.G1Point memory pubKey,
    uint weight,
    SortitionCredential memory sortitionCredential
) view returns (bool, uint256) {
    require(sortitionCredential.replicate < weight, "Credential's replicate value is not less than provider's weight");
    bool check = verifySortitionProof(round.seed, pubKey, sortitionCredential);
    uint256 vrfVal = sortitionCredential.gamma.x;

    return (check && vrfVal <= round.scoreCutoff, vrfVal);
}

function verifySortitionProof(
    uint256 seed,
    Bn256.G1Point memory pubKey,
    SortitionCredential memory sortitionCredential
) view returns (bool) {
    require(Bn256.isG1PointOnCurve(pubKey)); // this also checks that it is not zero
    require(Bn256.isG1PointOnCurve(sortitionCredential.gamma));
    Bn256.G1Point memory u = Bn256.g1Add(
        Bn256.scalarMultiply(pubKey, sortitionCredential.c),
        Bn256.scalarMultiply(Bn256.g1(), sortitionCredential.s)
    );

    Bn256.G1Point memory h = Bn256.g1HashToPoint(abi.encodePacked(seed, sortitionCredential.replicate));

    Bn256.G1Point memory v = Bn256.g1Add(
        Bn256.scalarMultiply(sortitionCredential.gamma, sortitionCredential.c),
        Bn256.scalarMultiply(h, sortitionCredential.s)
    );
    uint256 c2 = uint256(sha256(abi.encode(Bn256.g1(), h, pubKey, sortitionCredential.gamma, u, v)));
    c2 = c2 % Bn256.getQ();

    return c2 == sortitionCredential.c;
}
