// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import { ECPoint, ECPoint2, SortitionCredential, SortitionRound, verifySortitionCredential } from "../lib/Sortition.sol";
import { IVoterRegistry } from "../interface/IVoterRegistry.sol";
import { heapSort } from "../lib/Sort.sol";
import { IIFastUpdaters } from "../interface/IIFastUpdaters.sol";

contract FastUpdaters is IIFastUpdaters {
    struct StagedProviderData {
        bool present;
        ECPoint publicKey;
        uint seedScore;
    }

    mapping (address => StagedProviderData) stagedProviders;
    address[] stagedProviderAddresses;

    uint baseSeed;

    function stagedProviderData(
        ECPoint calldata publicKey, 
        uint score
    ) private pure returns(StagedProviderData memory) {
        return StagedProviderData(true, publicKey, score);
    }

    function registerNewProvider(NewProvider calldata newProvider) external override {
        SortitionRound memory round = SortitionRound(baseSeed, type(uint).max);
        (, uint score) = verifySortitionCredential(round, newProvider.publicKey, 0, newProvider.credential);
        stagedProviders[msg.sender] = stagedProviderData(newProvider.publicKey, score);
        stagedProviderAddresses.push(msg.sender);
    }

    function nextProviderRegistry(uint epochId) public override returns (ProviderRegistry memory registry) { // only governance
        uint totalWeight = voterRegistry.totalWeightPerRewardEpoch(epochId);
        (address[] memory voters, uint[] memory weights) = voterRegistry.votersForRewardEpoch(epochId);

        // Activate staged providers if they are registered voters
        // Here, we just pack them at the beginning of the already-allocated voters and weights arrays
        uint numProviders;
        for (uint i = 0; i < voters.length; ++i) {
            address voter = voters[i];
            StagedProviderData storage voterData = stagedProviders[voter];
            if (voterData.present) {
                voters[numProviders] = voter;
                // Assuming that weights have only up to (256 - VIRTUAL_PROVIDER_BITS) bits (= 244, a safe assumption)
                weights[numProviders] = (weights[i] << VIRTUAL_PROVIDER_BITS) / totalWeight;

                ++numProviders;
            }
        }

        // Allocate just the right amount of space for the return values
        uint[] memory seedScores = new uint[](numProviders);
        registry.providerAddresses = new address[](numProviders);
        registry.providerKeys = new ECPoint[](numProviders);
        registry.providerWeights = new uint[](numProviders);

        // Copy the packed arrays into the return values
        for (uint i = 0; i < numProviders; ++i) {
            address voter = voters[i];
            StagedProviderData storage voterData = stagedProviders[voter];

            registry.providerAddresses[i] = voter;
            registry.providerWeights[i] = weights[i];
            registry.providerKeys[i] = voterData.publicKey;
            seedScores[i] = voterData.seedScore;
        }

        // Recalculate the base seed
        heapSort(seedScores);
        registry.seed = baseSeed = uint(sha256(abi.encodePacked(seedScores)));

        // Finally, clear the staged providers for the next reward epoch
        for (uint i = 0; i < stagedProviderAddresses.length; ++i) {
            address addr = stagedProviderAddresses[i];
            delete stagedProviders[addr];
            delete stagedProviderAddresses[i];
        }
    }
}

// The number of units of weight distributed among providers is 1 << VIRTUAL_PROVIDER_BITS
uint constant VIRTUAL_PROVIDER_BITS = 12;