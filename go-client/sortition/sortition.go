package sortition

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	p            = bn254.ID.BaseField()
	q            = bn254.ID.ScalarField()
	g            = GeneratorG1()
	one          = big.NewInt(1)
	uint256Ty, _ = abi.NewType("uint256", "uint256", nil)
	bytes32Ty, _ = abi.NewType("bytes32", "bytes32", nil)
)

// Key structure represents a private and public key
// used in the sortition protocol
type Key struct {
	Pk *bn254.G1Affine
	Sk *big.Int
}

// Proof represents a generated verifiable randomness
// together with a proof of the correctness.
type Proof struct {
	Gamma *bn254.G1Affine
	C     *big.Int
	S     *big.Int
}

// UpdateProof contains all the information needed to prove the eligibility to
// submit updates.
type UpdateProof struct {
	Proof       *Proof
	BlockNumber *big.Int
	Replicate   *big.Int
}

// Signature represents a Schnorr signature of arbitrary message using
// the private/public key used in the sortition protocol.
type Signature struct {
	R *bn254.G1Affine
	S *big.Int
}

// KeyGen generates a private/public key pair to be used in the
// sortition protocol.
func KeyGen() (*Key, error) {
	sk, err := rand.Int(rand.Reader, q)
	if err != nil {
		return nil, err
	}
	pk := new(bn254.G1Affine).ScalarMultiplicationBase(sk)

	return &Key{pk, sk}, nil
}

// KeyFromString creates a a private/public key pair from the
// hexagonal string representation of the private key.
func KeyFromString(skString string) (*Key, error) {
	sk, check := new(big.Int).SetString(skString, 0)
	if !check {
		return nil, fmt.Errorf("failed to read the secret key")
	}
	pk := new(bn254.G1Affine).ScalarMultiplicationBase(sk)

	return &Key{pk, sk}, nil
}

// VerifiableRandomness creates a deterministic verifiable randomness (with proof) given a
// private key, seed, blockNum, and replicate number.
func VerifiableRandomness(key *Key, seed *big.Int, blockNum *big.Int, replicate *big.Int) (Proof, error) {
	arguments := abi.Arguments{{Type: uint256Ty}, {Type: uint256Ty}, {Type: uint256Ty}}
	toHash, err := arguments.Pack(seed, blockNum, replicate)
	if err != nil {
		return Proof{}, err
	}

	h := HashToG1(toHash)
	gamma := new(bn254.G1Affine).ScalarMultiplication(h, key.Sk)

	k, err := rand.Int(rand.Reader, q)
	if err != nil {
		return Proof{}, err
	}

	gToK := new(bn254.G1Affine).ScalarMultiplicationBase(k)
	hToK := new(bn254.G1Affine).ScalarMultiplication(h, k)

	arguments = abi.Arguments{{Type: uint256Ty}, {Type: uint256Ty}, {Type: uint256Ty},
		{Type: uint256Ty}, {Type: uint256Ty}, {Type: uint256Ty},
		{Type: uint256Ty}, {Type: uint256Ty}, {Type: uint256Ty},
		{Type: uint256Ty}, {Type: uint256Ty}, {Type: uint256Ty}}

	toHash, err = arguments.Pack(
		g.X.BigInt(new(big.Int)), g.Y.BigInt(new(big.Int)),
		h.X.BigInt(new(big.Int)), h.Y.BigInt(new(big.Int)),
		key.Pk.X.BigInt(new(big.Int)), key.Pk.Y.BigInt(new(big.Int)),
		gamma.X.BigInt(new(big.Int)), gamma.Y.BigInt(new(big.Int)),
		gToK.X.BigInt(new(big.Int)), gToK.Y.BigInt(new(big.Int)),
		hToK.X.BigInt(new(big.Int)), hToK.Y.BigInt(new(big.Int)),
	)
	if err != nil {
		return Proof{}, err
	}
	hash := sha256.New()
	hash.Write(toHash)
	buf := hash.Sum(nil)
	c := new(big.Int).SetBytes(buf)
	c.Mod(c, q)

	s := new(big.Int).Mul(key.Sk, c)
	s.Neg(s)
	s.Add(k, s)
	s.Mod(s, q)

	return Proof{Gamma: gamma, S: s, C: c}, nil
}

