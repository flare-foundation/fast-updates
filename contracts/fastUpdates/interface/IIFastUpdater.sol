// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import {IFastUpdater} from "./IFastUpdater.sol";
import {IIFastUpdateIncentiveManager} from "./IIFastUpdateIncentiveManager.sol";
import "./mocks/FlareSystemManager.sol";
import "./mocks/VoterRegistry.sol";
import {Governed} from "./Governed.sol";
import "../lib/FixedPointArithmetic.sol" as FPA;

abstract contract IIFastUpdater is IFastUpdater, Governed {
    VoterRegistry voterRegistry;
    FlareSystemManager flareSystemManager;
    IIFastUpdateIncentiveManager internal fastUpdateIncentiveManager;

    constructor(
        address _governance,
        VoterRegistry _voterRegistry,
        FlareSystemManager _flareSystemManager,
        IIFastUpdateIncentiveManager _fastUpdateIncentiveManager
    ) Governed(_governance) {
        _setVoterRegistry(_voterRegistry);
        _setFlareSystemManager(_flareSystemManager);
        _setFastUpdateIncentiveManager(_fastUpdateIncentiveManager);
    }

    function freeSubmitted() external virtual;

    function setSubmissionWindow(uint _submissionWindow) external virtual;

    function setPrices(FPA.Price[] memory _prices) external virtual;

    function setVoterRegistry(VoterRegistry _voterRegistry) external onlyGovernance {
        _setVoterRegistry(_voterRegistry);
    }

    function _setVoterRegistry(VoterRegistry _voterRegistry) private {
        voterRegistry = _voterRegistry;
    }

    function setFlareSystemManager(FlareSystemManager _flareSystemManager) external onlyGovernance {
        _setFlareSystemManager(_flareSystemManager);
    }

    function _setFlareSystemManager(FlareSystemManager _flareSystemManager) private {
        flareSystemManager = _flareSystemManager;
    }

    function setFastUpdateIncentiveManager(
        IIFastUpdateIncentiveManager _newFastUpdateIncentiveManager
    ) external onlyGovernance {
        _setFastUpdateIncentiveManager(_newFastUpdateIncentiveManager);
    }

    function _setFastUpdateIncentiveManager(IIFastUpdateIncentiveManager _newFastUpdateIncentiveManager) private {
        fastUpdateIncentiveManager = _newFastUpdateIncentiveManager;
    }
}
