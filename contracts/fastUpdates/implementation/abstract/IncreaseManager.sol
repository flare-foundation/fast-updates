// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;
pragma abicoder v2;

import { CircularListManager } from "./CircularListManager.sol";
import "../../lib/FixedPointArithmetic.sol" as FPA;

abstract contract IncreaseManager is CircularListManager {
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
        CircularListManager(_dur)
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

    function getIncentiveDuration() external view returns(uint) {
        return circularLength;
    }

    function setIncentiveDuration(uint _dur) external {
        setCircularLength(_dur);
        IncreaseManager.init();
    }
}