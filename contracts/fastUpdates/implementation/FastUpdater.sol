// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import { FastUpdaters, VIRTUAL_PROVIDER_BITS } from "./FastUpdaters.sol";
import { FastUpdateIncentiveManager } from "./FastUpdateIncentiveManager.sol";
import { Deltas } from "../lib/Deltas.sol";
import { SortitionRound, SortitionCredential, verifySortitionCredential } from "../lib/Sortition.sol";
import { IIFastUpdater } from "../interface/IIFastUpdater.sol";
import { IIFastUpdaters } from "../interface/IIFastUpdaters.sol";
import { IIFastUpdateIncentiveManager } from "../interface/IIFastUpdateIncentiveManager.sol";
import "../lib/FixedPointArithmetic.sol" as FPA;
import "../lib/Bn256.sol";
import "hardhat/console.sol";

contract FastUpdater is IIFastUpdater {
    // Circular list
    SortitionRound[] public activeSortitionRounds;
    FPA.Price[1000] public prices;
    FPA.Scale public scale;
    uint public submissionWindow;

    function setSubmissionWindow(uint w) public override { // only governance
        delete activeSortitionRounds;
        for (uint i = 0; i < w; ++i) {  
            activeSortitionRounds.push();
        }
        submissionWindow = w;
    }

    constructor(
        IIFastUpdaters _fastUpdaters, 
        IIFastUpdateIncentiveManager _fastUpdateIncentiveManager,
        FPA.Price[] memory _prices,
        uint _submissionWindow,
        uint epochId
    ) IIFastUpdater(_fastUpdaters, _fastUpdateIncentiveManager)
    {
        setPrices(_prices);
        setSubmissionWindow(_submissionWindow);
        finalizeBlock(true, epochId);
    }

    function getSortitionRound(uint blockNum) public view returns (uint seed, uint cutoff) {
        SortitionRound memory sortitionRound = activeSortitionRounds[blockNum % activeSortitionRounds.length];
        seed = sortitionRound.seed;
        cutoff = sortitionRound.scoreCutoff;
    }

    function setPrices(FPA.Price[] memory _prices) public { // only governance
        for (uint i = 0; i < _prices.length; ++i) {
            prices[i] = _prices[i];
        }
    }

    function setNextSortitionRound(bool newSeed, uint epochId, FPA.SampleSize newSampleSize) private {
        uint cutoff = getScoreCutoff(newSampleSize);
        uint seed;
        if (newSeed) { // TODO: this needs to be replaced with a real condition
            for (uint i = 0; i < activeProviderAddresses.length; ++i) {
                delete activeProviders[activeProviderAddresses[i]];
            }
            delete activeProviderAddresses;

            IIFastUpdaters.ProviderRegistry memory registry = fastUpdaters.nextProviderRegistry(epochId);

            for (uint i = 0; i < registry.providerAddresses.length; ++i) {
                address addr = registry.providerAddresses[i];
                activeProviders[addr] = ActiveProviderData(registry.providerKeys[i], registry.providerWeights[i]);
                activeProviderAddresses.push(addr);
            }
            seed = registry.seed;
        }
        else {
            seed = getPreviousSortitionRound(0).seed + 1;
        }

        setNextSortitionRound(SortitionRound(seed, cutoff));
    }

    // Called by Flare daemon at the end of each block
    function finalizeBlock(bool newSeed, uint epochId) public override { // only governance
        FPA.SampleSize newSampleSize;
        (newSampleSize, scale) = fastUpdateIncentiveManager.nextUpdateParameters();
        setNextSortitionRound(newSeed, epochId, newSampleSize);
    }

    function submitUpdates(FastUpdates calldata updates) external override {
        uint blocksAgo = block.number - updates.sortitionBlock;
        SortitionRound storage sortitionRound = getPreviousSortitionRound(blocksAgo);
        ActiveProviderData storage providerData = activeProviders[msg.sender];

        verifySortitionCredential(sortitionRound, providerData.publicKey, providerData.sortitionWeight, updates.sortitionCredential);
        applyUpdates(updates.deltas);
        uint epochId; // TODO: use a real value here
        emit FastUpdate(epochId, msg.sender);
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

    function fetchCurrentPrices(
        uint[] calldata feeds
    ) external view override returns(FPA.Price[] memory _prices) {
        _prices = new FPA.Price[](feeds.length);
        for (uint i = 0; i < feeds.length; ++i) {
            _prices[i] = prices[feeds[i]];
        }
    }

    function applyUpdates(Deltas calldata deltas) private {
        deltas.forEach(applyDelta); // TODO: optimize these calls for storage access
    }

    function applyDelta(
        int delta,
        uint feed
    ) private {
        if (delta == -1) {
            prices[feed] = FPA.div(prices[feed], scale);
        }
        else if (delta == 1) {
            prices[feed] = FPA.mul(prices[feed], scale);
        }
    }
}
