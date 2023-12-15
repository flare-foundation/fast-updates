// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import { IFastUpdater } from "./IFastUpdater.sol";
import { IIFastUpdaters } from "./IIFastUpdaters.sol";
import { IIFastUpdateIncentiveManager } from "./IIFastUpdateIncentiveManager.sol";

abstract contract IIFastUpdater is IFastUpdater {
    IIFastUpdaters internal fastUpdaters;
    IIFastUpdateIncentiveManager internal fastUpdateIncentiveManager;

    constructor(IIFastUpdaters _fastUpdaters, IIFastUpdateIncentiveManager _fastUpdateIncentiveManager) {
        setFastUpdaters(_fastUpdaters);
        setFastUpdateIncentiveManager(_fastUpdateIncentiveManager);
    }

    function setFastUpdaters(IIFastUpdaters newFastUpdaters) public { // only governance
        fastUpdaters = newFastUpdaters;
    }

    function setFastUpdateIncentiveManager(IIFastUpdateIncentiveManager newFastUpdateIncentiveManager) public { // only governance
        fastUpdateIncentiveManager = newFastUpdateIncentiveManager;
    }

    function setSubmissionWindow(uint) public virtual;

    function finalizeBlock() public virtual;
}