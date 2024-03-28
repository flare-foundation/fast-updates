// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

type Scale is uint; // 1x15, gives price granularity of 5-6 decimal places, more than is used in the FTSO
type Precision is uint; // 0x15; the fractional part of Scale
type SampleSize is uint; // 8x8
type Range is uint; // 8x8
type Price is uint; // 32x0 // An FTSO v2 price is 32-bit, as per the ftso-scaling repo
type Fractional is uint; // 0x16
type Fee is uint; // Same scale as currency units, with restricted bit length

Scale constant one = Scale.wrap(1 << 15); // 1.0000 0000 0000 000 binary
SampleSize constant zeroS = SampleSize.wrap(0);
Range constant zeroR = Range.wrap(0);
Fee constant zeroF = Fee.wrap(0);

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
        uint zWide = (xWide * yWide) >> 15;
        z = Scale.wrap(zWide);
    }
}

function mul(Price x, Scale y) pure returns (Price z) {
    unchecked {
        uint xWide = Price.unwrap(x);
        uint yWide = Scale.unwrap(y);
        uint zWide = (xWide * yWide) >> 15;
        z = Price.wrap(zWide);
    }
}

function mul(Fee x, Range y) pure returns (Fee z) {
    unchecked {
        uint xWide = Fee.unwrap(x);
        uint yWide = Range.unwrap(y);
        uint zWide = (xWide * yWide) >> 8;
        z = Fee.wrap(zWide);
    }
}

function mul(Fractional x, Fee y) pure returns (Fee z) {
    unchecked {
        uint xWide = Fractional.unwrap(x);
        uint yWide = Fee.unwrap(y);
        uint zWide = (xWide * yWide) >> 16;
        z = Fee.wrap(zWide);
    }
}

function mul(Fractional x, SampleSize y) pure returns (SampleSize z) {
    unchecked {
        uint xWide = Fractional.unwrap(x);
        uint yWide = SampleSize.unwrap(y);
        uint zWide = (xWide * yWide) >> 16;
        z = SampleSize.wrap(uint16(zWide));
    }
}

function frac(Range x, Range y) pure returns (Fractional z) {
    unchecked {
        uint xWide = Range.unwrap(x) << 16;
        uint yWide = Range.unwrap(y);
        uint zWide = xWide / yWide;
        z = Fractional.wrap(zWide);
    }
}

function frac(Fee x, Fee y) pure returns (Fractional z) {
    unchecked {
        uint xWide = Fee.unwrap(x) << 16;
        uint yWide = Fee.unwrap(y);
        uint zWide = xWide / yWide;
        z = Fractional.wrap(zWide);
    }
}

function div(Scale x, Scale y) pure returns (Scale z) {
    unchecked {
        uint xWide = Scale.unwrap(x) << 15;
        uint yWide = Scale.unwrap(y);
        uint zWide = xWide / yWide;
        z = Scale.wrap(zWide);
    }
}

function div(Range x, SampleSize y) pure returns (Precision z) {
    unchecked {
        uint xWide = Range.unwrap(x) << 15;
        uint yWide = SampleSize.unwrap(y);
        uint zWide = xWide / yWide;
        z = Precision.wrap(zWide);
    }
}

function div(Price x, Scale y) pure returns (Price z) {
    unchecked {
        uint xWide = Price.unwrap(x) << 15;
        uint yWide = Scale.unwrap(y);
        uint zWide = xWide / yWide;
        z = Price.wrap(zWide);
    }
}
