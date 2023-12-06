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
    FPA.Fee[] private excessOfferIncreases;

    // This is an optimization to prevent recalculation of this number in every offer.
    // Extra bookkeeping is required.
    FPA.Range range;

    // This is an optimization to prevent recalculation of this number in every offer.
    // Extra bookkeeping is required.
    FPA.Fee excessOfferValue;

    constructor(address payable _rp, FPA.SampleSize _bss, FPA.Range _br, FPA.SampleSize _sil, FPA.Fee _rip) {
        rewardPool = _rp;
        baseSampleSize = _bss;
        baseRange = range = _br;
        sampleIncreaseLimit = _sil;
        rangeIncreasePrice = _rip;
        excessOfferValue = FPA.Fee.wrap(1); // Arbitrary initial value, but must not be 0
    }

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

    // This is expected to be called only by FastUpdater, and only at the end of a block.
    function nextUpdateParameters() public override returns (FPA.SampleSize newSampleSize, FPA.Scale newScale) { // only governance
        newSampleSize = computeSampleSize();
        newScale = FPA.scaleWithPrecision(computePrecision());
        excessOfferValue = FPA.sub(excessOfferValue, excessOfferIncreases[0]);
        
        // sampleIncreases.circularZero16();
        // precisionIncreases.circularZero16();
        // rangeIncreases.circularZero16();
    }

    function incrementExpectedSampleSize(FPA.SampleSize de) private {
        // sampleIncreases.circularAdd16(inc8x8);
    }

    function incrementRange(FPA.Range dr) private {
        // rangeIncreases.circularAdd16(inc8x8);
    }

    function incrementExcessOffer(FPA.Fee dx) private {
        excessOfferValue = FPA.add(excessOfferValue, dx);
        // more...
    }

    function offerIncentive(IncentiveOffer calldata offer) external payable override {
        (FPA.Fee dc, FPA.Range dr) = processIncentiveOffer(offer);
        FPA.SampleSize de = sampleSizeIncrease(dc, dr);

        incrementExpectedSampleSize(de);
        incrementRange(dr);

        rewardPool.transfer(FPA.Fee.unwrap(dc));
        payable(msg.sender).transfer(msg.value - FPA.Fee.unwrap(dc));
    }

    function processIncentiveOffer(IncentiveOffer calldata offer) private returns (FPA.Fee contribution, FPA.Range rangeIncrease) {
        require(msg.value >> 240 == 0, "Incentive offer value capped at 240 bits");
        contribution = FPA.Fee.wrap(uint240(msg.value));
        rangeIncrease = offer.rangeIncrease;

        FPA.Range rangeNow = computeRange();
        if (FPA.lessThan(offer.rangeLimit, FPA.add(rangeNow, rangeIncrease))) {
            FPA.Range newRangeIncrease = FPA.lessThan(offer.rangeLimit, rangeNow) ? FPA.zeroR : FPA.sub(offer.rangeLimit, rangeNow);
            contribution = FPA.mul(FPA.frac(newRangeIncrease, rangeIncrease), contribution);
            rangeIncrease = newRangeIncrease;
        }
    }

    function sampleSizeIncrease(FPA.Fee dc, FPA.Range dr) private returns(FPA.SampleSize de) {
        FPA.Fee rangeCost = FPA.mul(rangeIncreasePrice, dr);
        require(!FPA.lessThan(rangeCost, dc), "Insufficient contribution to pay for range increase");
        FPA.Fee dx = FPA.sub(dc, rangeCost);
        incrementExcessOffer(dx);
        de = FPA.mul(FPA.frac(dx, excessOfferValue), sampleIncreaseLimit);
    }
}