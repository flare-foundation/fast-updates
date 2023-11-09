// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import { ECPoint2 } from "../lib/Sortition.sol";

contract FastUpdateManager {
    uint public submissionWindowLength;

    function setSubmissionWindowLength(uint w) public { // onlyGovernance
        submissionWindowLength = w;
    }

    uint32[1000] public numericDeltas;
    ECPoint2 private ecBasePoint;
    uint public baseSeed;

    uint16[] private activeSampleIncreases;
    function ixS(uint i) private view returns (uint) {
        return (i + block.number) % activeSampleIncreases.length;
    }
    function getSampleIncrease(uint i) private view returns (uint16) {
        assert(i < activeSampleIncreases.length);
        return activeSampleIncreases[ixS(i)];
    }
    function incrementCurrentSampleIncrease(uint16 inc) private {
        activeSampleIncreases[ixS(0)] += inc;
    }
    function setSampleIncreaseDuration(uint d) private {
        delete activeSampleIncreases;
        for (uint i = 0; i < d; ++i) {
            activeSampleIncreases.push();
        }
    }

    uint16[1000][] private activeDeltaIncreases;
    function ixD(uint i) private view returns (uint) {
        return (i + block.number) % activeDeltaIncreases.length;
    }
    function getDeltaIncrease(uint i, uint feed) private view returns (uint16) {
        assert(i < activeDeltaIncreases.length);
        return activeDeltaIncreases[ixD(i)][feed];
    }
    function scaleCurrentDeltaIncrease(uint16 scaleFactor8x8, uint feed) private {
        uint16 currentIncrease8x8 = activeDeltaIncreases[ixD(0)][feed];
        uint32 newIncrease = (uint32(currentIncrease8x8) * uint32(scaleFactor8x8)) >> 8; // fixed point product
        activeDeltaIncreases[ixD(0)][feed] = uint16(newIncrease);
    }
    function setDeltaIncreaseDuration(uint d) private {
        delete activeDeltaIncreases;
        for (uint i = 0; i < d; ++i) {
            activeDeltaIncreases.push();
        }
    }

    uint16 private expectedSampleSize8x8;
    uint16 private constant bitmask8x8low = uint8(int8(-1));

    function setBaseSeed(uint newSeed) public { // onlyGovernance
        baseSeed = newSeed;
    }

    function setExpectedSampleSize(uint16 newSize8x8) public { // onlyGovernance
        expectedSampleSize8x8 = newSize8x8;
        setSampleIncreaseDuration(activeSampleIncreases.length); // clears the array
    }

    function incrementExpectedSampleSize(uint16 inc) public { // only FastUpdateIncentiveManager
        uint24 newSize = uint24(expectedSampleSize8x8) + uint24(inc);
        if (newSize > type(uint16).max) {
            newSize = type(uint16).max;
            inc = uint16(newSize) - expectedSampleSize8x8;
        }
        expectedSampleSize8x8 = uint16(newSize);
        incrementCurrentSampleIncrease(inc);
    }

    function setNumericDeltas(uint32[1000] calldata newDeltas) public { // onlyGovernance
        numericDeltas = newDeltas;
        setDeltaIncreaseDuration(activeDeltaIncreases.length); // clears the array
    }

    function scaleNumericDelta(uint16 scaleFactor8x8, uint feed) public { // only FastUpdateIncentiveManager
        uint32 delta = numericDeltas[feed];
        uint48 newDelta = uint48(delta) * uint48(scaleFactor8x8);
        if (newDelta > type(uint32).max) {
            newDelta = type(uint32).max;
            scaleFactor8x8 = uint16(((newDelta << 16) / uint48(delta)) >> 16); // fixed point division
        }
        numericDeltas[feed] = uint32(newDelta);
        scaleNumericDelta(scaleFactor8x8, feed);
    }

    function getScoreCutoff(uint8 numParticipants) public view returns (uint) { // onlyGovernance
        // The formula is: (exp. s.size)/(num. part.) = (score)/(score range), score range = 2**256
        // So, return (expectedSampleSize << 256)/numParticipants, 
        // being careful about multiplication overflow and integer division:
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