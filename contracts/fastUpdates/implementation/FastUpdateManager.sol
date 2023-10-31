// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import { ECPoint2 } from "../lib/Sortition.sol";

contract FastUpdateManager {
    uint public submissionWindowLength;
    uint32[1000] public numericDeltas;
    ECPoint2 private ecBasePoint;

    function getECBasePoint() public view returns (ECPoint2 memory) {
        return ecBasePoint;
    }

    function getNumericDeltas(
        uint[] calldata feeds
    ) public view returns (uint[] memory feedDeltas) {
        feedDeltas = new uint[](feeds.length);
        for (uint i = 0; i < feeds.length; ++i) {
            feedDeltas[i] = numericDeltas[feeds[i]];
        }
    }
}