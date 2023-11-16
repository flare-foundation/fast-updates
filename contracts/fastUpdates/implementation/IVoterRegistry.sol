// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

// Copied as a stub from ftso-scaling repo 
abstract contract IVoterRegistry {

    function votersForRewardEpoch(
        uint256 _rewardEpochId
    ) external view virtual returns (address[] memory voters, uint256[] memory weights);

    mapping(uint256 => uint256) public totalWeightPerRewardEpoch;
    
}
