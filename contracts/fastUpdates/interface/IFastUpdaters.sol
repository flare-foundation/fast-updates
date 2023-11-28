// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import {SortitionCredential} from "../lib/Sortition.sol";
import "../lib/Bn256.sol";

abstract contract IFastUpdaters {
    struct NewProvider {
        Bn256.G1Point publicKey;
        SortitionCredential credential;
    }

    function registerNewProvider(NewProvider calldata) external virtual;
}
