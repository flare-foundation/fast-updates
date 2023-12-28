// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import "../lib/FixedPointArithmetic.sol" as FPA;

abstract contract IFastUpdateIncentiveManager {
    function getExpectedSampleSize() view external virtual returns(FPA.SampleSize);
    function getPrecision() view external virtual returns(FPA.Precision);
    function getRange() view external virtual returns(FPA.Range);

    uint public incentiveDuration;

    struct IncentiveOffer {
        FPA.Range rangeIncrease;
        FPA.Range rangeLimit;
    }

    event IncentiveOffered(
        FPA.Range rangeIncrease,
        FPA.SampleSize sampleSizeIncrease,
        FPA.Fee offerAmount
    );

    function offerIncentive(IncentiveOffer calldata) external payable virtual;
}