// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import { FastUpdater } from "./FastUpdater.sol";
import "../lib/CircularList.sol" as CL;
import "../lib/FixedPointArithmetic.sol" as FPA;
import { IIFastUpdateIncentiveManager } from "../interface/IIFastUpdateIncentiveManager.sol";

using { CL.circularGet16, CL.circularHead16, CL.circularZero16, CL.circularAdd16, CL.circularResize, CL.clear, CL.sum } for uint16[];

contract FastUpdateIncentiveManager is IIFastUpdateIncentiveManager {
    FPA.SampleSize[] private sampleIncreases;
    FPA.Precision[] private precisionIncreases;
    FPA.Range[] private rangeIncreases;

    FPA.SampleSize private baseSampleSize;
    FPA.Precision private basePrecision;
    FPA.Range private baseRange;

    FPA.SampleSize private sampleIncreaseLimit; // == 1/B
    uint private rangeIncreasePrice; // == A

    uint private excessIncentiveValue; // Must be positive

    function computeSampleSize() view private returns(FPA.SampleSize) {
        return FPA.add(baseSampleSize, FPA.sum(sampleIncreases));
    }

    function getExpectedSampleSize() view external override returns(FPA.SampleSize) {
        return computeSampleSize();
    }

    function computePrecision() view private returns(FPA.Precision p) {
        p = FPA.add(basePrecision, FPA.sum(precisionIncreases));
    }

    function getPrecision() view external override returns(FPA.Precision) {
        return computePrecision();
    }

    function computeRange() view private returns(FPA.Range) {
        return FPA.add(baseRange, FPA.sum(rangeIncreases));
    }

    function getRange() view external override returns(FPA.Range) {
        return computeRange();
    }

    function nextUpdateParameters() public override returns (FPA.SampleSize newSampleSize, FPA.Scale newScale) { // only governance
        newSampleSize = computeSampleSize();
        newScale = FPA.scaleWithPrecision(computePrecision());
        
        // sampleIncreases.circularZero16();
        // precisionIncreases.circularZero16();
        // rangeIncreases.circularZero16();
    }

    function incrementExpectedSampleSize(FPA.SampleSize de) private {
        // sampleIncreases.circularAdd16(inc8x8);
    }

    function incrementPrecision(FPA.Precision dp) private {
        // sampleIncreases.circularAdd16(inc1x15);
    }

    function incrementRange(FPA.Range dr) private {
        // rangeIncreases.circularAdd16(inc8x8);
    }

    function offerIncentive(IncentiveOffer calldata offer) external payable override {
        // let c = total amount received as incentive
        //     r = one-block relative variation range of numeric delta
        //     p = precision of numeric delta
        //     e = expected sample size
        // then: r = (1 + p)^e - 1 = pe (approximate, for small p)
        //       c = A r + C exp(B e)
        // equivalently: dc = A dr + B (c - A r) de
        // Given arguments dr = variationRangeIncrease, dc = msg.value, and current values of r, c: solve for de
        // Then r' = r + dr
        //      e' = e + de
        //      p' = r'/e'
        // and we update e', p'
        // TODO: fixed-point arithmetic

        uint contribution = msg.value;
        FPA.Range rangeIncrease = offer.rangeIncrease;
        require(!FPA.lessThan(rangeIncrease, FPA.zeroR), "Range increase must be nonnegative");
        FPA.Range rangeNow = computeRange();
        if (FPA.lessThan(offer.rangeLimit, FPA.add(rangeNow, rangeIncrease))) {
            FPA.Range newRangeIncrease = FPA.lessThan(offer.rangeLimit, rangeNow) ? FPA.zeroR : FPA.sub(offer.rangeLimit, rangeNow);
            contribution = FPA.mul(FPA.frac(newRangeIncrease, rangeIncrease), contribution);
            rangeIncrease = newRangeIncrease;
        }
        uint dx = contribution - FPA.mul(rangeIncreasePrice, rangeIncrease);
        excessIncentiveValue += dx;
        FPA.SampleSize de = FPA.mul(FPA.frac(dx, excessIncentiveValue), sampleIncreaseLimit);

        require(!FPA.lessThan(de, FPA.zeroS), "Incentive offer must not decrease the sample size");
        incrementExpectedSampleSize(de);
        FPA.SampleSize sampleSizeNow = computeSampleSize();

        rangeNow = FPA.add(rangeNow, rangeIncrease);
        incrementRange(rangeIncrease);

        FPA.Precision precisionNow = computePrecision();
        FPA.Precision newPrecision = FPA.div(rangeNow, sampleSizeNow);
        incrementPrecision(FPA.sub(newPrecision, precisionNow));

        payable(msg.sender).transfer(msg.value - contribution);
        rewardPool.transfer(contribution);
    }
}  