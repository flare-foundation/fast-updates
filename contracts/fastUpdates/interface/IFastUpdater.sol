// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import { ECPoint, SortitionCredential } from "../lib/Sortition.sol";
import { Deltas } from "../lib/Deltas.sol";

abstract contract IFastUpdater {
    struct ActiveProviderData {
        ECPoint publicKey;
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
    function fetchCurrentPrices(uint[] calldata) external view virtual returns(uint[] memory);
}