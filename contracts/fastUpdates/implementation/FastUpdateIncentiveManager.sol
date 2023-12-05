// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import { FastUpdater } from "./FastUpdater.sol";
import "../lib/CircularList.sol" as CL;
import "../lib/FixedPointArithmetic.sol" as FPA;
import { IIFastUpdateIncentiveManager } from "../interface/IIFastUpdateIncentiveManager.sol";

using { CL.circularGet16, CL.circularHead16, CL.circularZero16, CL.circularAdd16, CL.circularResize, CL.clear, CL.sum } for uint16[];

contract FastUpdateIncentiveManager is IIFastUpdateIncentiveManager {
    FPA.SampleSize[] private sampleIncreases;
    FPA.Range[] private rangeIncreases;

    FPA.SampleSize private baseSampleSize;
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

    function computePrecision() view private returns(FPA.Precision) {
        return FPA.div(computeRange(), computeSampleSize());
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
        (uint dc, FPA.Range dr) = processIncentiveOffer(offer);
        FPA.SampleSize de = sampleSizeIncrease(dc, dr);

        incrementExpectedSampleSize(de);
        incrementRange(dr);

        rewardPool.transfer(dc);
        payable(msg.sender).transfer(msg.value - dc);
    }

    function processIncentiveOffer(IncentiveOffer calldata offer) private returns (uint contribution, FPA.Range rangeIncrease) {
        contribution = msg.value;
        rangeIncrease = offer.rangeIncrease;

        FPA.Range rangeNow = computeRange();
        if (FPA.lessThan(offer.rangeLimit, FPA.add(rangeNow, rangeIncrease))) {
            FPA.Range newRangeIncrease = FPA.lessThan(offer.rangeLimit, rangeNow) ? FPA.zeroR : FPA.sub(offer.rangeLimit, rangeNow);
            contribution = FPA.mul(FPA.frac(newRangeIncrease, rangeIncrease), contribution);
            rangeIncrease = newRangeIncrease;
        }
    }

    function sampleSizeIncrease(uint dc, FPA.Range dr) private returns(FPA.SampleSize de) {
        uint rangeCost = FPA.mul(rangeIncreasePrice, dr);
        require(dc >= rangeCost, "Insufficient contribution to pay for range increase");
        uint dx = dc - rangeCost;
        excessIncentiveValue += dx;
        de = FPA.mul(FPA.frac(dx, excessIncentiveValue), sampleIncreaseLimit);
    }
}