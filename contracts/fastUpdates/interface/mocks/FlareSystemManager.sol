// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

// Copied as a stub from ftso-scaling repo 
interface FlareSystemManager {
    function getCurrentRandom() external view returns(uint256 _currentRandom);
    function getCurrentRewardEpochId() external view returns(uint24 _currentRewardEpochId);
}