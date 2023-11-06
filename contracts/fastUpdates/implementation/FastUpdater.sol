// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import { FastUpdaters } from "./FastUpdaters.sol";
import { FastUpdateManager } from "./FastUpdateManager.sol";
import { Deltas } from "../lib/Deltas.sol";
import { ECPoint, ECPoint2, SortitionRound, SortitionCredential, verifySortitionCredential } from "../lib/Sortition.sol";

contract FastUpdater {
    FastUpdaters private fastUpdaters;
    FastUpdateManager private fastUpdateManager;

    uint32[1000] private anchorPrices;
    int8[1000] private totalUnitDeltas; // maintained so that the true price is never negative, nor overflows

    // A circular buffer
    SortitionRound[] private activeSortitionRounds;
    function ix(uint i) private view returns (uint) {
        return (i + block.number) % activeSortitionRounds.length;
    }
    function getSortitionRound(uint i) private view returns (SortitionRound storage) {
        assert(i < activeSortitionRounds.length);
        return activeSortitionRounds[ix(i)];
    }
    function setCurrentSortitionRound(SortitionRound memory x) private {
        activeSortitionRounds[ix(0)] = x;
    }
    function setSubmissionWindow(uint w) private {
        delete activeSortitionRounds;
        for (uint i = 0; i < w; ++i) {
            activeSortitionRounds.push();
        }
    }

    // Called by Flare daemon at the end of each block
    function finalizeBlock(bool newSeed) public {
        uint numParticipants = fastUpdaters.numParticipants();
        uint cutoff = fastUpdateManager.getScoreCutoff(numParticipants);
        uint seed = newSeed ? fastUpdateManager.baseSeed() : getSortitionRound(0).seed + 1;
        setCurrentSortitionRound(SortitionRound(seed, cutoff));
    }

    function setFastUpdaters(FastUpdaters addr) public { // onlyGovernance
        fastUpdaters = addr;
    }

    function setFastUpdateManager(FastUpdateManager addr) public { // onlyGovernance
        fastUpdateManager = addr;
    }

    function submitUpdates(
        uint64 sortitionBlock,
        SortitionCredential calldata sortitionCredential,
        Deltas calldata deltas
    ) public {
        uint blocksAgo = block.number - sortitionBlock;
        SortitionRound storage sortitionRound = getSortitionRound(blocksAgo);
        ECPoint memory publicKey = fastUpdaters.sortitionPublicKey(msg.sender);
        ECPoint2 memory basePoint = fastUpdateManager.getECBasePoint();

        verifySortitionCredential(sortitionRound, publicKey, basePoint, sortitionCredential);
        applyUpdates(deltas);
    }

    function fetchCurrentPrices(
        uint[] calldata feeds
    ) public view returns(uint[] memory prices) {
        uint[] memory feedDeltas = fastUpdateManager.getNumericDeltas(feeds);
        prices = new uint[](feeds.length);
        for (uint i = 0; i < feeds.length; ++i) {
            prices[i] = currentPrice(feedDeltas[i], feeds[i]);
        }
    }

    function currentPrice(
        uint feedDelta,
        uint feed
    ) private view returns (uint) {
        // assumption: currentPrice() >= 0
        return uint(currentPrice(anchorPrices[feed], feedDelta, totalUnitDeltas[feed]));
    }

    function currentPrice(
        uint anchorPrice,
        uint feedDelta, // assumption: int(feedDelta) >= 0
        int totalUnitDelta
    ) private pure returns(int) {
        int ap = int(uint(anchorPrice));
        int pd = int(feedDelta) * totalUnitDelta;
        return ap + pd;
    }

    function applyUpdates(
        Deltas calldata deltas
    ) private {
        deltas.forEach(applyDelta); // TODO: optimize these calls for storage access
    }

    function applyDelta(
        int delta,
        uint feed
    ) private  {
        int newTotalUnitDelta = totalUnitDeltas[feed] + delta;
        if (newTotalUnitDelta < type(int8).min || newTotalUnitDelta > type(int8).max) {
            uint[] memory feed1 = new uint[](1);
            feed1[0] = feed;
            // i.e. getNumericDeltas([feed])
            uint[] memory feedDeltas = fastUpdateManager.getNumericDeltas(feed1); // TODO: optimize these calls
            int newAnchorPrice = currentPrice(anchorPrices[feed], feedDeltas[0], newTotalUnitDelta);
            anchorPrices[feed] = newAnchorPrice < 0 ? 0 : uint32(uint(newAnchorPrice));
            totalUnitDeltas[feed] = 0;
        }
        else {
            totalUnitDeltas[feed] = int8(newTotalUnitDelta);
        }
    }
}