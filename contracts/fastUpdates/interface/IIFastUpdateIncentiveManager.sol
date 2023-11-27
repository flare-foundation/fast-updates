// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import { IFastUpdateIncentiveManager } from "./IFastUpdateIncentiveManager.sol";

abstract contract IIFastUpdateIncentiveManager is IFastUpdateIncentiveManager {
    address payable internal rewardPool;

    function setRewardPool(address payable newRewardPool) external {
        rewardPool = newRewardPool;
    }

    function nextSortitionParameters() public virtual returns(uint16 newRange8x8, uint16 newPrecision1x15);
}