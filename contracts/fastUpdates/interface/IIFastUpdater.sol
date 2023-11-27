// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import { IFastUpdater } from "./IFastUpdater.sol";
import { IIFastUpdaters } from "./IIFastUpdaters.sol";
import { IIFastUpdateIncentiveManager } from "./IIFastUpdateIncentiveManager.sol";

abstract contract IIFastUpdater is IFastUpdater {
    IIFastUpdaters internal fastUpdaters;
    IIFastUpdateIncentiveManager internal fastUpdateIncentiveManager;

    function setFastUpdaters(IIFastUpdaters newFastUpdaters) public { // only governance
        fastUpdaters = newFastUpdaters;
    }

    function setFastUpdateIncentiveManager(IIFastUpdateIncentiveManager newFastUpdateIncentiveManager) public { // only governance
        fastUpdateIncentiveManager = newFastUpdateIncentiveManager;
    }

    function setSubmissionWindow(uint) external virtual;

    function finalizeBlock() public virtual;
}