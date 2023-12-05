// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import { IFastUpdateIncentiveManager } from "./IFastUpdateIncentiveManager.sol";
import "../lib/FixedPointArithmetic.sol" as FPA;

abstract contract IIFastUpdateIncentiveManager is IFastUpdateIncentiveManager {
    address payable internal rewardPool;

    function setRewardPool(address payable newRewardPool) external { // only governance
        rewardPool = newRewardPool;
    }

    function nextUpdateParameters() public virtual returns(FPA.SampleSize newSampleSize, FPA.Scale newScale);
}