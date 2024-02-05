// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import "hardhat/console.sol";
import "../lib/Bn256.sol";
import {SortitionState, SortitionCredential, verifySortitionCredential, verifySortitionProof} from "../lib/Sortition.sol";

contract TestSortitionContract {
    function testVerifySortitionCredential(
        SortitionState calldata sortitionState,
        SortitionCredential calldata sortitionCredential
    ) public view returns (bool) {
        bool check;
        uint256 score;
        (check, score) = verifySortitionCredential(sortitionState, sortitionCredential);
        return check;
    }

    function testVerifySortitionProof(
        SortitionState calldata sortitionState,
        SortitionCredential calldata sortitionCredential
    ) public view returns (bool) {
        return verifySortitionProof(sortitionState, sortitionCredential);
    }
}
