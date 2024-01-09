// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import { IICircular } from "../interface/IICircular.sol";
import "../lib/FixedPointArithmetic.sol" as FPA;
import { IIFastUpdateIncentiveManager } from "../interface/IIFastUpdateIncentiveManager.sol";
import "hardhat/console.sol";

abstract contract IncreaseManager is IICircular {
    // Circular lists all
    FPA.SampleSize[] private sampleIncreases;
    FPA.Range[] private rangeIncreases;
    FPA.Fee[] private excessOfferIncreases;

    // This is an optimization to prevent recalculation of these numbers in every offer.
    // Extra bookkeeping is required.
    FPA.SampleSize sampleSize;
    FPA.Range range;
    FPA.Fee excessOfferValue;

    constructor(uint _dur)
        IICircular(_dur)
    {
        init();
    }

    function init() internal {
        delete sampleIncreases;
        delete rangeIncreases;
        delete excessOfferIncreases;

        for (uint i = 0; i < circularLength; ++i) {
            sampleIncreases.push();
            rangeIncreases.push();
            excessOfferIncreases.push();
        }
    }

    function step() internal {
        // Bookkeeping for the cached values
        excessOfferValue = FPA.sub(excessOfferValue, excessOfferIncreases[nextIx()]);
        range = FPA.sub(range, rangeIncreases[nextIx()]);
        sampleSize = FPA.sub(sampleSize, sampleIncreases[nextIx()]);
        
        sampleIncreases[nextIx()] = FPA.zeroS;
        rangeIncreases[nextIx()] = FPA.zeroR;
        excessOfferIncreases[nextIx()] = FPA.zeroF;
    }

    function increaseSampleSize(FPA.SampleSize de) internal {
        sampleIncreases[thisIx()] = FPA.add(sampleIncreases[thisIx()], de);
        sampleSize = FPA.add(sampleSize, de);
    }

    function increaseRange(FPA.Range dr) internal {
        rangeIncreases[thisIx()] = FPA.add(rangeIncreases[thisIx()], dr);
        range = FPA.add(range, dr);
    }

    function increaseExcessOfferValue(FPA.Fee dx) internal {
        excessOfferIncreases[thisIx()] = FPA.add(excessOfferIncreases[thisIx()], dx);
        excessOfferValue = FPA.add(excessOfferValue, dx);
    }
}

contract FastUpdateIncentiveManager is IIFastUpdateIncentiveManager, IncreaseManager {
    constructor(address payable _rp, FPA.SampleSize _bss, FPA.Range _br, FPA.SampleSize _sil, FPA.Fee _rip, uint _dur) 
        IIFastUpdateIncentiveManager(_rp, _bss, _br, _sil, _rip, _dur)
        IncreaseManager(_dur)
    {
        sampleSize = baseSampleSize;
        range = baseRange;
        excessOfferValue = FPA.Fee.wrap(1); // Arbitrary initial value, but must not be 0
    }

    function setIncentiveDuration(uint _dur) public override {
        setCircularLength(_dur);
        IncreaseManager.init();
    }

    function getIncentiveDuration() external view override returns(uint) {
        return circularLength;
    }

    function getExpectedSampleSize() view external override returns(FPA.SampleSize) {
        return sampleSize;
    }

    function getRange() view external override returns(FPA.Range) {
        return range;
    }

    function getPrecision() view external override returns(FPA.Precision) {
        return computePrecision();
    }

    function computePrecision() view private returns(FPA.Precision) {
        return FPA.div(range, sampleSize);
    }

    // This is expected to be called only by FastUpdater, and only at the end of a block.
    function nextUpdateParameters() public override returns (FPA.SampleSize newSampleSize, FPA.Scale newScale) { // only governance
        newSampleSize = sampleSize;
        newScale = FPA.scaleWithPrecision(computePrecision());
        IncreaseManager.step();
    }

    function offerIncentive(IncentiveOffer calldata offer) external payable override {
        (FPA.Fee dc, FPA.Range dr) = processIncentiveOffer(offer);
        FPA.SampleSize de = sampleSizeIncrease(dc, dr);

        IncreaseManager.increaseSampleSize(de);
        IncreaseManager.increaseRange(dr);

        rewardPool.transfer(FPA.Fee.unwrap(dc));
        emit IncentiveOffered(dr, de, dc);
        payable(msg.sender).transfer(msg.value - FPA.Fee.unwrap(dc));
    }

    function processIncentiveOffer(IncentiveOffer calldata offer) private returns (FPA.Fee contribution, FPA.Range rangeIncrease) {
        require(msg.value >> 240 == 0, "Incentive offer value capped at 240 bits");
        contribution = FPA.Fee.wrap(uint240(msg.value));
        rangeIncrease = offer.rangeIncrease;

        FPA.Range finalRange = FPA.add(range, rangeIncrease);
        if (FPA.lessThan(offer.rangeLimit, finalRange)) {
            finalRange = offer.rangeLimit;
            FPA.Range newRangeIncrease = FPA.lessThan(finalRange, range) ? FPA.zeroR : FPA.sub(finalRange, range);
            contribution = FPA.mul(FPA.frac(newRangeIncrease, rangeIncrease), contribution);
            rangeIncrease = newRangeIncrease;
        }
        require(FPA.lessThan(finalRange, sampleSize), "Offer would make the precision greater than 100%");
    }

    function sampleSizeIncrease(FPA.Fee dc, FPA.Range dr) private returns(FPA.SampleSize de) {
        FPA.Fee rangeCost = FPA.mul(rangeIncreasePrice, dr);
        require(!FPA.lessThan(dc, rangeCost), "Insufficient contribution to pay for range increase");
        FPA.Fee dx = FPA.sub(dc, rangeCost);

        IncreaseManager.increaseExcessOfferValue(dx);

        de = FPA.mul(FPA.frac(dx, excessOfferValue), sampleIncreaseLimit);
    }
}