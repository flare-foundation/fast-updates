// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

// it is just a stub, not a live deployment;
// we are fine with experimental feature
/* solium-disable-next-line */
pragma experimental ABIEncoderV2;

import "../lib/Bn256.sol";

contract TestBn256 {
    Bn256.G1Point g1 = Bn256.g1();

    function runHashingTest() public view {
        string memory hello = "hello!";
        string memory goodbye = "goodbye.";
        Bn256.G1Point memory p_1;
        Bn256.G1Point memory p_2;
        p_1 = Bn256.g1HashToPoint(bytes(hello));
        p_2 = Bn256.g1HashToPoint(bytes(goodbye));

        require(p_1.x != 0, "X should not equal 0 in a hashed point.");
        require(p_1.y != 0, "Y should not equal 0 in a hashed point.");
        require(p_2.x != 0, "X should not equal 0 in a hashed point.");
        require(p_2.y != 0, "Y should not equal 0 in a hashed point.");

        require(Bn256.isG1PointOnCurve(p_1), "Hashed points should be on the curve.");
        require(Bn256.isG1PointOnCurve(p_2), "Hashed points should be on the curve.");
    }

    function runHashAndAddTest() public view {
        string memory hello = "hello!";
        string memory goodbye = "goodbye.";
        Bn256.G1Point memory p_1;
        Bn256.G1Point memory p_2;
        p_1 = Bn256.g1HashToPoint(bytes(hello));
        p_2 = Bn256.g1HashToPoint(bytes(goodbye));

        Bn256.G1Point memory p_3;
        Bn256.G1Point memory p_4;

        p_3 = Bn256.g1Add(p_1, p_2);
        p_4 = Bn256.g1Add(p_2, p_1);

        require(p_3.x == p_4.x, "Point addition should be commutative.");
        require(p_3.y == p_4.y, "Point addition should be commutative.");

        require(Bn256.isG1PointOnCurve(p_3), "Added points should be on the curve.");
    }

    function runHashAndScalarMultiplyTest() public view {
        string memory hello = "hello!";
        Bn256.G1Point memory p_1;
        Bn256.G1Point memory p_2;
        p_1 = Bn256.g1HashToPoint(bytes(hello));

        p_2 = Bn256.scalarMultiply(p_1, 12);

        require(Bn256.isG1PointOnCurve(p_2), "Multiplied point should be on the curve.");
    }

    function publicG1Add(Bn256.G1Point memory a, Bn256.G1Point memory b) public view returns (Bn256.G1Point memory c) {
        c = Bn256.g1Add(a, b);
    }

    function publicG1ScalarMultiply(Bn256.G1Point memory a, uint256 s) public view returns (Bn256.G1Point memory c) {
        c = Bn256.scalarMultiply(a, s);
    }
}
