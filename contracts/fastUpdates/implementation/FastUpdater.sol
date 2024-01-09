// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;
pragma abicoder v2;

import { FastUpdaters, VIRTUAL_PROVIDER_BITS } from "./FastUpdaters.sol";
import { FastUpdateIncentiveManager } from "./FastUpdateIncentiveManager.sol";
import { Deltas } from "../lib/Deltas.sol";
import { SortitionRound, sortitionRound, SortitionCredential, verifySortitionCredential } from "../lib/Sortition.sol";
import { IIFastUpdater } from "../interface/IIFastUpdater.sol";
import { IIFastUpdaters } from "../interface/IIFastUpdaters.sol";
import { IIFastUpdateIncentiveManager } from "../interface/IIFastUpdateIncentiveManager.sol";
import { IICircular } from "../interface/IICircular.sol";
import "../lib/FixedPointArithmetic.sol" as FPA;
import "../lib/Bn256.sol";
import "hardhat/console.sol";

abstract contract RoundManager is IICircular {
    // Circular list
    SortitionRound[] public activeSortitionRounds;

    constructor(uint _dur)
        IICircular(_dur)
    {
        init();
    }

    function init() internal {
        delete activeSortitionRounds;
        for (uint i = 0; i < circularLength; ++i) {  
            activeSortitionRounds.push();
        }
    }

    function step(uint cutoff) internal {
        step(activeSortitionRounds[thisIx()].seed + 1, cutoff);
    }

    function step(uint seed, uint cutoff) internal {
        activeSortitionRounds[nextIx()] = sortitionRound(seed, cutoff);
    }

    // A round could be missing either because it was too many blocks ago and was dropped,
    // or because it was in a previous reward epoch, and was dropped.
    function getRound(uint blockNum) view internal returns (SortitionRound storage round) {
        string memory failMsg = "Sortition round for the given block is no longer available";
        uint ix = blockIx(blockNum, failMsg);
        round = activeSortitionRounds[ix];
        require(round.present, failMsg);
    }
}

contract FastUpdater is IIFastUpdater, RoundManager {
    FPA.Price[1000] public prices;
    FPA.Scale public scale;

    constructor(
        IIFastUpdaters _fastUpdaters, 
        IIFastUpdateIncentiveManager _fastUpdateIncentiveManager,
        FPA.Price[] memory _prices,
        uint _submissionWindow,
        uint epochId
    ) 
        IIFastUpdater(_fastUpdaters, _fastUpdateIncentiveManager)
        RoundManager(_submissionWindow)
    {
        setPrices(_prices);
        setSubmissionWindow(_submissionWindow);
        finalizeBlock(true, epochId);
    }
    
    function getSubmissionWindow() public view override returns(uint) {
        return circularLength;
    }

    function setSubmissionWindow(uint w) public override { // only governance
        setCircularLength(w);
        RoundManager.init();
    }

    function getSortitionRound(uint blockNum) external view returns (SortitionRound memory) {
        return getRound(blockNum);
    }

    function setPrices(FPA.Price[] memory _prices) public { // only governance
        for (uint i = 0; i < _prices.length; ++i) {
            prices[i] = _prices[i];
        }
    }

    function setNextSortitionRound(bool newSeed, uint epochId, FPA.SampleSize newSampleSize) private {
        uint cutoff = getScoreCutoff(newSampleSize);
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
            
            // We can't allow the submission window to cross a change of providers, since we don't store both the
            // old and the new provider registries.
            RoundManager.init();
            RoundManager.step(registry.seed, cutoff);
        }
        else {
            RoundManager.step(cutoff);
        }
    }

    // Called by Flare daemon at the end of each block
    function finalizeBlock(bool newSeed, uint epochId) public override { // only governance
        FPA.SampleSize newSampleSize;
        (newSampleSize, scale) = fastUpdateIncentiveManager.nextUpdateParameters();
        setNextSortitionRound(newSeed, epochId, newSampleSize);
    }

    function submitUpdates(FastUpdates calldata updates) external override {
        SortitionRound storage round = getRound(updates.sortitionBlock);
        ActiveProviderData storage providerData = activeProviders[msg.sender];

        verifySortitionCredential(round, providerData.publicKey, providerData.sortitionWeight, updates.sortitionCredential);
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
