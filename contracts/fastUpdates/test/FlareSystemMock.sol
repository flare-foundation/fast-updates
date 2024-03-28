// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import "../interface/mocks/VoterRegistry.sol";
import "../interface/mocks/FlareSystemManager.sol";

contract FlareSystemMock is VoterRegistry, FlareSystemManager {
    uint randomSeed;
    uint epochLen;

    struct Policy {
        bytes32 pk_1;
        bytes32 pk_2;
        uint weight;
    }

    mapping(uint => mapping(address => Policy)) policies;
    mapping(uint => uint) totalWeights;

    constructor(uint _randomSeed, uint _epochLen) {
        randomSeed = _randomSeed;
        epochLen = _epochLen;
    }

    function getCurrentRandom() external view returns (uint256 _currentRandom) {
        return uint(sha256(abi.encodePacked(block.number / epochLen, randomSeed)));
    }

    function getCurrentRewardEpochId() external view returns (uint24 _currentRewardEpochId) {
        return uint24(block.number / epochLen);
    }

    function getPublicKeyAndNormalisedWeight(
        uint256 _rewardEpochId,
        address _signingPolicyAddress
    ) external view returns (bytes32 _publicKeyPart1, bytes32 _publicKeyPart2, uint16 _normalisedWeight) {
        Policy storage policy = policies[_rewardEpochId][_signingPolicyAddress];
        _publicKeyPart1 = policy.pk_1;
        _publicKeyPart2 = policy.pk_2;
        require(_publicKeyPart1 != 0 || _publicKeyPart2 != 0, "Invalid signing policy address");

        uint weightsSum = totalWeights[_rewardEpochId];
        _normalisedWeight = uint16((policy.weight * type(uint16).max) / weightsSum);
    }

    // Combines the functionality of VoterRegistry.registerVoter and EntityManager.registerPublicKey
    // from flare-smart-contracts-v2
    function registerAsVoter(uint epoch, address sender, Policy calldata policy) external {
        require(policy.weight != 0, "Weight must be nonzero");
        policies[epoch][sender] = policy;
        totalWeights[epoch] += policy.weight;
    }
}
