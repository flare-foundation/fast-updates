// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

function heapSort(uint[] memory array) pure {
    makeHeap(array);
    unmakeHeap(array);
}

// After this, array is a max-heap
function makeHeap(uint[] memory array) pure {
    for (uint i = 1; i < array.length; ++i) { // Yes, 1
        addToHeap(array, i);
    }
}

// Repeatedly peels off the maximum element, after which array is sorted
function unmakeHeap(uint[] memory array) pure {
    for (uint lastIndex = array.length - 1; lastIndex > 0; --lastIndex) {
        popHeap(array, lastIndex);
    }
}

// Extends the max-heap of the first i-1 elements of array to include the i'th element
function addToHeap(uint[] memory array, uint i) pure {
    while (i > 0) {
        uint parentI = (i - 1)/2;
        uint parent = array[parentI];
        uint value = array[i];
        if (value > parent) {
            array[parentI] = value;
            array[i] = parent;
        }
        else return;
    }
    return;
}

// Reduces the heap by moving the max to the end and fixing so that the rest is still a heap
function popHeap(uint[] memory array, uint lastIndex) pure {
    uint head = array[0];
    uint last = array[lastIndex];
    array[lastIndex] = head;
    array[0] = last;
    restoreHeap(array, lastIndex);
}

// Put the out-of-order first element at the end of the correct path, reestablishing the heap
function restoreHeap(uint[] memory array, uint length) pure {
    uint j = 0;
    while (j < length) {
        uint aIndex = j;
        uint bIndex = 2*j + 1;
        uint cIndex = 2*j + 2;
        uint a = array[aIndex];
        uint b = array[bIndex];
        uint c = array[cIndex];
        if (a > b) {
            if (a > c) return;
            else {
                array[aIndex] = c;
                array[cIndex] = a;
                j = cIndex;
            }
        }
        else {
            if (b < c) {
                array[aIndex] = c;
                array[cIndex] = a;
                j = cIndex;
            }
            else {
                array[aIndex] = b;
                array[bIndex] = a;
                j = bIndex;
            }
        }
    }
}