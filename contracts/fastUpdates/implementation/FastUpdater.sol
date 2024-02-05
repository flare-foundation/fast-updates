// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;
pragma abicoder v2;

import {FastUpdateIncentiveManager} from "./FastUpdateIncentiveManager.sol";
import {Deltas} from "../lib/Deltas.sol";
import {SortitionState, SortitionCredential, verifySortitionCredential} from "../lib/Sortition.sol";
import {IIFastUpdater} from "../interface/IIFastUpdater.sol";
import {IIFastUpdateIncentiveManager} from "../interface/IIFastUpdateIncentiveManager.sol";
import {CircularListManager} from "./abstract/CircularListManager.sol";
import "../interface/mocks/FlareSystemManager.sol";
import "../interface/mocks/VoterRegistry.sol";
import "../lib/FixedPointArithmetic.sol" as FPA;
import "../lib/Bn256.sol";
import "hardhat/console.sol";

// The number of units of weight distributed among providers is 1 << VIRTUAL_PROVIDER_BITS
uint constant VIRTUAL_PROVIDER_BITS = 12;
// value 128 bellow can be replaced with x such that x >=  numBits(FPA.SampleSize) + VIRTUAL_PROVIDER_BITS
// and x <= 256 - numBits(FPA.SampleSize)
uint constant UINT_SPLIT = 128;
uint constant SMALL_P = Bn256.p & (2 ** (UINT_SPLIT) - 1);
uint constant BIG_P = Bn256.p >> UINT_SPLIT;

struct Submitted {
    uint[] gammaXs;
    Bn256.G1Point[] pks; // todo: we could also have just pkXs or even avoided
}

contract FastUpdater is IIFastUpdater, CircularListManager {
    FPA.Price[1000] public prices;
    uint submissionWindow;

    Submitted[] internal submitted;

    constructor(
        VoterRegistry _voterRegistry,
        FlareSystemManager _flareSystemManager,
        IIFastUpdateIncentiveManager _fastUpdateIncentiveManager,
        FPA.Price[] memory _prices,
        uint _submissionWindow
    )
        IIFastUpdater(_voterRegistry, _flareSystemManager, _fastUpdateIncentiveManager)
        CircularListManager(_submissionWindow + 1)
    {
        setSubmissionWindow(_submissionWindow);
        setPrices(_prices);
        setSubmitted();
    }

    function setSubmissionWindow(uint _submissionWindow) public {
        // only governance
        submissionWindow = _submissionWindow;
    }

    function setPrices(FPA.Price[] memory _prices) public {
        // only governance
        for (uint i = 0; i < _prices.length; ++i) {
            prices[i] = _prices[i];
        }
    }

    function setSubmitted() public {
        // only governance
        for (uint i = 0; i < circularLength; ++i) {
            submitted.push();
        }
    }

    function freeSubmitted() public {
        // only governance
        delete submitted[nextIx()];
    }

    function getSubmitted(uint blockNum) internal view returns (Submitted storage submittedI) {
        string memory failMsg = "Sortition round for the given block is no longer or not yet available";
        uint ix = blockIx(blockNum, failMsg);
        submittedI = submitted[ix];
    }

    function submitUpdates(FastUpdates calldata updates) external override {
        require(
            block.number < updates.sortitionBlock + submissionWindow,
            "Updates no longer accepted for the given block"
        );
        require(block.number >= updates.sortitionBlock, "Updates not yet available for the given block");

        (Bn256.G1Point memory key, uint weight) = _providerData(msg.sender);
        SortitionState memory sortitionState = SortitionState({
            baseSeed: flareSystemManager.getCurrentRandom(),
            blockNumber: updates.sortitionBlock,
            scoreCutoff: currentScoreCutoff(),
            weight: weight,
            pubKey: key
        });

        Submitted storage submittedI = getSubmitted(updates.sortitionBlock);
        for (uint j = 0; j < submittedI.gammaXs.length; j++) {
            if (submittedI.gammaXs[j] == updates.sortitionCredential.gamma.x) {
                if (submittedI.pks[j].x == key.x && submittedI.pks[j].y == key.y) {
                    revert("submission already provided");
                }
            }
        }

        submittedI.gammaXs.push(updates.sortitionCredential.gamma.x);
        submittedI.pks.push(key);

        bool check;
        (check, ) = verifySortitionCredential(sortitionState, updates.sortitionCredential);
        require(check, "sortition proof invalid");

        FPA.Scale scale = fastUpdateIncentiveManager.getScale();
        applyUpdates(scale, updates.deltas);
        emit FastUpdate(msg.sender);
    }

    function _providerData(address voter) private view returns (Bn256.G1Point memory key, uint weight) {
        uint epochId = flareSystemManager.getCurrentRewardEpochId();
        (bytes32 pk_1, bytes32 pk_2, uint16 normalizedWeight) = voterRegistry.getPublicKeyAndNormalisedWeight(
            epochId,
            voter
        );
        key = Bn256.G1Point(uint(pk_1), uint(pk_2));
        weight = (uint(normalizedWeight) << VIRTUAL_PROVIDER_BITS) >> 16;
    }

    function currentSortitionWeight(address voter) public view override returns (uint weight) {
        (, weight) = _providerData(voter);
    }

    function currentScoreCutoff() public view override returns (uint cutoff) {
        FPA.SampleSize expectedSampleSize = fastUpdateIncentiveManager.getExpectedSampleSize();
        // The formula is: (exp. s.size)/(num. prov.) = (score)/(score range)
        //   score range = p = 21888242871839275222246405745257275088696311157297823662689037894645226208583
        //   num. prov.  = 2**VIRTUAL_PROVIDER_BITS
        //   exp. s.size = "expectedSampleSize8x8 >> 8", in that we keep the fractional bits:
        cutoff = (BIG_P * uint(FPA.SampleSize.unwrap(expectedSampleSize))) << (UINT_SPLIT - 8 - VIRTUAL_PROVIDER_BITS);
        cutoff += (SMALL_P * uint(FPA.SampleSize.unwrap(expectedSampleSize))) >> (8 + VIRTUAL_PROVIDER_BITS);
    }

    function fetchCurrentPrices(uint[] calldata feeds) external view override returns (FPA.Price[] memory _prices) {
        _prices = new FPA.Price[](feeds.length);
        for (uint i = 0; i < feeds.length; ++i) {
            _prices[i] = prices[feeds[i]];
        }
    }

    function applyUpdates(FPA.Scale scale, Deltas calldata deltas) private {
        deltas.forEach(scale, applyDelta); // TODO: optimize these calls for storage access
    }

    function applyDelta(int delta, uint feed, FPA.Scale scale) private {
        if (delta == -1) {
            prices[feed] = FPA.div(prices[feed], scale);
        } else if (delta == 1) {
            prices[feed] = FPA.mul(prices[feed], scale);
        }
    }
}
