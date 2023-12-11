// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

// it is just a stub, not a live deployment;
// we are fine with experimental feature
/* solium-disable-next-line */
pragma experimental ABIEncoderV2;

import "../lib/FixedPointArithmetic.sol" as FPA;

contract TestFixedPointArithmetic {
    function identityScaleTest(FPA.Scale x) public pure returns (FPA.Scale y) {
        y = x;
    }
    function identityPrecisionTest(FPA.Precision x) public pure returns (FPA.Precision y) {
        y = x;
    }
    function identitySampleSizeTest(FPA.SampleSize x) public pure returns (FPA.SampleSize y) {
        y = x;
    }
    function identityRangeTest(FPA.Range x) public pure returns (FPA.Range y) {
        y = x;
    }
    function identityPriceTest(FPA.Price x) public pure returns (FPA.Price y) {
        y = x;
    }
    function identityDeltaTest(FPA.Delta x) public pure returns (FPA.Delta y) {
        y = x;
    }
    function identityFractionalTest(FPA.Fractional x) public pure returns (FPA.Fractional y) {
        y = x;
    }
    function identityFeeTest(FPA.Fee x) public pure returns (FPA.Fee y) {
        y = x;
    }

    function oneTest(FPA.Scale x) public pure returns(FPA.Scale y1, FPA.Scale y2) {
        y1 = FPA.mul(FPA.one, x);
        y2 = FPA.mul(x, FPA.one);
    }
    function zeroDTest(FPA.Delta x) public pure returns(FPA.Delta y1, FPA.Delta y2) {
        y1 = FPA.add(FPA.zeroD, x);
        y2 = FPA.add(x, FPA.zeroD);
    }
    function zeroSTest(FPA.SampleSize x) public pure returns(FPA.SampleSize y1, FPA.SampleSize y2) {
        y1 = FPA.add(FPA.zeroS, x);
        y2 = FPA.add(x, FPA.zeroS);
    }
    function zeroRTest(FPA.Range x) public pure returns(FPA.Range y1, FPA.Range y2) {
        y1 = FPA.add(FPA.zeroR, x);
        y2 = FPA.add(x, FPA.zeroR);
    }

    // Addition/subtraction tests

    function addDeltaTest(FPA.Delta x, FPA.Delta y) public pure returns(FPA.Delta z) {
        z = FPA.add(x, y);
    }
    function addSampleSizeTest(FPA.SampleSize x, FPA.SampleSize y) public pure returns(FPA.SampleSize z1, FPA.SampleSize z2) {
        z1 = FPA.add(x, y);
        z2 = FPA.sub(x, y);
    }
    function addRangeTest(FPA.Range x, FPA.Range y) public pure returns(FPA.Range z1, FPA.Range z2) {
        z1 = FPA.add(x, y);
        z2 = FPA.sub(x, y);
    }
    function addFeeTest(FPA.Fee x, FPA.Fee y) public pure returns(FPA.Fee z1, FPA.Fee z2) {
        z1 = FPA.add(x, y);
        z2 = FPA.sub(x, y);
    }

    // Multiplication/division tests

    function mulScaleTest(FPA.Scale x, FPA.Scale y) public pure returns (FPA.Scale z1, FPA.Scale z2) {
        z1 = FPA.mul(x, y);
        z2 = FPA.div(x, y);
    }
    function mulPriceScaleTest(FPA.Price x, FPA.Scale y) public pure returns (FPA.Price z) {
        z = FPA.mul(x, y);
    }
    function mulFeeRangeTest(FPA.Fee x, FPA.Range y) public pure returns (FPA.Fee z) {
        z = FPA.mul(x, y);
    }
    function mulFractionalFeeTest(FPA.Fractional x, FPA.Fee y) public pure returns (FPA.Fee z) {
        z = FPA.mul(x, y);
    }
    function mulFractionalSampleSizeTest(FPA.Fractional x, FPA.SampleSize y) public pure returns (FPA.SampleSize z) {
        z = FPA.mul(x, y);
    }
    function divRangeTest(FPA.Range x, FPA.Range y) public pure returns (FPA.Fractional z) {
        z = FPA.frac(x, y);
    }
    function divFeeTest(FPA.Fee x, FPA.Fee y) public pure returns (FPA.Fractional z) {
        z = FPA.frac(x, y);
    }
    function divRangeSampleSizeTest(FPA.Range x, FPA.SampleSize y) public pure returns (FPA.Precision z) {
        z = FPA.div(x, y);
    }
}