// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

abstract contract IICircular {
    uint circularLength;

    constructor(uint _l) {
        setCircularLength(_l);
    }

    function setCircularLength(uint _l) internal {
        circularLength = _l;
    }

    function ix(uint i) internal view returns (uint) {
        return (i + block.number) % circularLength;
    }

    function blockIx(uint blockNum, string memory failMsg) internal view returns (uint) {
        require(blockNum <= block.number && block.number < blockNum + circularLength, failMsg);
        uint blocksAgo = block.number - blockNum;
        return backIx(blocksAgo);
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