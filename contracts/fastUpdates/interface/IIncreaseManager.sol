// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;
pragma abicoder v2;

interface IIncreaseManager {
    function getIncentiveDuration() external view returns(uint);
    function setIncentiveDuration(uint _dur) external;
}