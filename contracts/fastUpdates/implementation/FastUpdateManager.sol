// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

contract FastUpdateManager {
    uint public submissionWindowLength;
    uint32[1000] public numericDeltas;

    function getNumericDeltas(
        uint[] calldata feeds
    ) public view returns (uint[] memory feedDeltas) {
        feedDeltas = new uint[](feeds.length);
        for (uint i = 0; i < feeds.length; ++i) {
            feedDeltas[i] = numericDeltas[feeds[i]];
        }
    }
}