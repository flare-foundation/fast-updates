// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import {SortitionCredential, SortitionRound, verifySortitionCredential} from "../lib/Sortition.sol";
import {IVoterRegistry} from "./IVoterRegistry.sol";
import "../lib/Bn256.sol";

contract FastUpdaters {
    struct ActiveProviderData {
        Bn256.G1Point publicKey;
        uint sortitionWeight;
    }

    struct StagedProviderData {
        bool present;
        Bn256.G1Point publicKey;
        uint seedScore;
    }

    mapping(address => ActiveProviderData) public activeProviders;
    address[] activeProviderAddresses; // Must have uint8 length

    function numProviders() public view returns (uint) {
        return activeProviderAddresses.length;
    }

    mapping(address => StagedProviderData) stagedProviders;
    address[] stagedProviderAddresses;

    Bn256.G1Point private ecBasePoint;
    uint public baseSeed;

    IVoterRegistry voterRegistry;

    function setVoterRegistry(IVoterRegistry registry) public {
        // only governance
        voterRegistry = registry;
    }

    // This is because Solidity's autogen'd getters intentionally screw up returned structs
    function getECBasePoint() public view returns (Bn256.G1Point memory) {
        return ecBasePoint;
    }

    function sortitionPublicKey(address provider) public view returns (Bn256.G1Point memory) {
        return activeProviders[provider].publicKey;
    }

    function stagedProviderData(
        Bn256.G1Point calldata publicKey,
        uint score
    ) private pure returns (StagedProviderData memory) {
        return StagedProviderData(true, publicKey, score);
    }

    function registerNewProvider(Bn256.G1Point calldata publicKey, SortitionCredential calldata credential) public {
        SortitionRound memory round = SortitionRound(baseSeed, type(uint).max);
        (, uint score) = verifySortitionCredential(round, publicKey, 0, credential);
        stagedProviders[msg.sender] = stagedProviderData(publicKey, score);
        stagedProviderAddresses.push(msg.sender);
    }

    function finalizeRewardEpoch(uint epochId) public {
        // only governance
        // Clear active providers
        for (uint i = 0; i < activeProviderAddresses.length; ++i) {
            address provider = activeProviderAddresses[i];
            delete activeProviders[provider];
        }
        delete activeProviderAddresses;

        // Activate staged providers if they are registered voters
        uint totalWeight = voterRegistry.totalWeightPerRewardEpoch(epochId);
        (address[] memory voters, uint[] memory weights) = voterRegistry.votersForRewardEpoch(epochId);
        uint[] memory seedScores = new uint[](voters.length);
        for (uint i = 0; i < voters.length; ++i) {
            address voter = voters[i];
            StagedProviderData storage voterData = stagedProviders[voter];
            if (voterData.present) {
                // Assuming that weights have only up to (256 - VIRTUAL_PROVIDER_BITS) bits (= 244, a safe assumption)
                uint sortitionWeight = (weights[i] << VIRTUAL_PROVIDER_BITS) / totalWeight;

                activeProviders[voter] = ActiveProviderData(voterData.publicKey, sortitionWeight);
                activeProviderAddresses.push(voter);
                seedScores[activeProviderAddresses.length] = voterData.seedScore;
            }
        }

        // Recalculate the base seed
        // The end-padding of the list of scores with 0s is unfortunate but harmless
        baseSeed = uint(sha256(abi.encodePacked(seedScores)));
    }
}

// The number of units of weight distributed among providers is 1 << VIRTUAL_PROVIDER_BITS
uint constant VIRTUAL_PROVIDER_BITS = 12;
