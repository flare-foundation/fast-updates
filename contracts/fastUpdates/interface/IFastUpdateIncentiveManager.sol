// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import "../lib/FixedPointArithmetic.sol" as FPA;

interface IFastUpdateIncentiveManager {
    struct IncentiveOffer {
        FPA.Range rangeIncrease;
        FPA.Range rangeLimit;
    }

    function getExpectedSampleSize() external view returns (FPA.SampleSize);

    function getPrecision() external view returns (FPA.Precision);

    function getRange() external view returns (FPA.Range);

    function getScale() external view returns (FPA.Scale);

    event IncentiveOffered(FPA.Range rangeIncrease, FPA.SampleSize sampleSizeIncrease, FPA.Fee indexed offerAmount);

    function offerIncentive(IncentiveOffer calldata) external payable;
}
