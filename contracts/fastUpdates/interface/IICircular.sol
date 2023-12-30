// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

abstract contract IICircular {
    uint circularLength;

    constructor(uint _l) {
        circularLength = _l;
    }

    function ix(uint i) internal view returns (uint) {
        return (i + block.number) % circularLength;
    }

    function blockIx(uint blockNum, string memory failMsg) internal view returns (uint) {
        require(block.number < blockNum + circularLength, failMsg);
        return ix(blockNum);
    }

    function backIx(uint i) internal view returns (uint) {
        assert(i < circularLength);
        return ix(circularLength - i);
    }

    function prevIx() internal view returns (uint) {
        return backIx(1);
    }

    function thisIx() internal view returns (uint) {
        return ix(0);
    }

    function nextIx() internal view returns (uint) {
        return ix(1);
    }
}