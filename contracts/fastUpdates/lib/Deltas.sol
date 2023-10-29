// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

// An array of 1000 2-bit entries, packed
struct Deltas {
    bytes32[7] mainParts;
    bytes26 tailPart;
}

// f is applied to each delta with index, as though mapping over an array
function forEach(Deltas calldata deltas, function(int, uint) f) {
    for (uint i = 0; i < 7; ++i) {
        deltas.mainParts[i].forEachPackedBytes32n(i, f, 32);
    }
    bytes32(deltas.tailPart).forEachPackedBytes32n(7, f, 26);
}

// n is the number of bytes to operate on, starting from index 0
function forEachPackedBytes32n(bytes32 packedBytes, uint i, function(int, uint) f, uint n) {
    i *= 32;
    for (uint j = 0; j < n; ++j) {
        packedBytes[j].forEachPackedBits2(i + j, f);
    }
}

// f is applied to the signed 2-bit integers packed into a bytes1
// The value -2 is rejected.
function forEachPackedBits2(bytes1 packedBits2, uint ij, function(int, uint) f) {
    ij *= 4;
    for (uint k = 0; k < 8; k += 2) {
        int entry = int8(uint8(packedBits2 << k)) >> 6;
        assert(entry != -2);
        f(entry, ij + k);
    }
}

using { forEach } for Deltas global;
using { forEachPackedBytes32n } for bytes32;
using { forEachPackedBits2 } for bytes1;