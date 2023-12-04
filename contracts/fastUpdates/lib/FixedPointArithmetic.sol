// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

type Scale is uint16; // 1x15, gives price granularity of 5-6 decimal places, more than is used in the FTSO
type Precision is uint16; // 0x15; the fractional part of Scale
type SampleSize is uint16; // 7x8
type Range is uint16; // 8x8
type Price is uint32; // 32x0
type Delta is int8; // 7x0 signed

Scale constant one = Scale.wrap(1 << 15); // 1.0000 0000 0000 000 binary
Precision constant zeroP = Precision.wrap(0);
Delta constant zeroD = Delta.wrap(0);
SampleSize constant zeroS = SampleSize.wrap(0);
Range constant zeroR = Range.wrap(0);

function add(Delta x, Delta y) pure returns (Delta z) {
    z = Delta.wrap(Delta.unwrap(x) + Delta.unwrap(y));
}

function add(SampleSize x, SampleSize y) pure returns (SampleSize z) {
    z = SampleSize.wrap(SampleSize.unwrap(x) + SampleSize.unwrap(y));
}

function add(Precision x, Precision y) pure returns (Precision z) {
    z = Precision.wrap(Precision.unwrap(x) + Precision.unwrap(y));
}

function add(Range x, Range y) pure returns (Range z) {
    z = Range.wrap(Range.unwrap(x) + Range.unwrap(y));
}

function sub(Delta x, Delta y) pure returns (Delta z) {
    z = Delta.wrap(Delta.unwrap(x) - Delta.unwrap(y));
}

function sub(SampleSize x, SampleSize y) pure returns (SampleSize z) {
    z = SampleSize.wrap(SampleSize.unwrap(x) - SampleSize.unwrap(y));
}

function sub(Precision x, Precision y) pure returns (Precision z) {
    z = Precision.wrap(Precision.unwrap(x) - Precision.unwrap(y));
}

function sub(Range x, Range y) pure returns (Range z) {
    z = Range.wrap(Range.unwrap(x) - Range.unwrap(y));
}

function sum(SampleSize[] storage list) view returns (SampleSize z) {
    z = zeroS;
    for (uint i = 0; i < list.length; ++i) {
        z = add(z, list[i]);
    }
}

function sum(Precision[] storage list) view returns (Precision z) {
    z = zeroP;
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

function scaleWithPrecision(Precision p) pure returns (Scale s) {
    return Scale.wrap(Precision.unwrap(add(Precision.wrap(1), p)));
}

function maxDelta(Delta x) pure returns (bool) {
    return Delta.unwrap(x) == type(int8).max;
}

function minDelta(Delta x) pure returns (bool) {
    return Delta.unwrap(x) == type(int8).min;
}

function positive(Precision x) pure returns (bool) {
    return Precision.unwrap(x) > 0;
}

function lessThan(Range x, Range y) pure returns (bool) {
    return Range.unwrap(x) < Range.unwrap(y);
}

function lessThan(SampleSize x, SampleSize y) pure returns (bool) {
    return SampleSize.unwrap(x) < SampleSize.unwrap(y);
}

function mul(Scale x, Scale y) pure returns(Scale z) {
    uint32 xWide = uint32(Scale.unwrap(x));
    uint32 yWide = uint32(Scale.unwrap(y));
    uint32 zWide = (xWide * yWide) >> 15;
    z = Scale.wrap(uint16(zWide));
}

function mul(Price x, Scale y) pure returns (Price z) {
    uint32 xWide = Price.unwrap(x);
    uint32 yWide = uint32(Scale.unwrap(y));
    uint48 zWide = (xWide * yWide) >> 15;
    z = Price.wrap(uint32(zWide));
}

function div(Scale x, Scale y) pure returns (Scale z) {
    uint32 xWide = uint32(Scale.unwrap(x)) << 15;
    uint32 yWide = uint32(Scale.unwrap(y));
    uint32 zWide = (xWide / yWide) >> 15;
    z = Scale.wrap(uint16(zWide));
}

function div(Range x, Range y) pure returns (Range z) {
    uint32 xWide = uint32(Range.unwrap(x)) << 8;
    uint32 yWide = uint32(Range.unwrap(y));
    uint32 zWide = (xWide / yWide) >> 8;
    z = Range.wrap(uint16(zWide));
}

function div(Range x, SampleSize y) pure returns (Precision z) {
    uint32 xWide = uint32(Range.unwrap(x)) << 8;
    uint32 yWide = uint32(Range.unwrap(y));
    uint32 zWide = (xWide / yWide) >> 8;
    z = Range.wrap(uint16(zWide));
}

function pow(Scale[8] storage binaryPowers, Delta _power) view returns (Scale result) {
    int8 power = Delta.unwrap(_power); // 2s-complement, big-endian indexing
    result = one;
    for(uint i = 0; i < 7; ++i) {
        if (power & 1 != 0) result = mul(result, binaryPowers[i]);
        power >>= 1;
    }
    if (power & 1 != 0) result = div(result, binaryPowers[7]);
}

function powersInto(Scale x, Scale[8] storage binaryPowers) {
    Scale y = binaryPowers[0] = x;
    for (uint i = 1; i < 8; ++i) {
        y = mul(y, y);
        binaryPowers[i] = y;
    }
}