// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import { ECPoint, SortitionCredential } from "../lib/Sortition.sol";

abstract contract IFastUpdaters {
    struct NewProvider {
        ECPoint publicKey;
        SortitionCredential credential;
    }

    function registerNewProvider(NewProvider calldata) external virtual;
}