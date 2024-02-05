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

    function setRewardPool(address payable _rp) public { // only governance
        rewardPool = _rp;
    }

    function setBaseSampleSize(FPA.SampleSize _sz) public { // only governance
        baseSampleSize = _sz;
    }

    function setBaseRange(FPA.Range _rn) public { // only governance
        baseRange = _rn;
    }

    function setSampleIncreaseLimit(FPA.SampleSize _lim) public { // only governance
        sampleIncreaseLimit = _lim;
    }

    function setRangeIncreasePrice(FPA.Fee _price) public { // only governance
        rangeIncreasePrice = _price;
    }

    constructor(address payable _rp, FPA.SampleSize _bss, FPA.Range _br, FPA.SampleSize _sil, FPA.Fee _rip) {
        setRewardPool(_rp);
        setBaseSampleSize(_bss);
        setBaseRange(_br);
        setSampleIncreaseLimit(_sil);
        setRangeIncreasePrice(_rip);
    }

    function advance() public virtual;
}