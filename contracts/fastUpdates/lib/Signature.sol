pragma solidity 0.8.18;


struct Signature {
    uint8 v;
    bytes32 r;
    bytes32 s;
}

function recoverSigner(bytes32 _hash, Signature memory _signature) pure returns (address) {
    bytes memory prefix = "\x19Ethereum Signed Message:\n32";
    bytes32 prefixedHashMessage = keccak256(abi.encodePacked(prefix, _hash));
    address signer = ecrecover(prefixedHashMessage, _signature.v, _signature.r, _signature.s);
    require(signer != address(0), "ECDSA: invalid signature");
    return signer;
}
