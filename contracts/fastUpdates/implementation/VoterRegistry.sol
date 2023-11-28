// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import "../interface/IVoterRegistry.sol";

contract VoterRegistry is IVoterRegistry {
    // rewardEpochId => voter => weight
    mapping(uint256 => mapping(address => uint256)) public weightForRewardEpoch;

    // mapping(uint256 => uint256) public totalWeightPerRewardEpoch;
    mapping(uint256 => address[]) public rewardEpochToAllVoters;

    function registerAsAVoter(uint256 _rewardEpochId, uint256 _weight) public {
        require(weightForRewardEpoch[_rewardEpochId][msg.sender] == 0, "voter already registered");
        weightForRewardEpoch[_rewardEpochId][msg.sender] = _weight;
        rewardEpochToAllVoters[_rewardEpochId].push(msg.sender);
        totalWeightPerRewardEpoch[_rewardEpochId] += _weight;
    }

    function getVoterWeightForRewardEpoch(address _voter, uint256 _rewardEpochId) public view returns (uint256) {
        return weightForRewardEpoch[_rewardEpochId][_voter];
    }

    function votersForRewardEpoch(
        uint256 _rewardEpochId
    ) public view override returns (address[] memory voters, uint256[] memory weights) {
        voters = rewardEpochToAllVoters[_rewardEpochId];
        weights = voterWeightsInRewardEpoch(_rewardEpochId, voters);
    }

    function voterWeightsInRewardEpoch(
        uint256 _rewardEpochId,
        address[] memory _voters
    ) internal view returns (uint256[] memory weights) {
        weights = new uint256[](_voters.length);
        for (uint256 i = 0; i < _voters.length; i++) {
            weights[i] = getVoterWeightForRewardEpoch(_voters[i], _rewardEpochId);
        }
    }
}
