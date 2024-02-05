// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

// Copied as a stub from ftso-scaling repo 
interface VoterRegistry {
    /**
     * Returns voter's public key and normalised weight for a given reward epoch and signing policy address
     */
    function getPublicKeyAndNormalisedWeight(
        uint256 _rewardEpochId,
        address _signingPolicyAddress
    )
        external view
        returns (
            bytes32 _publicKeyPart1,
            bytes32 _publicKeyPart2,
            uint16 _normalisedWeight
        );    
}
