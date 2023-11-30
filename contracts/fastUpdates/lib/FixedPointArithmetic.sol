// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

type Precision is uint16; // 1x15
type SampleSize is uint16; // 7x8
type Range is uint16; // 8x8
type Price is uint32; // 32x0
type Delta is int8; // 7x0 signed

Precision constant one = Precision.wrap(1 << 15); // 1.0000 0000 0000 000 binary

function mul(Precision x, Precision y) pure returns(Precision z) {
    uint32 xWide = uint32(Precision.unwrap(x));
    uint32 yWide = uint32(Precision.unwrap(y));
    uint32 zWide = (xWide * yWide) >> 15;
    z = Precision.wrap(uint16(zWide));
}

function mul(Price x, Precision y) pure returns (Price z) {
    uint32 xWide = Price.unwrap(x);
    uint32 yWide = uint32(Precision.unwrap(y));
    uint48 zWide = (xWide * yWide) >> 15;
    z = Price.wrap(uint32(zWide));
}

function div(Precision x, Precision y) pure returns (Precision z) {
    uint32 xWide = uint32(Precision.unwrap(x)) << 15;
    uint32 yWide = uint32(Precision.unwrap(y));
    uint32 zWide = (xWide / yWide) >> 15;
    z = Precision.wrap(uint16(zWide));
}

function pow(Precision[8] storage binaryPowers, Delta power) view returns (Precision result) {
    bytes1 powerBits = bytes1(uint8(Delta.unwrap(power))); // big-endian indexing
    bytes1 mask = bytes1(hex"01");
    result = one;
    for(uint i = 0; i < 6; ++i) {
        if (powerBits & mask != bytes1(0)) result = mul(result, binaryPowers[i]);
        powerBits >>= 1;
        ++i;
    }
    if (powerBits & mask != bytes1(0)) result = div(result, binaryPowers[7]);
}