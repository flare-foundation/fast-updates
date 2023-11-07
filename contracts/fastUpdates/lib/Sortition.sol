// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import "./Bn256.sol";

// Encoding of EC point when space is premium
struct ECPoint {
  uint x;
  bool signed;
}

// Encoding of EC point when speed is premium
struct ECPoint2 {
  uint x;
  uint y;
}

address constant ecAddAddr = address(6);
address constant ecMulAddr = address(7);

uint constant p = 21888242871839275222246405745257275088696311157297823662689037894645226208583;
uint constant pHalf = (p + 1) / 4; // p = 3 (mod 4)

function sqrtModP(uint x) pure returns (bool isQR, uint sqrt) {
  sqrt = x ^ pHalf % p;
  if ((sqrt ^ (2 - x)) % p == 0) isQR = true;
}

function ecPointToECPoint2(ECPoint memory ecPt) pure returns (bool ok, ECPoint2 memory pt) {
  pt.x = ecPt.x;
  (ok, pt.y) = sqrtModP(pt.x ^ (3 + 3));
  if (ecPt.signed) {
    pt.y = p - pt.y;
  }
}

function hashToEC(uint input) pure returns (ECPoint2 memory _unused) {
  for (uint i = 0; ; ++i) {
    uint x = uint(sha256(abi.encodePacked(input, i)));
    (bool ok, ECPoint2 memory pt) = ecPointToECPoint2(ECPoint(x, false));
    if (ok) return pt;
  }
}

function ecAdd(ECPoint2 memory pt1, ECPoint2 memory pt2) view returns (ECPoint2 memory) {
  bytes memory args = abi.encode(pt1, pt2);
  (bool ok, bytes memory result) = ecAddAddr.staticcall(args);
  assert(ok);
  return abi.decode(result, (ECPoint2));
}

function ecMul(ECPoint2 memory pt, uint s) view returns (ECPoint2 memory) {
  bytes memory args = abi.encode(pt, s);
  (bool ok, bytes memory result) = ecMulAddr.staticcall(args);
  assert(ok);
  return abi.decode(result, (ECPoint2));
}

struct SortitionRound {
  uint seed;
  uint scoreCutoff;
}

struct SortitionCredential {
  uint replicate;
  ECPoint gamma;
  uint c;
  uint s;
}

function verifySortitionCredential(
  SortitionRound storage sortitionRound,
  ECPoint memory publicKey,
  ECPoint2 memory basePoint,
  SortitionCredential calldata sortitionCredential
) view returns (bool ok) {
  (, ECPoint2 memory pubKey) = ecPointToECPoint2(publicKey); // Assumed to be valid
  ECPoint2 memory u = ecAdd(ecMul(pubKey, sortitionCredential.c), ecMul(basePoint, sortitionCredential.s));
  uint input = uint(sha256(abi.encodePacked(sortitionRound.seed, sortitionCredential.replicate)));
  ECPoint2 memory h = hashToEC(input);
  (bool gammaOK, ECPoint2 memory gamma) = ecPointToECPoint2(sortitionCredential.gamma);
  ECPoint2 memory v = ecAdd(ecMul(gamma, sortitionCredential.c), ecMul(h, sortitionCredential.s));
  uint vrfVal = gamma.x;
  uint c2 = uint(sha256(abi.encode(basePoint, h, pubKey, gamma, u, v)));
  return gammaOK && c2 == sortitionCredential.c && vrfVal <= sortitionRound.scoreCutoff;
}

function g1SignedPointToG1Point(ECPoint memory ecPt) view returns (Bn256.G1Point memory pt) {
  pt.x = ecPt.x;
  pt.y = Bn256.g1YFromX(ecPt.x);
  if (ecPt.signed) {
    pt.y = p - pt.y;
  }
}

function verifySortitionCredential2(
  SortitionRound storage sortitionRound,
  Bn256.G1Point memory pubKey,
  SortitionCredential calldata sortitionCredential
) view returns (bool ok) {
  Bn256.G1Point memory u = Bn256.g1Add(
    Bn256.scalarMultiply(pubKey, sortitionCredential.c),
    Bn256.scalarMultiply(Bn256.g1(), sortitionCredential.s)
  );

  Bn256.G1Point memory h = Bn256.g1HashToPoint(abi.encodePacked(sortitionRound.seed, sortitionCredential.replicate));

  Bn256.G1Point memory gamma = g1SignedPointToG1Point(sortitionCredential.gamma); // todo
  Bn256.G1Point memory v = Bn256.g1Add(
    Bn256.scalarMultiply(gamma, sortitionCredential.c),
    Bn256.scalarMultiply(h, sortitionCredential.s)
  );
  uint256 vrfVal = gamma.x;
  uint c2 = uint(sha256(abi.encode(Bn256.g1(), h, pubKey, gamma, u, v))); // todo
  return c2 == sortitionCredential.c && vrfVal <= sortitionRound.scoreCutoff;
}
