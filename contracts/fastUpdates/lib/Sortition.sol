// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

struct ECPoint {
    int x;
}

struct SortitionRound {
    uint256 seed;
    uint256 scoreCutoff;
}

struct SortitionCredential {
    int x;
}