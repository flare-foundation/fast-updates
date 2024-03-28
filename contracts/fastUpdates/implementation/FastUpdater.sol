// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;
pragma abicoder v2;

import {FastUpdateIncentiveManager} from "./FastUpdateIncentiveManager.sol";
import {SortitionState, SortitionCredential, verifySortitionCredential} from "../lib/Sortition.sol";
import {IIFastUpdater} from "../interface/IIFastUpdater.sol";
import {IIFastUpdateIncentiveManager} from "../interface/IIFastUpdateIncentiveManager.sol";
import {CircularListManager} from "./abstract/CircularListManager.sol";
import "../interface/mocks/FlareSystemManager.sol";
import "../interface/mocks/VoterRegistry.sol";
import "../lib/FixedPointArithmetic.sol" as FPA;
import "../lib/Bn256.sol";
import {recoverSigner} from "../lib/Signature.sol";
import {Signature} from "../lib/Signature.sol";


// The number of units of weight distributed among providers is 1 << VIRTUAL_PROVIDER_BITS
uint constant VIRTUAL_PROVIDER_BITS = 12;
// value 128 bellow can be replaced with x such that x >=  numBits(FPA.SampleSize) + VIRTUAL_PROVIDER_BITS
// and x <= 256 - numBits(FPA.SampleSize)
uint constant UINT_SPLIT = 128;
uint constant SMALL_P = Bn256.p & (2 ** (UINT_SPLIT) - 1);
uint constant BIG_P = Bn256.p >> UINT_SPLIT;

struct SubmittedHashes {
    bytes32[] hashes;
}

