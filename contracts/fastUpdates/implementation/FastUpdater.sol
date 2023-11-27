// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import { FastUpdaters, VIRTUAL_PROVIDER_BITS } from "./FastUpdaters.sol";
import { FastUpdateIncentiveManager } from "./FastUpdateIncentiveManager.sol";
import { Deltas } from "../lib/Deltas.sol";
import { ECPoint, ECPoint2, SortitionRound, SortitionCredential, verifySortitionCredential } from "../lib/Sortition.sol";

contract FastUpdater {
    FastUpdaters private fastUpdaters;
    FastUpdateIncentiveManager private fastUpdateIncentiveManager;

    struct ActiveProviderData {
        ECPoint publicKey;
        uint sortitionWeight;
    }

    SortitionRound[] private activeSortitionRounds;
    mapping (address => ActiveProviderData) public activeProviders;
    address[] activeProviderAddresses; // Must have uint8 length

    function setFastUpdaters(FastUpdaters addr) public { // onlyGovernance
        fastUpdaters = addr;
    }

    function setSortitionParameters() private returns(uint16 expectedSampleSize8x8) {
        uint16 newPrecision1x15;
        (expectedSampleSize8x8, newPrecision1x15) = fastUpdateIncentiveManager.nextSortitionParameters();
        setPrecision(newPrecision1x15);
    }

    function setNextSortitionRound(bool newSeed, uint16 expectedSampleSize8x8) private {
        uint epochId; // TODO: Get this correctly
        uint cutoff = getScoreCutoff(expectedSampleSize8x8);
        uint seed;
        if (newSeed) { // TODO: this needs to be replaced with a real condition
            for (uint i = 0; i < activeProviderAddresses.length; ++i) {
                delete activeSortitionRounds[i];
            }
            ECPoint[] memory providerKeys;
            uint[] memory providerWeights;
            (seed, activeProviderAddresses, providerKeys, providerWeights) = fastUpdaters.nextProviderData(epochId);
            for (uint i = 0; i < activeProviderAddresses.length; ++i) {
                activeProviders[activeProviderAddresses[i]] = ActiveProviderData(providerKeys[i], providerWeights[i]);
            }
        }
        else {
            seed = getPreviousSortitionRound(0).seed + 1;
        }
        setNextSortitionRound(SortitionRound(seed, cutoff));
    }

    // Called by Flare daemon at the end of each block
    function finalizeBlock(bool newSeed) public { // only governance
        uint16 expectedSampleSize8x8 = setSortitionParameters();
        setNextSortitionRound(newSeed, expectedSampleSize8x8);
    }

    function submitUpdates(
        uint64 sortitionBlock,
        SortitionCredential calldata sortitionCredential,
        Deltas calldata deltas
    ) public {
        uint blocksAgo = block.number - sortitionBlock;
        SortitionRound storage sortitionRound = getPreviousSortitionRound(blocksAgo);
        ActiveProviderData storage providerData = activeProviders[msg.sender];

        verifySortitionCredential(sortitionRound, providerData.publicKey, providerData.sortitionWeight, sortitionCredential);
        applyUpdates(deltas);
    }

    function getScoreCutoff(uint16 expectedSampleSize8x8) private pure returns (uint) {
        // The formula is: (exp. s.size)/(num. prov.) = (score)/(score range)
        //   score range = 2**256
        //   num. prov.  = 2**VIRTUAL_PROVIDER_BITS
        //   exp. s.size = "expectedSampleSize8x8 >> 8", in that we keep the fractional bits:
        return uint(expectedSampleSize8x8) << (256 - VIRTUAL_PROVIDER_BITS - 8);
    }

    function ix(uint i) private view returns (uint) {
        return (i + block.number) % activeSortitionRounds.length;
    }
    function getPreviousSortitionRound(uint i) private view returns (SortitionRound storage) {
        assert(i < activeSortitionRounds.length);
        return activeSortitionRounds[ix(activeSortitionRounds.length - i)];
    }
    function setNextSortitionRound(SortitionRound memory x) private {
        activeSortitionRounds[ix(1)] = x;
    }
    function setSubmissionWindow(uint w) private { // only governance
        delete activeSortitionRounds;
        for (uint i = 0; i < w; ++i) {
            activeSortitionRounds.push();
        }
    }

    uint32[1000] private anchorPrices;
    int8[1000] private totalUnitDeltas;

    // stand-in for uint16[8]; precisionPowers[i] = 1 + (0x15 bit fraction) = precision1x15 ** (2 ** i)
    bytes16 private precisionPowers; // Must call setPrecision before this is used!

    function padRight16(uint16 x) private pure returns(bytes16) {
        return bytes16(bytes2(x));
    }

    function mulFixed1x15(uint16 x, uint16 y) private pure returns(uint16) {
        return uint16((uint32(x) * uint32(y)) >> 15);
    }

    function setPrecision(uint16 scale1x15) private {
        // Unavoidably expensive: when the precision changes the meaning of totalUnitDeltas also changes
        for (uint feed = 0; feed < 1000; ++feed) {
            weighAnchorPrice(feed);
        }

        bytes16 powers = padRight16(scale1x15);
        for (uint i = 1; i < 7; ++i) {
            scale1x15 = mulFixed1x15(scale1x15, scale1x15);
            powers |= padRight16(scale1x15) >> (16 * i);
        }
        precisionPowers = powers;
    }

    function deltaFactor(int8 totalUnitDelta) private view returns (uint16 factor1x15) {
        bytes16 powers = precisionPowers;
        bytes1 deltaBinary = bytes1(uint8(totalUnitDelta));
        factor1x15 = uint16(bytes2(hex"a0_00"));

        bytes1 deltaBitMask = hex"01";
        bytes16 powerMask = hex"ff_00_00_00_00_00_00_00";

        while (deltaBinary != bytes1(0)) {
            if (deltaBinary & deltaBitMask != 0) {
                uint16 power1x15 = uint16(bytes2(precisionPowers & powerMask));
                factor1x15 = mulFixed1x15(factor1x15, power1x15);
            }

            powers <<= 16;
            deltaBinary >>= 1;
        }
    }

    function fetchCurrentPrices(
        uint32[] calldata feeds
    ) public view returns(uint[] memory prices) {
        prices = new uint[](feeds.length);
        for (uint i = 0; i < feeds.length; ++i) {
            uint32 feed = feeds[i];
            prices[i] = computePrice(anchorPrices[feed], totalUnitDeltas[feed]);
        }
    }

    function computePrice(uint32 anchorPrice, int8 totalUnitDelta) private view returns(uint32) {
        return uint32(uint(anchorPrice) * uint(deltaFactor(totalUnitDelta)) >> 15);
    }

    function applyUpdates(
        Deltas calldata deltas
    ) private {
        deltas.forEach(applyDelta); // TODO: optimize these calls for storage access
    }

    function applyDelta(
        int delta,
        uint feed
    ) private {
        int8 totalUnitDelta = totalUnitDeltas[feed];

        if (totalUnitDelta == type(int8).min || totalUnitDelta == type(int8).max) {
            weighAnchorPrice(feed, int8(delta));
        }
        else {
            totalUnitDeltas[feed] = int8(totalUnitDelta + delta);
        }
    }

    function weighAnchorPrice(uint feed, int8 extraDelta) private {
        anchorPrices[feed] = computePrice(anchorPrices[feed], totalUnitDeltas[feed]);
        totalUnitDeltas[feed] = extraDelta;
    }

    function weighAnchorPrice(uint feed) private {
        weighAnchorPrice(feed, 0);
    }
}