// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

type Scale is uint16; // 1x15, gives price granularity of 5-6 decimal places, more than is used in the FTSO
type Precision is uint16; // 0x15; the fractional part of Scale
type SampleSize is uint16; // 8x8
type Range is uint16; // 8x8
type Price is uint32; // 32x0 // An FTSO v2 price is 32-bit, as per the ftso-scaling repo
type Fractional is uint16; // 0x16
type Fee is uint240; // Same scale as currency units, with restricted bit length
import "hardhat/console.sol";

Scale constant one = Scale.wrap(1 << 15); // 1.0000 0000 0000 000 binary
SampleSize constant zeroS = SampleSize.wrap(0);
Range constant zeroR = Range.wrap(0);
Fee constant zeroF = Fee.wrap(0);

function add(SampleSize x, SampleSize y) pure returns (SampleSize z) {
    z = SampleSize.wrap(SampleSize.unwrap(x) + SampleSize.unwrap(y));
}

function add(Range x, Range y) pure returns (Range z) {
    z = Range.wrap(Range.unwrap(x) + Range.unwrap(y));
}

function add(Fee x, Fee y) pure returns (Fee z) {
    z = Fee.wrap(Fee.unwrap(x) + Fee.unwrap(y));
}

function sub(SampleSize x, SampleSize y) pure returns (SampleSize z) {
    z = SampleSize.wrap(SampleSize.unwrap(x) - SampleSize.unwrap(y));
}

function sub(Range x, Range y) pure returns (Range z) {
    z = Range.wrap(Range.unwrap(x) - Range.unwrap(y));
}

function sub(Fee x, Fee y) pure returns (Fee z) {
    z = Fee.wrap(Fee.unwrap(x) - Fee.unwrap(y));
}

function sum(SampleSize[] storage list) view returns (SampleSize z) {
    z = zeroS;
    for (uint i = 0; i < list.length; ++i) {
        z = add(z, list[i]);
    }
}

function sum(Range[] storage list) view returns (Range z) {
    z = zeroR;
    for (uint i = 0; i < list.length; ++i) {
        z = add(z, list[i]);
    }
}

function sum(Fee[] storage list) view returns (Fee z) {
    for (uint i = 0; i < list.length; ++i) {
        z = add(z, list[i]);
    }
}

function scaleWithPrecision(Precision p) pure returns (Scale s) {
    return Scale.wrap(Scale.unwrap(one) + Precision.unwrap(p));
}

function lessThan(Range x, Range y) pure returns (bool) {
    return Range.unwrap(x) < Range.unwrap(y);
}

function lessThan(Fee x, Fee y) pure returns (bool) {
    return Fee.unwrap(x) < Fee.unwrap(y);
}

function lessThan(Range x, SampleSize y) pure returns (bool) {
    return Range.unwrap(x) < SampleSize.unwrap(y);
}

function mul(Scale x, Scale y) pure returns(Scale z) {
    uint32 xWide = uint32(Scale.unwrap(x));
    uint32 yWide = uint32(Scale.unwrap(y));
    uint32 zWide = (xWide * yWide) >> 15;
    z = Scale.wrap(uint16(zWide));
}

function mul(Price x, Scale y) pure returns (Price z) {
    uint48 xWide = uint48(Price.unwrap(x));
    uint48 yWide = uint48(Scale.unwrap(y));
    uint48 zWide = (xWide * yWide) >> 15;
    z = Price.wrap(uint32(zWide));
}

function mul(Fee x, Range y) pure returns (Fee z) {
    uint xWide = uint(Fee.unwrap(x));
    uint yWide = uint(Range.unwrap(y));
    uint zWide = (xWide * yWide) >> 8;
    z = Fee.wrap(uint240(zWide));
}

function mul(Fractional x, Fee y) pure returns (Fee z) {
    uint xWide = uint(Fractional.unwrap(x));
    uint yWide = uint(Fee.unwrap(y));
    uint zWide = (xWide * yWide) >> 16;
    z = Fee.wrap(uint240(zWide));
}

function mul(Fractional x, SampleSize y) pure returns (SampleSize z) {
    uint32 xWide = uint32(Fractional.unwrap(x));
    uint32 yWide = uint32(SampleSize.unwrap(y));
    uint32 zWide = (xWide * yWide) >> 16;
    z = SampleSize.wrap(uint16(zWide));
}

function frac(Range x, Range y) pure returns (Fractional z) {
    uint32 xWide = uint32(Range.unwrap(x)) << 16;
    uint32 yWide = uint32(Range.unwrap(y));
    uint32 zWide = xWide / yWide;
    z = Fractional.wrap(uint16(zWide));
}

function frac(Fee x, Fee y) pure returns (Fractional z) {
    uint xWide = uint(Fee.unwrap(x)) << 16;
    uint yWide = uint(Fee.unwrap(y));
    uint zWide = xWide / yWide;
    z = Fractional.wrap(uint16(zWide));
}

function div(Scale x, Scale y) pure returns (Scale z) {
    uint32 xWide = uint32(Scale.unwrap(x)) << 15;
    uint32 yWide = uint32(Scale.unwrap(y));
    uint32 zWide = xWide / yWide;
    z = Scale.wrap(uint16(zWide));
}

function div(Range x, SampleSize y) pure returns (Precision z) {
    uint32 xWide = uint32(Range.unwrap(x)) << 15;
    uint32 yWide = uint32(SampleSize.unwrap(y));
    uint32 zWide = xWide / yWide;
    z = Precision.wrap(uint16(zWide));
}

function div(Price x, Scale y) pure returns (Price z) {
    uint48 xWide = uint48(Price.unwrap(x)) << 15;
    uint48 yWide = uint48(Scale.unwrap(y));
    uint48 zWide = xWide / yWide;
    z = Price.wrap(uint32(zWide));
}
