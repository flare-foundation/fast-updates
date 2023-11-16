// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import "hardhat/console.sol";
import "./Bn256.sol";
import {SortitionRound, SortitionCredential, VerifySortitionCredential, VerifySortitionProof} from "./Sortition.sol";

contract TestSortitionContract {
    function TestVerifySortitionCredential(
        SortitionRound memory sortitionRound,
        Bn256.G1Point memory pubKey,
        SortitionCredential memory sortitionCredential
    ) public view returns (bool) {
        bool check;
        uint256 score;
        (check, score) = VerifySortitionCredential(sortitionRound, pubKey, 0, sortitionCredential);
        return check;
    }

    function TestVerifySortitionProof(
        uint256 seed,
        Bn256.G1Point memory pubKey,
        SortitionCredential memory sortitionCredential
    ) public view returns (bool) {
        return VerifySortitionProof(seed, pubKey, sortitionCredential);
    }
}
