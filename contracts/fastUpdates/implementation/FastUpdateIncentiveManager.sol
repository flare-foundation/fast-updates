// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import { FastUpdater } from "./FastUpdater.sol";
import "../lib/CircularList.sol" as CL;
import { IIFastUpdateIncentiveManager } from "../interface/IIFastUpdateIncentiveManager.sol";

using { CL.circularGet16, CL.circularHead16, CL.circularZero16, CL.circularAdd16, CL.circularResize, CL.clear, CL.sum } for uint16[];

contract FastUpdateIncentiveManager is IIFastUpdateIncentiveManager {
    address payable rewardPool;

    uint private excessIncentiveValue;

    uint16[] private sampleIncreases;
    uint16[] private precisionIncreases;
    uint16[] private rangeIncreases;

    uint16 private baseSampleSize8x8;
    uint16 private basePrecision0x16;
    uint16 private baseRange8x8;

    uint private rangeIncreasePrice;
    uint private sampleIncreaseLimit;

    function nextSortitionParameters() public override returns (uint16 newSampleSize, uint16 newPrecision) { // only governance
        newSampleSize = baseSampleSize8x8 + sampleIncreases.sum();
        newPrecision = basePrecision0x16 + precisionIncreases.sum();
        sampleIncreases.circularZero16();
        precisionIncreases.circularZero16();
    }

    function incrementExpectedSampleSize(uint16 inc8x8) private {
        sampleIncreases.circularAdd16(inc8x8);
    }

    function incrementPrecision(uint16 inc1x15) private {
        sampleIncreases.circularAdd16(inc1x15);
    }

    function incrementRange(uint16 inc8x8) private {
        rangeIncreases.circularAdd16(inc8x8);
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
        uint16 variationRangeIncrease0x16 = offer.variationRangeIncrease0x16;
        if (offer.rangeLimit8x8 < range8x8 + variationRangeIncrease0x16) {
            uint16 newRangeIncrease = offer.rangeLimit8x8 - range8x8;
            if (newRangeIncrease < 0) newRangeIncrease = 0;
            contribution *= newRangeIncrease / variationRangeIncrease0x16;
            variationRangeIncrease0x16 = newRangeIncrease;
        }
        uint dx = contribution - rangeIncreasePrice * variationRangeIncrease0x16;
        excessIncentiveValue += dx;
        uint16 de = uint16((dx / excessIncentiveValue) / sampleIncreaseLimit);

        expectedSampleSize8x8 += de;
        incrementExpectedSampleSize(de);

        range8x8 += variationRangeIncrease0x16;
        incrementRange(variationRangeIncrease0x16);

        uint16 newPrecision0x16 = range8x8 / expectedSampleSize8x8;
        incrementPrecision(newPrecision0x16 - precision0x16);
        precision0x16 = newPrecision0x16;

        assert(de >= 0);
        payable(msg.sender).transfer(msg.value - contribution);
        rewardPool.transfer(contribution);
    }
}