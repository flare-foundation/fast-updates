// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import {SortitionCredential} from "../lib/Sortition.sol";
import "../lib/Bn256.sol";

abstract contract IFastUpdaters {
    struct NewProvider {
        Bn256.G1Point publicKey;
        SortitionCredential credential;
    }

    event NewProviderKey(
        uint indexed rewardEpochId,
        address indexed providerAddress,
        Bn256.G1Point indexed providerPublicKey
    );

    function registerNewProvider(NewProvider calldata) external virtual;
}
