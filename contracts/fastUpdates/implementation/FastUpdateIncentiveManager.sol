// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import {IncreaseManager} from "./abstract/IncreaseManager.sol";
import "../lib/FixedPointArithmetic.sol" as FPA;
import {IIFastUpdateIncentiveManager} from "../interface/IIFastUpdateIncentiveManager.sol";

contract FastUpdateIncentiveManager is IIFastUpdateIncentiveManager, IncreaseManager {
    constructor(
        address _governance,
        address payable _rp,
        FPA.SampleSize _bss,
        FPA.Range _br,
        FPA.SampleSize _sil,
        FPA.Fee _rip,
        uint _dur
    ) IIFastUpdateIncentiveManager(_governance, _rp, _bss, _br, _sil, _rip) IncreaseManager(_dur) {
        sampleSize = baseSampleSize;
        range = baseRange;
        excessOfferValue = FPA.Fee.wrap(1); // Arbitrary initial value, but must not be 0
    }

    function getExpectedSampleSize() external view returns (FPA.SampleSize) {
        return sampleSize;
    }

    function getRange() external view returns (FPA.Range) {
        return range;
    }

    function getPrecision() external view returns (FPA.Precision) {
        return computePrecision();
    }

    function computePrecision() private view returns (FPA.Precision) {
        return FPA.div(range, sampleSize); // todo: if range is bigger than sample size, this is problematic
    }

    function getScale() public view returns (FPA.Scale) {
        return FPA.scaleWithPrecision(computePrecision());
    }

    // This is expected to be called once per block by the Flare daemon
    function advance() external override onlyGovernance {
        IncreaseManager.step();
    }

    function offerIncentive(IncentiveOffer calldata offer) external payable {
        (FPA.Fee dc, FPA.Range dr) = processIncentiveOffer(offer);
        FPA.SampleSize de = sampleSizeIncrease(dc, dr);

        IncreaseManager.increaseSampleSize(de);
        IncreaseManager.increaseRange(dr);

        rewardPool.transfer(FPA.Fee.unwrap(dc));
        emit IncentiveOffered(dr, de, dc);
        payable(msg.sender).transfer(msg.value - FPA.Fee.unwrap(dc));
    }

    function processIncentiveOffer(
        IncentiveOffer calldata offer
    ) private returns (FPA.Fee contribution, FPA.Range rangeIncrease) {
        require(msg.value >> 240 == 0, "Incentive offer value capped at 240 bits");
        contribution = FPA.Fee.wrap(uint240(msg.value));
        rangeIncrease = offer.rangeIncrease;

        FPA.Range finalRange = FPA.add(range, rangeIncrease);
        if (FPA.lessThan(offer.rangeLimit, finalRange)) {
            finalRange = offer.rangeLimit;
            FPA.Range newRangeIncrease = FPA.lessThan(finalRange, range) ? FPA.zeroR : FPA.sub(finalRange, range);
            contribution = FPA.mul(FPA.frac(newRangeIncrease, rangeIncrease), contribution);
            rangeIncrease = newRangeIncrease;
        }
        require(FPA.lessThan(finalRange, sampleSize), "Offer would make the precision greater than 100%");
    }

    function setIncentiveDuration(uint _duration) external override onlyGovernance {
        _setIncentiveDuration(_duration);
    }

    function sampleSizeIncrease(FPA.Fee dc, FPA.Range dr) private returns (FPA.SampleSize de) {
        FPA.Fee rangeCost = FPA.mul(rangeIncreasePrice, dr);
        require(!FPA.lessThan(dc, rangeCost), "Insufficient contribution to pay for range increase");
        FPA.Fee dx = FPA.sub(dc, rangeCost);

        IncreaseManager.increaseExcessOfferValue(dx);

        de = FPA.mul(FPA.frac(dx, excessOfferValue), sampleIncreaseLimit);
    }
}
