// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import { ECPoint2 } from "../lib/Sortition.sol";

contract FastUpdateManager {
    uint public submissionWindowLength;
    uint32[1000] public numericDeltas;
    ECPoint2 private ecBasePoint;
    uint public baseSeed;

    uint16 private expectedSampleSize8x8;
    uint16 private constant bitmask8x8low = uint8(int8(-1));

    function setBaseSeed(uint newSeed) public { // onlyGovernance
        baseSeed = newSeed;
    }

    function setExpectedSampleSize(uint16 newSize8x8) public { // onlyGovernance
        expectedSampleSize8x8 = newSize8x8;
    }

    function getScoreCutoff(uint8 numParticipants) public view returns (uint) { // onlyGovernance
        // Being careful about multiplication overflow and integer division
        //
        // e<<256 / n = (eH<<256 + eL<<248) / n = (eL<<248 + eH * (2**256 % n))/ n + eH * (2**256//n)
        // 2**256 % n = (uint.max % n + 1) % n
        // 2**256//n  = uint.max//n + (uint.max % n == -1)

        uint eH = expectedSampleSize8x8 >> 8;
        uint eL = expectedSampleSize8x8 & bitmask8x8low;

        uint u = type(uint).max % numParticipants;
        uint x = (u + 1) % numParticipants;
        uint y = type(uint).max / numParticipants + (u == numParticipants - 1 ? 1 : 0);

        return (eL << 248 + eH * x) / numParticipants + eH * y;
    }

    // This is because Solidity's autogen'd getters intentionally screw up returned structs
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