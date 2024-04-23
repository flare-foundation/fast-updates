package updates

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"fast-updates-client/contracts-interface/fast_updater"
	"fast-updates-client/sortition"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	uint256Ty, _ = abi.NewType("uint256", "uint256", nil)
	bytesTy, _   = abi.NewType("bytes", "bytes", nil)
)

// PrepareUpdates creates a struct that can be submitted to the FastUpdates contract.
func PrepareUpdates(updateProof *sortition.UpdateProof, deltas []byte, privateKey *ecdsa.PrivateKey) (*fast_updater.IFastUpdaterFastUpdates, error) {
	// prepare credential
	gamma := fast_updater.Bn256G1Point{
		X: updateProof.Proof.Gamma.X.BigInt(new(big.Int)),
		Y: updateProof.Proof.Gamma.Y.BigInt(new(big.Int)),
	}
	sortitionCredential := fast_updater.SortitionCredential{
		Replicate: updateProof.Replicate, Gamma: gamma, C: updateProof.Proof.C, S: updateProof.Proof.S,
	}

	// sign the update
	arguments := abi.Arguments{{Type: uint256Ty}, {Type: uint256Ty}, {Type: uint256Ty},
		{Type: uint256Ty}, {Type: uint256Ty}, {Type: uint256Ty}, {Type: bytesTy},
	}
	toHash, err := arguments.Pack(
		updateProof.BlockNumber, updateProof.Replicate, gamma.X,
		gamma.Y, updateProof.Proof.C, updateProof.Proof.S, deltas,
	)
	if err != nil {
		return nil, fmt.Errorf("PrepareUpdates: Pack: %w", err)
	}
	hashFunc := sha256.New()
	_, err = hashFunc.Write(toHash)
	if err != nil {
		return nil, fmt.Errorf("PrepareUpdates: Write: %w", err)
	}
	buf := hashFunc.Sum(nil)
	prefix := "\x19Ethereum Signed Message:\n32"
	hashed := crypto.Keccak256([]byte(prefix), buf)
	signature, err := crypto.Sign(hashed, privateKey)
	if err != nil {
		return nil, fmt.Errorf("PrepareUpdates: Sign: %w", err)
	}
	var r [32]byte
	copy(r[:], signature[0:32])
	var s [32]byte
	copy(s[:], signature[32:64])
	v := uint8(signature[64]) + 27
	sig := fast_updater.IFastUpdaterSignature{R: r, S: s, V: v}

	// prepare the update
	update := &fast_updater.IFastUpdaterFastUpdates{SortitionBlock: updateProof.BlockNumber, Deltas: deltas,
		SortitionCredential: sortitionCredential, Signature: sig}

	return update, nil
}

// PrepareUpdatesSubmission creates the bytes representation of the fast updates submission call
// stripped of the selector. This can be used to submit to the submission contract instead of directly
// to the Fast Updates contract.
func PrepareUpdatesSubmission(update *fast_updater.IFastUpdaterFastUpdates) ([]byte, error) {
	parsed, err := fast_updater.FastUpdaterMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("PrepareUpdatesSubmission: GetAbi: %w", err)
	}
	packedCall, err := parsed.Pack("submitUpdates", *update)
	if err != nil {
		return nil, fmt.Errorf("PrepareUpdatesSubmission: Pack: %w", err)
	}

	if len(packedCall) < 4 {
		return nil, fmt.Errorf("PrepareUpdatesSubmission: bytes representation of the call is too short")
	}

	return packedCall[4:], nil
}
