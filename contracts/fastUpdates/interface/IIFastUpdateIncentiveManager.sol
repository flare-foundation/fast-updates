// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import { IFastUpdateIncentiveManager } from "./IFastUpdateIncentiveManager.sol";
import "../lib/FixedPointArithmetic.sol" as FPA;

abstract contract IIFastUpdateIncentiveManager is IFastUpdateIncentiveManager {
    address payable internal rewardPool;

    FPA.SampleSize internal baseSampleSize;
    FPA.Range internal baseRange;

    FPA.SampleSize internal sampleIncreaseLimit;
    FPA.Fee internal rangeIncreasePrice;

    function setRewardPool(address payable _rp) external { // only governance
        rewardPool = _rp;
    }

    function setBaseSampleSize(FPA.SampleSize _sz) external { // only governance
        baseSampleSize = _sz;
    }

    function setBaseRange(FPA.Range _rn) external { // only governance
        baseRange = _rn;
    }

    function setSampleIncreaseLimit(FPA.SampleSize _lim) external { // only governance
        sampleIncreaseLimit = _lim;
    }

    function setRangeIncreasePrice(FPA.Fee _price) external { // only governance
        rangeIncreasePrice = _price;
    }

    function nextUpdateParameters() public virtual returns(FPA.SampleSize newSampleSize, FPA.Scale newScale);
}