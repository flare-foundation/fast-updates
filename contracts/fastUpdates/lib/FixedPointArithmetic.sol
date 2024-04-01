// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

/* 
   Opaque type synonyms to enforce arithemtic correctness.
   All of these are internally uint to avert solc's restricted-bit-size internal handling.
   Since the space is available, the fractional parts of all (except Price, which is not controlled by us) are very wide.
*/

type Scale is uint; // 1x127
type Precision is uint; // 0x127; the fractional part of Scale, top bit always 0
type SampleSize is uint; // 8x120; current gas usage and block gas limit force <32 update transactions per block
type Range is uint; // 8x120, with some space for >100% fluctuations (measured volatility per block is ~1e-3 at most)
type Fractional is uint; // 0x128

type Fee is uint; // 128x0; same scale as currency units, restricted to bottom 128 bits (1e18 integer and fractional parts) to accommodate arithmetic
type Price is uint; // 32x0; an FTSO v2 price is 32-bit, as per the ftso-scaling repo

Scale constant one = Scale.wrap(1 << 127);
SampleSize constant zeroS = SampleSize.wrap(0);
Range constant zeroR = Range.wrap(0);
Fee constant zeroF = Fee.wrap(0);

function _check(uint x) pure returns(bool) {
    return x < 1<<128;
}

function check(Scale x) pure returns(bool) {
    return _check(Scale.unwrap(x));
}

function check(Precision x) pure returns(bool) {
    return _check(Precision.unwrap(x));
}

function check(SampleSize x) pure returns(bool) {
    return _check(SampleSize.unwrap(x));
}

function check(Range x) pure returns(bool) {
    return _check(Range.unwrap(x));
}

function check(Fractional x) pure returns(bool) {
    return _check(Fractional.unwrap(x));
}

function check(Fee x) pure returns(bool) {
    return _check(Fee.unwrap(x));
}

function check(Price x) pure returns(bool) {
    return Price.unwrap(x) < 1<<32;
}

function add(SampleSize x, SampleSize y) pure returns (SampleSize z) {
    unchecked {
        z = SampleSize.wrap(SampleSize.unwrap(x) + SampleSize.unwrap(y));
    }
}

function add(Range x, Range y) pure returns (Range z) {
    unchecked {
        z = Range.wrap(Range.unwrap(x) + Range.unwrap(y));
    }
}

function add(Fee x, Fee y) pure returns (Fee z) {
    unchecked {
        z = Fee.wrap(Fee.unwrap(x) + Fee.unwrap(y));
    }
}

function sub(SampleSize x, SampleSize y) pure returns (SampleSize z) {
    unchecked {
        z = SampleSize.wrap(SampleSize.unwrap(x) - SampleSize.unwrap(y));
    }
}

function sub(Range x, Range y) pure returns (Range z) {
    unchecked {
        z = Range.wrap(Range.unwrap(x) - Range.unwrap(y));
    }
}

function sub(Fee x, Fee y) pure returns (Fee z) {
    unchecked {
        z = Fee.wrap(Fee.unwrap(x) - Fee.unwrap(y));
    }
}

function sum(SampleSize[] storage list) view returns (SampleSize z) {
    unchecked {
        z = zeroS;
        for (uint i = 0; i < list.length; ++i) {
            z = add(z, list[i]);
        }
    }
}

function sum(Range[] storage list) view returns (Range z) {
    unchecked {
        z = zeroR;
        for (uint i = 0; i < list.length; ++i) {
            z = add(z, list[i]);
        }
    }
}

function sum(Fee[] storage list) view returns (Fee z) {
    unchecked {
        for (uint i = 0; i < list.length; ++i) {
            z = add(z, list[i]);
        }
    }
}

function scaleWithPrecision(Precision p) pure returns (Scale s) {
    unchecked {
        return Scale.wrap(Scale.unwrap(one) + Precision.unwrap(p));
    }
}

function lessThan(Range x, Range y) pure returns (bool) {
    unchecked {
        return Range.unwrap(x) < Range.unwrap(y);
    }
}

function lessThan(Fee x, Fee y) pure returns (bool) {
    unchecked {
        return Fee.unwrap(x) < Fee.unwrap(y);
    }
}

function lessThan(Range x, SampleSize y) pure returns (bool) {
    unchecked {
        return Range.unwrap(x) < SampleSize.unwrap(y);
    }
}

function mul(Scale x, Scale y) pure returns(Scale z) {
    unchecked {
        uint xWide = Scale.unwrap(x);
        uint yWide = Scale.unwrap(y);
        uint zWide = (xWide * yWide) >> 127;
        z = Scale.wrap(zWide);
    }
}

function mul(Price x, Scale y) pure returns (Price z) {
    unchecked {
        uint xWide = Price.unwrap(x);
        uint yWide = Scale.unwrap(y);
        uint zWide = (xWide * yWide) >> 127;
        z = Price.wrap(zWide);
    }
}

function mul(Fee x, Range y) pure returns (Fee z) {
    unchecked {
        uint xWide = Fee.unwrap(x);
        uint yWide = Range.unwrap(y);
        uint zWide = (xWide * yWide) >> 120;
        z = Fee.wrap(zWide);
    }
}

function mul(Fractional x, Fee y) pure returns (Fee z) {
    unchecked {
        uint xWide = Fractional.unwrap(x);
        uint yWide = Fee.unwrap(y);
        uint zWide = (xWide * yWide) >> 128;
        z = Fee.wrap(zWide);
    }
}

function mul(Fractional x, SampleSize y) pure returns (SampleSize z) {
    unchecked {
        uint xWide = Fractional.unwrap(x);
        uint yWide = SampleSize.unwrap(y);
        uint zWide = (xWide * yWide) >> 128;
        z = SampleSize.wrap(zWide);
    }
}

function frac(Range x, Range y) pure returns (Fractional z) {
    unchecked {
        uint xWide = Range.unwrap(x) << 128;
        uint yWide = Range.unwrap(y);
        uint zWide = xWide / yWide;
        z = Fractional.wrap(zWide);
    }
}

function frac(Fee x, Fee y) pure returns (Fractional z) {
    unchecked {
        uint xWide = Fee.unwrap(x) << 128;
        uint yWide = Fee.unwrap(y);
        uint zWide = xWide / yWide;
        z = Fractional.wrap(zWide);
    }
}

function div(Range x, SampleSize y) pure returns (Precision z) {
    unchecked {
        uint xWide = Range.unwrap(x) << 127;
        uint yWide = SampleSize.unwrap(y);
        uint zWide = xWide / yWide;
        z = Precision.wrap(zWide);
    }
}

function div(Price x, Scale y) pure returns (Price z) {
    unchecked {
        uint xWide = Price.unwrap(x) << 127;
        uint yWide = Scale.unwrap(y);
        uint zWide = xWide / yWide;
        z = Price.wrap(zWide);
    }
}
