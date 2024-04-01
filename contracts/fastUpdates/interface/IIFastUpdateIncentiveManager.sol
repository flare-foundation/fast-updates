// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import {IFastUpdateIncentiveManager} from "./IFastUpdateIncentiveManager.sol";
import {Governed} from "./Governed.sol";
import "../lib/FixedPointArithmetic.sol" as FPA;

abstract contract IIFastUpdateIncentiveManager is IFastUpdateIncentiveManager, Governed {
    address payable internal rewardPool;

    FPA.SampleSize internal baseSampleSize;
    FPA.Range internal baseRange;

    FPA.SampleSize internal sampleIncreaseLimit;
    FPA.Fee internal rangeIncreasePrice;

    function setRewardPool(address payable _rp) external onlyGovernance {
        _setRewardPool(_rp);
    }

    function _setRewardPool(address payable _rp) private {
        rewardPool = _rp;
    }

    function setBaseSampleSize(FPA.SampleSize _sz) external onlyGovernance {
        _setBaseSampleSize(_sz);
    }

    function _setBaseSampleSize(FPA.SampleSize _sz) private {
        require(FPA.check(_sz), "Base sample size too large");
        baseSampleSize = _sz;
    }

    function setBaseRange(FPA.Range _rn) external onlyGovernance {
        _setBaseRange(_rn);
    }

    function _setBaseRange(FPA.Range _rn) private {
        require(FPA.check(_rn), "Base range too large");
        baseRange = _rn;
    }

    function setSampleIncreaseLimit(FPA.SampleSize _lim) external onlyGovernance {
        _setSampleIncreaseLimit(_lim);
    }

    function _setSampleIncreaseLimit(FPA.SampleSize _lim) private {
        require(FPA.check(_lim), "Sample increase limit too large");
        sampleIncreaseLimit = _lim;
    }

    function setRangeIncreasePrice(FPA.Fee _price) external onlyGovernance {
        _setRangeIncreasePrice(_price);
    }

    function _setRangeIncreasePrice(FPA.Fee _price) private {
        require(FPA.check(_price), "Range increase price too large");
        rangeIncreasePrice = _price;
    }

    constructor(
        address _governance,
        address payable _rp,
        FPA.SampleSize _bss,
        FPA.Range _br,
        FPA.SampleSize _sil,
        FPA.Fee _rip
    ) Governed(_governance) {
        _setRewardPool(_rp);
        _setBaseSampleSize(_bss);
        _setBaseRange(_br);
        _setSampleIncreaseLimit(_sil);
        _setRangeIncreasePrice(_rip);
    }

    function advance() external virtual;
    function setIncentiveDuration(uint _duration) external virtual;
}
