// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

// Copied as a stub from ftso-scaling repo
interface IVoterRegistry {

    function votersForRewardEpoch(
        uint256 _rewardEpochId
    ) external view returns (address[] memory, uint256[] memory);

    function thresholdForRewardEpoch(
        uint256 _rewardEpochId
    ) external view returns (uint256);

    function getVoterWeightForRewardEpoch(
        address _voter,
        uint256 _rewardEpochId
    ) external view returns (uint256);

    function getDelegatorWeightForRewardEpochAndVoter(
        address _delegator,
        address _voter,
        uint256 _rewardEpochId
    ) external view returns (uint256);
}