// VerifyRandomness verifies that the provided randomness corresponds to the providers public key,
// and given seed, block number, and replicate number. Used for tests, actual verification
// is done at the contract.
func VerifyRandomness(proof Proof, pk *bn254.G1Affine, seed *big.Int, blockNum *big.Int, replicate *big.Int) (bool, error) {
	pkToC := new(bn254.G1Affine).ScalarMultiplication(pk, proof.C)
	gToS := new(bn254.G1Affine).ScalarMultiplicationBase(proof.S)
	u := new(bn254.G1Affine).Add(pkToC, gToS)

	arguments := abi.Arguments{{Type: uint256Ty}, {Type: uint256Ty}, {Type: uint256Ty}}
	toHash, err := arguments.Pack(seed, blockNum, replicate)
	if err != nil {
		return false, err
	}
	h := HashToG1(toHash)

	gammaToC := new(bn254.G1Affine).ScalarMultiplication(proof.Gamma, proof.C)
	hToS := new(bn254.G1Affine).ScalarMultiplication(h, proof.S)
	v := new(bn254.G1Affine).Add(gammaToC, hToS)

	arguments = abi.Arguments{{Type: uint256Ty}, {Type: uint256Ty}, {Type: uint256Ty},
		{Type: uint256Ty}, {Type: uint256Ty}, {Type: uint256Ty},
		{Type: uint256Ty}, {Type: uint256Ty}, {Type: uint256Ty},
		{Type: uint256Ty}, {Type: uint256Ty}, {Type: uint256Ty}}

	toHash, err = arguments.Pack(
		g.X.BigInt(new(big.Int)), g.Y.BigInt(new(big.Int)),
		h.X.BigInt(new(big.Int)), h.Y.BigInt(new(big.Int)),
		pk.X.BigInt(new(big.Int)), pk.Y.BigInt(new(big.Int)),
		proof.Gamma.X.BigInt(new(big.Int)), proof.Gamma.Y.BigInt(new(big.Int)),
		u.X.BigInt(new(big.Int)), u.Y.BigInt(new(big.Int)),
		v.X.BigInt(new(big.Int)), v.Y.BigInt(new(big.Int)),
	)
	if err != nil {
		return false, err
	}
	var buf []byte
	hash := sha256.New()
	hash.Write(toHash)
	buf = hash.Sum(nil)
	c := new(big.Int).SetBytes(buf)
	c.Mod(c, q)

	if c.Cmp(proof.C) != 0 {
		return false, nil
	}

	return true, nil
}

// FindUpdateProofs searches among replicate numbers smaller than the provided weight for
// a randomness that is generated by the provided key, seed and block number, and is
// smaller than the provided cutoff. Such a randomness proves that the client can submit
// an update to the chain.
func FindUpdateProofs(key *Key, seed, cutoff *big.Int, blockNum *big.Int, weight uint64) ([]*UpdateProof, error) {
	updateProofs := make([]*UpdateProof, 0)
	for rep := 0; rep < int(weight); rep++ {
		proof, err := VerifiableRandomness(key, seed, blockNum, big.NewInt(int64(rep)))
		if err != nil {
			return nil, fmt.Errorf("VerifiableRandomness: %w", err)
		}
		if proof.Gamma.X.BigInt(new(big.Int)).Cmp(cutoff) < 0 {
			updateProof := UpdateProof{Proof: &proof, BlockNumber: blockNum, Replicate: big.NewInt(int64(rep))}
			updateProofs = append(updateProofs, &updateProof)
		}
	}

	return updateProofs, nil
}

// HashToG1 hashes an arbitrary message to a point in an elliptic group.
func HashToG1(msg []byte) *bn254.G1Affine {
	var buf []byte
	hash := sha256.New()
	hash.Write(msg)
	buf = hash.Sum(nil)
	x := new(big.Int).SetBytes(buf)

	for {
		x3 := new(big.Int).Exp(x, big.NewInt(3), p)
		x3.Add(x3, big.NewInt(3))

		y := new(big.Int).ModSqrt(x3, p)
		if y != nil {
			g := new(bn254.G1Affine)
			g.X.SetBigInt(x)
			g.Y.SetBigInt(y)

			return g
		}
		x.Add(x, one)
	}
}

func Sign(key *Key, msg [32]byte) (*Signature, error) {
	k, err := rand.Int(rand.Reader, q)
	if err != nil {
		return nil, err
	}
	r := new(bn254.G1Affine).ScalarMultiplicationBase(k)

	arguments := abi.Arguments{
		{Type: uint256Ty}, {Type: uint256Ty}, {Type: bytes32Ty},
		{Type: uint256Ty}, {Type: uint256Ty},
	}

	toHash, err := arguments.Pack(
		key.Pk.X.BigInt(new(big.Int)), key.Pk.Y.BigInt(new(big.Int)),
		msg,
		r.X.BigInt(new(big.Int)), r.Y.BigInt(new(big.Int)),
	)
	if err != nil {
		return nil, err
	}

	eBytes := crypto.Keccak256(toHash)
	e := new(big.Int).SetBytes(eBytes)

	s := new(big.Int).Sub(k, new(big.Int).Mul(key.Sk, e))
	s.Mod(s, q)

	return &Signature{R: r, S: s}, nil
}

// GeneratorG1 returns the generator of the elliptic group used in
// the sortition protocol.
func GeneratorG1() bn254.G1Affine {
	_, _, g, _ := bn254.Generators()

	return g
}
