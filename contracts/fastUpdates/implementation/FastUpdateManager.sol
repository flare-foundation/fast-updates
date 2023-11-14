// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import { ECPoint2 } from "../lib/Sortition.sol";
import { VIRTUAL_PROVIDER_BITS } from "./FastUpdaters.sol";

contract FastUpdateManager {
    uint public submissionWindowLength;

    function setSubmissionWindowLength(uint w) public { // onlyGovernance
        submissionWindowLength = w;
    }

    uint32[1000] public numericDeltas;

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

    function getScoreCutoff() public view returns (uint) { // onlyGovernance
        // The formula is: (exp. s.size)/(num. prov.) = (score)/(score range)
        //   score range = 2**256
        //   num. prov.  = 2**VIRTUAL_PROVIDER_BITS
        //   exp. s.size = "expectedSampleSize8x8 >> 8", in that we keep the fractional bits:
        return uint(expectedSampleSize8x8) << (256 - VIRTUAL_PROVIDER_BITS - 8);
    }

    function getNumericDeltas(
        uint[] calldata feeds
    ) public view returns (uint[] memory feedDeltas) {
        feedDeltas = new uint[](feeds.length);
        for (uint i = 0; i < feeds.length; ++i) {
            feedDeltas[i] = numericDeltas[feeds[i]];
        }
    }

    function finalizeBlock() public { // only governance
        expectedSampleSize8x8 -= getSampleIncrease(0);
        for (uint feed = 0; feed < activeDeltaIncreases.length; ++feed) {
            activeDeltaIncreases[0][feed] -= getDeltaIncrease(0, feed);
        }
    }
}