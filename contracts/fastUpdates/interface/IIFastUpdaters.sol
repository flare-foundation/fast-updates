// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import { IFastUpdaters } from "./IFastUpdaters.sol";
import { IVoterRegistry } from "./IVoterRegistry.sol";
import { ECPoint } from "../lib/Sortition.sol";

abstract contract IIFastUpdaters is IFastUpdaters {
    IVoterRegistry voterRegistry;

    function setVoterRegistry(IVoterRegistry newVoterRegistry) external {
        voterRegistry = newVoterRegistry;
    }

    struct ProviderRegistry {
        uint seed;
        address[] providerAddresses;
        ECPoint[] providerKeys;
        uint[] providerWeights;
    }

    function nextProviderRegistry(uint epochId) public virtual returns(ProviderRegistry memory);
}