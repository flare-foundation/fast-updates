// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

abstract contract IFastUpdateIncentiveManager {
    uint16 public expectedSampleSize8x8;
    uint16 public precision0x16;
    uint16 public range8x8;

    struct IncentiveOffer {
        uint16 variationRangeIncrease0x16;
        uint16 rangeLimit8x8;
    }

    function offerIncentive(IncentiveOffer calldata) external payable virtual;
}