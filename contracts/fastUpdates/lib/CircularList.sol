// SPDX-License-Identifier: MIT
// Implements a list with wrap-around indexing, for uint16 elements fitting in one storage slot
// This is a pragmatic representation of a FIFO that is rotated one notch each block
// Allows look-back for up to 16 blocks (16 = 256 / 16) with operations on "current" element
pragma solidity 0.8.18;

function ix(uint16[] storage list, uint i) view returns (uint) {
    return (i + block.number) % list.length;
}

function clear(uint16[] storage list) {
    for (uint i = 0; i < list.length; ++i) {
        list[i] = 0;
    }
}

function circularGet16(uint16[] storage list, uint i) view returns (uint16) {
    assert(i < list.length);
    return list[list.ix(0)];
}

function circularHead16(uint16[] storage list) view returns (uint16) {
    return circularGet16(list, 0);
}

function circularZero16(uint16[] storage list) {
    list[list.ix(0)] = 0;
}

function circularAdd16(uint16[] storage list, uint16 inc) {
    uint32 newVal = list[list.ix(0)] + inc;
    if (newVal > type(uint16).max) newVal = type(uint16).max;
    list[list.ix(0)] = uint16(newVal);
}

function circularResize(uint16[] storage list, uint newSize) {
    assert(newSize <= 16);

    list.clear();
    if (newSize < list.length) {
        for (uint i = 0; i < list.length - newSize; ++i) {
            list.pop();
        }
    }
    else {
        for (uint i = 0; i < newSize - list.length; ++i) {
            list.push();
        }
    }
}

function sum(uint16[] storage list) view returns (uint16 result) {
    for (uint i = 0; i < list.length; ++i) {
        result += list[i];
    }
}

using { ix, clear } for uint16[];