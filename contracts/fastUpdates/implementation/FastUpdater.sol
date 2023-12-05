// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import { FastUpdaters, VIRTUAL_PROVIDER_BITS } from "./FastUpdaters.sol";
import { FastUpdateIncentiveManager } from "./FastUpdateIncentiveManager.sol";
import { Deltas } from "../lib/Deltas.sol";
import { SortitionRound, SortitionCredential, verifySortitionCredential } from "../lib/Sortition.sol";
import { IIFastUpdater } from "../interface/IIFastUpdater.sol";
import { IIFastUpdaters } from "../interface/IIFastUpdaters.sol";
import "../lib/FixedPointArithmetic.sol" as FPA;
import "../lib/Bn256.sol";

// TODO: governance functions to change the anchor prices, e.g. at the time of deployment of this contract
contract FastUpdater is IIFastUpdater {
    // Circular list
    SortitionRound[] private activeSortitionRounds;
    // An FTSO v2 price is 32-bit, as per the ftso-scaling repo
    FPA.Price[1000] private anchorPrices;
    // We accumulate -128..127 unit deltas before recomputing the anchor price
    FPA.Delta[1000] private totalUnitDeltas;
    // scalePowers[i] = 1 + (0x15 bit precision) = scale ** (2 ** i)
    FPA.Scale[8] private scalePowers; // Must call setScale before this is used!

    function setSortitionParameters() private returns(FPA.SampleSize newSampleSize) {
        FPA.Scale newScale;
        (newSampleSize, newScale) = fastUpdateIncentiveManager.nextUpdateParameters();
        setScale(newScale);
    }

    function setNextSortitionRound(bool newSeed, FPA.SampleSize newSampleSize) private {
        uint epochId; // TODO: Get this correctly
        uint cutoff = getScoreCutoff(newSampleSize);
        uint seed;
        if (newSeed) { // TODO: this needs to be replaced with a real condition
            for (uint i = 0; i < activeProviderAddresses.length; ++i) {
                delete activeSortitionRounds[i];
            }
            IIFastUpdaters.ProviderRegistry memory registry = fastUpdaters.nextProviderRegistry(epochId);
            for (uint i = 0; i < registry.providerAddresses.length; ++i) {
                activeProviders[registry.providerAddresses[i]] = ActiveProviderData(registry.providerKeys[i], registry.providerWeights[i]);
            }
            seed = registry.seed;
        }
        else {
            seed = getPreviousSortitionRound(0).seed + 1;
        }
        setNextSortitionRound(SortitionRound(seed, cutoff));
    }

    // Called by Flare daemon at the end of each block
    function finalizeBlock() public override { // only governance
        bool newSeed; // TODO: use a real thing here
        FPA.SampleSize newSampleSize = setSortitionParameters();
        setNextSortitionRound(newSeed, newSampleSize);
    }

    function submitUpdates(FastUpdates calldata updates) external override {
        uint blocksAgo = block.number - updates.sortitionBlock;
        SortitionRound storage sortitionRound = getPreviousSortitionRound(blocksAgo);
        ActiveProviderData storage providerData = activeProviders[msg.sender];

        verifySortitionCredential(sortitionRound, providerData.publicKey, providerData.sortitionWeight, updates.sortitionCredential);
        applyUpdates(updates.deltas);
    }

    function getScoreCutoff(FPA.SampleSize expectedSampleSize) private pure returns (uint) {
        // The formula is: (exp. s.size)/(num. prov.) = (score)/(score range)
        //   score range = 2**256
        //   num. prov.  = 2**VIRTUAL_PROVIDER_BITS
        //   exp. s.size = "expectedSampleSize8x8 >> 8", in that we keep the fractional bits:
        return uint(FPA.SampleSize.unwrap(expectedSampleSize)) << (256 - VIRTUAL_PROVIDER_BITS - 8);
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
    function setSubmissionWindow(uint w) external override { // only governance
        delete activeSortitionRounds;
        for (uint i = 0; i < w; ++i) {  
            activeSortitionRounds.push();
        }
    }

    function setScale(FPA.Scale x) private {
        // Unavoidably expensive: when the precision changes the meaning of totalUnitDeltas also changes
        for (uint feed = 0; feed < 1000; ++feed) {
            weighAnchorPrice(feed);
        }

        FPA.powersInto(x, scalePowers);
    }

    function fetchCurrentPrices(
        uint[] calldata feeds
    ) external view override returns(uint[] memory prices) {
        prices = new uint[](feeds.length);
        for (uint i = 0; i < feeds.length; ++i) {
            uint feed = feeds[i];
            prices[i] = FPA.Price.unwrap(computePrice(anchorPrices[feed], totalUnitDeltas[feed]));
        }
    }

    function computePrice(FPA.Price anchorPrice, FPA.Delta totalUnitDelta) private view returns(FPA.Price) {
        return FPA.mul(anchorPrice, FPA.pow(scalePowers, totalUnitDelta));
    }

    function applyUpdates(Deltas calldata deltas) private {
        deltas.forEach(applyDelta); // TODO: optimize these calls for storage access
    }

    function applyDelta(
        FPA.Delta delta,
        uint feed
    ) private {
        FPA.Delta totalUnitDelta = totalUnitDeltas[feed];

        if (FPA.minDelta(totalUnitDelta) || FPA.maxDelta(totalUnitDelta)) {
            weighAnchorPrice(feed, delta);
        }
        else {
            totalUnitDeltas[feed] = FPA.add(totalUnitDelta, delta);
        }
    }

    function weighAnchorPrice(uint feed, FPA.Delta extraDelta) private {
        anchorPrices[feed] = computePrice(anchorPrices[feed], totalUnitDeltas[feed]);
        totalUnitDeltas[feed] = extraDelta;
    }

    function weighAnchorPrice(uint feed) private {
        weighAnchorPrice(feed, FPA.zeroD);
    }
}
