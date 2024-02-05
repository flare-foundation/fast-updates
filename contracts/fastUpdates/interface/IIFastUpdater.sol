// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import { IFastUpdater } from "./IFastUpdater.sol";
import { IIFastUpdateIncentiveManager } from "./IIFastUpdateIncentiveManager.sol";
import "./mocks/FlareSystemManager.sol";
import "./mocks/VoterRegistry.sol";

abstract contract IIFastUpdater is IFastUpdater {
    VoterRegistry voterRegistry;
    FlareSystemManager flareSystemManager;
    IIFastUpdateIncentiveManager internal fastUpdateIncentiveManager;

    constructor(
        VoterRegistry _voterRegistry,
        FlareSystemManager _flareSystemManager,
        IIFastUpdateIncentiveManager _fastUpdateIncentiveManager
    ) {
        setVoterRegistry(_voterRegistry);
        setFlareSystemManager(_flareSystemManager);
        setFastUpdateIncentiveManager(_fastUpdateIncentiveManager);
    }

    function setVoterRegistry(VoterRegistry _voterRegistry) public { // only governance
        voterRegistry = _voterRegistry;
    }

    function setFlareSystemManager(FlareSystemManager _flareSystemManager) public { // only governance
        flareSystemManager = _flareSystemManager;    
    }

    function setFastUpdateIncentiveManager(IIFastUpdateIncentiveManager newFastUpdateIncentiveManager) public { // only governance
        fastUpdateIncentiveManager = newFastUpdateIncentiveManager;
    }
}
