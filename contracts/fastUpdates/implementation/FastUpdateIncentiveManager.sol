// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import { IICircular } from "../interface/IICircular.sol";
import "../lib/FixedPointArithmetic.sol" as FPA;
import { IIFastUpdateIncentiveManager } from "../interface/IIFastUpdateIncentiveManager.sol";
import "hardhat/console.sol";

contract FastUpdateIncentiveManager is IIFastUpdateIncentiveManager, IICircular {
    // Circular lists all
    FPA.SampleSize[] private sampleIncreases;
    FPA.Range[] private rangeIncreases;
    FPA.Fee[] private excessOfferIncreases;

    function getIncentiveDuration() external view override returns(uint) {
        return circularLength;
    }

    function setIncentiveDuration(uint _duration) public override { // only governance
        delete sampleIncreases;
        delete rangeIncreases;
        delete excessOfferIncreases;

        for (uint i = 0; i < _duration; ++i) {
            sampleIncreases.push();
            rangeIncreases.push();
            excessOfferIncreases.push();
        }
    }

    // This is an optimization to prevent recalculation of these numbers in every offer.
    // Extra bookkeeping is required.
    FPA.SampleSize sampleSize;
    FPA.Range range;
    FPA.Fee excessOfferValue;

    constructor(address payable _rp, FPA.SampleSize _bss, FPA.Range _br, FPA.SampleSize _sil, FPA.Fee _rip, uint _dur) 
        IIFastUpdateIncentiveManager(_rp, _bss, _br, _sil, _rip, _dur)
        IICircular(_dur)
    {
        setIncentiveDuration(_dur);
        sampleSize = baseSampleSize;
        range = baseRange;
        excessOfferValue = FPA.Fee.wrap(1); // Arbitrary initial value, but must not be 0
    }

    function getExpectedSampleSize() view external override returns(FPA.SampleSize) {
        return sampleSize;
    }

    function computePrecision() view private returns(FPA.Precision) {
        return FPA.div(range, sampleSize);
    }

    function getPrecision() view external override returns(FPA.Precision) {
        return computePrecision();
    }

    function getRange() view external override returns(FPA.Range) {
        return range;
    }

    // This is expected to be called only by FastUpdater, and only at the end of a block.
    function nextUpdateParameters() public override returns (FPA.SampleSize newSampleSize, FPA.Scale newScale) { // only governance
        newSampleSize = sampleSize;
        newScale = FPA.scaleWithPrecision(computePrecision());

        // Bookkeeping for the cached values
        excessOfferValue = FPA.sub(excessOfferValue, excessOfferIncreases[nextIx()]);
        range = FPA.sub(range, rangeIncreases[nextIx()]);
        sampleSize = FPA.sub(sampleSize, sampleIncreases[nextIx()]);
        
        sampleIncreases[nextIx()] = FPA.zeroS;
        rangeIncreases[nextIx()] = FPA.zeroR;
        excessOfferIncreases[nextIx()] = FPA.zeroF;
    }

    function offerIncentive(IncentiveOffer calldata offer) external payable override {
        (FPA.Fee dc, FPA.Range dr) = processIncentiveOffer(offer);
        FPA.SampleSize de = sampleSizeIncrease(dc, dr);

        sampleIncreases[thisIx()] = FPA.add(sampleIncreases[thisIx()], de);
        sampleSize = FPA.add(sampleSize, de);

        rangeIncreases[thisIx()] = FPA.add(rangeIncreases[thisIx()], dr);
        range = FPA.add(range, dr);

        require(FPA.lessThan(offer.rangeLimit, sampleSize), "Offer would make the precision greater than 100%");

        rewardPool.transfer(FPA.Fee.unwrap(dc));
        payable(msg.sender).transfer(msg.value - FPA.Fee.unwrap(dc));
    }

    function processIncentiveOffer(IncentiveOffer calldata offer) private returns (FPA.Fee contribution, FPA.Range rangeIncrease) {
        require(msg.value >> 240 == 0, "Incentive offer value capped at 240 bits");
        contribution = FPA.Fee.wrap(uint240(msg.value));
        rangeIncrease = offer.rangeIncrease;

        if (FPA.lessThan(offer.rangeLimit, FPA.add(range, rangeIncrease))) {
            FPA.Range newRangeIncrease = FPA.lessThan(offer.rangeLimit, range) ? FPA.zeroR : FPA.sub(offer.rangeLimit, range);
            contribution = FPA.mul(FPA.frac(newRangeIncrease, rangeIncrease), contribution);
            rangeIncrease = newRangeIncrease;
        }
    }

    function sampleSizeIncrease(FPA.Fee dc, FPA.Range dr) private returns(FPA.SampleSize de) {
        FPA.Fee rangeCost = FPA.mul(rangeIncreasePrice, dr);
        require(!FPA.lessThan(dc, rangeCost), "Insufficient contribution to pay for range increase");
        FPA.Fee dx = FPA.sub(dc, rangeCost);

        excessOfferIncreases[thisIx()] = FPA.add(excessOfferIncreases[thisIx()], dx);
        excessOfferValue = FPA.add(excessOfferValue, dx);

        de = FPA.mul(FPA.frac(dx, excessOfferValue), sampleIncreaseLimit);
    }
}