// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import { SortitionCredential } from "../lib/Sortition.sol";
import "../lib/Bn256.sol";
import "../lib/FixedPointArithmetic.sol" as FPA;
import { Deltas } from "../lib/Deltas.sol";

abstract contract IFastUpdater {
    struct ActiveProviderData {
        Bn256.G1Point publicKey;
        uint sortitionWeight;
    }

    mapping (address => ActiveProviderData) public activeProviders;
    address[] public activeProviderAddresses;

    struct FastUpdates {
        uint sortitionBlock;
        SortitionCredential sortitionCredential;
        Deltas deltas;
    }

    function submitUpdates(FastUpdates calldata) external virtual;
    function fetchCurrentPrices(uint[] calldata) external view virtual returns(FPA.Price[] memory);
}