contract FastUpdater is IIFastUpdater, CircularListManager {
    bytes32[] private prices;
    uint submissionWindow;

    SubmittedHashes[] internal submitted;
    bytes[] internal submittedDeltas;
    uint currentDelta;
    uint backlogDelta;

    constructor(
        address _governance,
        VoterRegistry _voterRegistry,
        FlareSystemManager _flareSystemManager,
        IIFastUpdateIncentiveManager _fastUpdateIncentiveManager,
        FPA.Price[] memory _prices,
        uint _submissionWindow,
        uint _deltasBacklog
    )
        IIFastUpdater(_governance, _voterRegistry, _flareSystemManager, _fastUpdateIncentiveManager)
        CircularListManager(_submissionWindow + 1)
    {
        _setPrices(_prices);

        _setSubmissionWindow(_submissionWindow);
        _initSubmitted();

        for (uint slot = 0; slot < _deltasBacklog; ++slot) {
            submittedDeltas.push();
        }
    }

    function setSubmissionWindow(uint _submissionWindow) external override onlyGovernance {
        _setSubmissionWindow(_submissionWindow);
    }

    function _setSubmissionWindow(uint _submissionWindow) private {
        submissionWindow = _submissionWindow;
    }

    function setPrices(FPA.Price[] memory _prices) external override onlyGovernance {
        _setPrices(_prices);
    }

    function _setPrices(FPA.Price[] memory _prices) private {
        prices = new bytes32[](((_prices.length - 1) / 8) + 1);

        uint feed;
        for (uint slot = 0; slot < prices.length; ++slot) {
            bytes32 newSlot;
            for (uint entry = 0; entry < 8; ++entry) {
                newSlot <<= 32;
                if (feed < _prices.length) {
                    newSlot |= bytes32(uint(FPA.Price.unwrap(_prices[feed])));
                }
                ++feed;
            }
            prices[slot] = newSlot;
            if (feed >= _prices.length) {
                break;
            }
        }
    }

    function _initSubmitted() private {
        for (uint i = 0; i < circularLength; ++i) {
            submitted.push();
        }
    }

    function freeSubmitted() external override onlyGovernance {
        delete submitted[nextIx()];
    }

    function getSubmitted(uint blockNum) private view returns (SubmittedHashes storage submittedI) {
        string memory failMsg = "Sortition round for the given block is no longer or not yet available";
        uint ix = blockIx(blockNum, failMsg);
        submittedI = submitted[ix];
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

    function currentSortitionWeight(address voter) external view returns (uint weight) {
        (, weight) = _providerData(voter);
    }

    function currentScoreCutoff() public view returns (uint cutoff) {
        FPA.SampleSize expectedSampleSize = fastUpdateIncentiveManager.getExpectedSampleSize();
        // The formula is: (exp. s.size)/(num. prov.) = (score)/(score range)
        //   score range = p = 21888242871839275222246405745257275088696311157297823662689037894645226208583
        //   num. providers  = 2**VIRTUAL_PROVIDER_BITS
        //   exp. sample size = "expectedSampleSize8x8 >> 8", in that we keep the fractional bits:
        cutoff = (BIG_P * uint(FPA.SampleSize.unwrap(expectedSampleSize))) << (UINT_SPLIT - 8 - VIRTUAL_PROVIDER_BITS);
        cutoff += (SMALL_P * uint(FPA.SampleSize.unwrap(expectedSampleSize))) >> (8 + VIRTUAL_PROVIDER_BITS);
    }

    function submitUpdates(FastUpdates calldata updates) external {
        require(
            block.number < updates.sortitionBlock + submissionWindow,
            "Updates no longer accepted for the given block"
        );
        require(block.number >= updates.sortitionBlock, "Updates not yet available for the given block");
        require((updates.deltas.length * 4) <= prices.length * 8, "More updates than available prices");

        bytes32 msgHashed = sha256(abi.encode(updates.sortitionBlock, updates.sortitionCredential, updates.deltas));
        // console.log(uint(msgHashed));
        address sender = recoverSigner(msgHashed, updates.signature);

        (Bn256.G1Point memory key, uint weight) = _providerData(sender);
        SortitionState memory sortitionState = SortitionState({
            baseSeed: flareSystemManager.getCurrentRandom(),
            blockNumber: updates.sortitionBlock,
            scoreCutoff: currentScoreCutoff(),
            weight: weight,
            pubKey: key
        });

        SubmittedHashes storage submittedI = getSubmitted(updates.sortitionBlock);
        bytes32 hashedRandomness = sha256(abi.encode(key, updates.sortitionBlock, updates.sortitionCredential.replicate));

        for (uint j = 0; j < submittedI.hashes.length; j++) {
            if (submittedI.hashes[j] == hashedRandomness) {
                revert("submission already provided");
            }
        }
        submittedI.hashes.push(hashedRandomness);

        (bool check, ) = verifySortitionCredential(sortitionState, updates.sortitionCredential);
        require(check, "sortition proof invalid");

        submitDeltas(updates.deltas);

        emit FastUpdate(msg.sender);
    }

    function submitDeltas(bytes calldata deltas) internal {
        submittedDeltas[currentDelta] = deltas;
        currentDelta = (currentDelta + 1) % submittedDeltas.length;
    }


    function applySubmitted() external onlyGovernance {
        FPA.Scale scale = fastUpdateIncentiveManager.getScale();
        assembly {
            let arg := mload(0x40)
            let slot := add(arg, 0x20)
            let loacationDeltas := add(arg, 0x40)
            let loacationPrices := add(arg, 0x60)
            mstore(add(arg, 0x80), sload(submittedDeltas.slot)) // add(arg, 0x80) not saved in a variable to avoid the stack too long error
            mstore(add(arg, 0xa0), sload(currentDelta.slot)) // add(arg, 0xa0) not saved in a variable to avoid the stack too long error
            let deltaReduced
            let price
            let priceReduced
            let newPrice
            let delta
            let deltasLen

            mstore(loacationPrices, prices.slot)
            mstore(loacationPrices, keccak256(loacationPrices, 32))
            mstore(slot, submittedDeltas.slot)

            // iterate over updates in the backlog
            for { let i := sload(backlogDelta.slot) } iszero(eq(i, mload(add(arg, 0xa0)))) { i := mod(add(i, 1),  mload(add(arg, 0x80)))} {
                // get location of the i-th update in the updates backlog
                mstore(loacationDeltas, add(keccak256(slot, 32), i))
                deltasLen := sload(mload(loacationDeltas))

                // data is stored differnetly if there is more or less than 31 bytes
                // more than 31 bytes
                if eq(mod(deltasLen, 2), 1) {
                    mstore(loacationDeltas, keccak256(loacationDeltas, 32))
                    deltasLen := div(deltasLen, 2)
                    for { let j := 0 } lt(j, add(div(sub(deltasLen, 1), 32), 1)) { j := add(j, 1) } {
                        // load from storage a bytes32 element containing 128 deltas
                        delta := sload(add(mload(loacationDeltas), j))

                        // the delta consists of 128 updates saved in 256 bits
                        for { let k := 0 } lt(k, 256) {} {
                            if iszero(lt(add(mul(j, 32), div(k, 8)), deltasLen)) {
                                break
                            }
                            // load from storage the value that covers 8 prices 
                            price := sload(add(mload(loacationPrices), add(mul(j, 16), div(k, 16))))

                            // use 8 updates to change 8 prices
                            newPrice := 0
                            for { let l := 0 } lt(l, 8) {l := add(l, 1) } {
                                priceReduced := shl(mul(l, 32), price)
                                priceReduced := shr(224, priceReduced)

                                if lt(add(mul(j, 32), div(k, 8)), deltasLen) {
                                    deltaReduced := shl(k, delta)
                                    deltaReduced := shr(254, deltaReduced)

                                    if eq(deltaReduced, 1) {
                                        // mul
                                        priceReduced := mul(priceReduced, scale)
                                        priceReduced := shr(15, priceReduced)
                                    }
                                    if eq(deltaReduced, 3) {
                                        // div
                                        priceReduced := shl(15, priceReduced)
                                        priceReduced := div(priceReduced, scale)
                                    }
                                }
                                priceReduced := shl(sub(224, mul(l, 32)), priceReduced)
                                newPrice := or(newPrice, priceReduced)

                                k:= add(k, 2)
                            }

                            // store a new value covering 8 prices
                            sstore(add(mload(loacationPrices), add(mul(j, 16), div(sub(k, 16), 16))), newPrice)
                        }
                    }
                    deltasLen := 1 // to avoid executing the other if case
                }

                // less than 32 bytes
                if eq(mod(deltasLen, 2), 0) {
                    delta := deltasLen
                    deltasLen := div(mod(deltasLen, 64), 2)

                    // the delta consists of 128 updates saved in 256 bits
                    for { let k := 0 } lt(k, 256) {} {
                        if iszero(lt(div(k, 8), deltasLen)) {
                            break
                        }
                        // load from storage the value that covers 8 prices 
                        price := sload(add(mload(loacationPrices), div(k, 16)))

                        // use 8 updates to change 8 prices
                        newPrice := 0
                        for { let l := 0 } lt(l, 8) {l := add(l, 1) } {
                            priceReduced := shl(mul(l, 32), price)
                            priceReduced := shr(224, priceReduced)
                            if lt(div(k, 8), deltasLen) {
                                deltaReduced := shl(k, delta)
                                deltaReduced := shr(254, deltaReduced)
                                if eq(deltaReduced, 1) {
                                    // mul
                                    priceReduced := mul(priceReduced, scale)
                                    priceReduced := shr(15, priceReduced)
                                }
                                if eq(deltaReduced, 3) {
                                    // div
                                    priceReduced := shl(15, priceReduced)
                                    priceReduced := div(priceReduced, scale)
                                }
                            }
                            priceReduced := shl(sub(224, mul(l, 32)), priceReduced)
                            newPrice := or(newPrice, priceReduced)

                            k:= add(k, 2)
                        }

                        // store a new value covering 8 prices
                        sstore(add(mload(loacationPrices), div(sub(k, 16), 16)), newPrice)
                    }
                }
            }
            // change backlogDelta variable to equal currentDelta
            sstore(backlogDelta.slot, mload(add(arg, 0xa0)))
        }
    }

    // feeds should be sorted for better performance
    function fetchCurrentPrices(uint[] calldata feeds) external view returns (uint256[] memory _prices) {
        _prices = new uint256[](feeds.length);

        FPA.Scale scale = fastUpdateIncentiveManager.getScale();
    
        assembly {
            let arg := mload(0x40)
            // let arg2 := add(arg, 0x20) // not defined to lower the stack size
            // let arg3 := add(arg, 0x40) // not defined to lower the stack size
            // let locationDeltasBacklog := add(arg, 0x60) // not defined to lower the stack size
            let loacationDeltas := add(arg, 0x80)
            let loacationPrices := add(arg, 0xa0)

            let length
            let delta
            let price
            let tmpVar

            length := mul(sload(prices.slot), 8)
            mstore(loacationPrices, prices.slot)
            mstore(loacationPrices, keccak256(loacationPrices, 32))
            mstore(add(arg, 0x60), submittedDeltas.slot) // location of deltas backlog, not defined in a variable to lower the stack size
            mstore(add(arg, 0x60), keccak256(add(arg, 0x60), 32))
            mstore(add(arg, 0xc0), sload(submittedDeltas.slot)) // size of deltas backlog, not defined in a variable to lower the stack size
            mstore(add(arg, 0xe0), sload(currentDelta.slot)) // current position in the backlog, not defined in a variable to lower the stack size

            calldatacopy(arg, feeds.offset, 0x20)
            mstore(add(arg, 0x20), div(mload(arg), 8)) // slot
            mstore(add(arg, 0x40), mod(mload(arg), 8))  // position
            price := sload(add(mload(loacationPrices), mload(add(arg, 0x40))))

            for { let j := 0 } lt(j, feeds.length) { j := add(j, 1) } {
                calldatacopy(arg, add(feeds.offset, mul(j, 0x20)), 0x20)
                if iszero(lt(mload(arg), length)) {
                    revert(0, 0)
                }
                
                tmpVar := div(mload(arg), 8)  // use tmpVar for temporary value of slot
                mstore(add(arg, 0x40), mod(mload(arg), 8))  // position
                if iszero(eq(tmpVar, mload(add(arg, 0x20)))) {
                    mstore(add(arg, 0x20), tmpVar)
                    price := sload(add(mload(loacationPrices), tmpVar))
                }

                tmpVar := shl(mul(mload(add(arg, 0x40)), 32), price) // use tmpVar for temporary value extracting price
                mstore(add(add(_prices, 0x20), mul(j, 0x20)), shr(224, tmpVar))
            }

            // iterate over updates in the backlog
            for { let i := sload(backlogDelta.slot) } iszero(eq(i, mload(add(arg, 0xe0)))) { i := mod(add(i, 1),  mload(add(arg, 0xc0)))} {
                // get location of the i-th update in the updates backlog
                mstore(loacationDeltas, add(mload(add(arg, 0x60)), i))
                length := sload(mload(loacationDeltas))

                // data is stored differnetly if there is more or less than 31 bytes
                // more than 31 bytes
                if eq(mod(length, 2), 1) {
                    mstore(loacationDeltas, keccak256(loacationDeltas, 32))
                    length := div(length, 2)
                    // length := add(div(sub(length, 1), 32), 1)
                    // get indexes of the first feed
                    calldatacopy(arg, feeds.offset, 0x20)
                    mstore(add(arg, 0x20), div(mload(arg), 128))  // slot offset

                    // get delta that has information about the first feed, if it exists
                    if lt(mload(arg), mul(length, 4)) {
                        delta := sload(add(mload(loacationDeltas), mload(add(arg, 0x20))))
                    }

                    for { let j := 0 } lt(j, feeds.length) { j := add(j, 1) } {
                        calldatacopy(arg, add(feeds.offset, mul(j, 0x20)), 0x20)
                        tmpVar := div(mload(arg), 128) // use tmpVar for temporary value of slot

                        // if this update did not update the requested feed, skip it
                        if iszero(lt(mload(arg), mul(length, 4))) {
                            continue
                        }

                        mstore(add(arg, 0x40), mod(mload(arg), 128))  // position

                        if iszero(eq(tmpVar, mload(add(arg, 0x20)))) {
                            mstore(add(arg, 0x20), tmpVar)
                            delta := sload(add(mload(loacationDeltas), tmpVar))
                        }
                        
                        tmpVar := shl(mul(mload(add(arg, 0x40)), 2), delta) // use tmpVar for temporary value extracting one delta
                        tmpVar := shr(254, tmpVar)
                        if eq(tmpVar, 1) {
                            // mul
                            price := mul(mload(add(add(_prices, 0x20), mul(j, 0x20))), scale)
                            mstore(add(add(_prices, 0x20), mul(j, 0x20)), shr(15, price))
                        }
                        if eq(tmpVar, 3) {
                            // div
                            price := shl(15, mload(add(add(_prices, 0x20), mul(j, 0x20))))
                            mstore(add(add(_prices, 0x20), mul(j, 0x20)), div(price, scale))
                        }
                    }
                    length := 1 // to avoid executing the other if case
                }

                // less than 32 bytes
                if eq(mod(length, 2), 0) {
                    delta := length
                    length := div(mod(length, 64), 2)

                    // get indexes of the first feed
                    calldatacopy(arg, feeds.offset, 0x20)
                    mstore(add(arg, 0x20), div(mload(arg), 128))  // slot offset

                    for { let j := 0 } lt(j, feeds.length) { j := add(j, 1) } {
                        calldatacopy(arg, add(feeds.offset, mul(j, 0x20)), 0x20)

                        // if this update did not update the requested feed, skip it
                        if iszero(lt(mload(arg), mul(length, 4))) {
                            continue
                        }

                        mstore(add(arg, 0x40), mod(mload(arg), 128))  // position
                        
                        tmpVar := shl(mul(mload(add(arg, 0x40)), 2), delta) // use tmpVar for temporary value extracting one delta
                        tmpVar := shr(254, tmpVar)
                        if eq(tmpVar, 1) {
                            // mul
                            price := mul(mload(add(add(_prices, 0x20), mul(j, 0x20))), scale)
                            mstore(add(add(_prices, 0x20), mul(j, 0x20)), shr(15, price))
                        }
                        if eq(tmpVar, 3) {
                            // div
                            price := shl(15, mload(add(add(_prices, 0x20), mul(j, 0x20))))
                            mstore(add(add(_prices, 0x20), mul(j, 0x20)), div(price, scale))
                        }
                    }
                }
            }
        }
    }
}


