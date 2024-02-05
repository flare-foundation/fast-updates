// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import { SortitionCredential } from "../lib/Sortition.sol";
import "../lib/Bn256.sol";
import "../lib/FixedPointArithmetic.sol" as FPA;
import { Deltas } from "../lib/Deltas.sol";

interface IFastUpdater {
    struct FastUpdates {
        uint sortitionBlock;
        SortitionCredential sortitionCredential;
        Deltas deltas;
    }

    event FastUpdate(
        address indexed providerAddress
    );

    function submitUpdates(FastUpdates calldata) external;
    function fetchCurrentPrices(uint[] calldata) external view  returns(FPA.Price[] memory);
    function currentScoreCutoff() external view returns (uint);
    function currentSortitionWeight(address voter) external view returns (uint);
}