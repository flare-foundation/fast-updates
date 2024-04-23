package sortition_test

import (
	"crypto/rand"
	"fast-updates-client/sortition"
	"fmt"
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc/bn254"
)

func TestSortition(t *testing.T) {
	key, err := sortition.KeyGen()
	if err != nil {
		t.Fatal(err)
	}

	blockNum, err := rand.Int(rand.Reader, new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil))
	if err != nil {
		t.Fatal(err)
	}
	seed, err := rand.Int(rand.Reader, new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil))
	if err != nil {
		t.Fatal(err)
	}
	for i := int64(0); i < 1000; i++ {
		replicate := big.NewInt(i)
		proof, err := sortition.VerifiableRandomness(key, seed, blockNum, replicate)
		if err != nil {
			t.Fatal(err)
		}

		check, err := sortition.VerifyRandomness(proof, key.Pk, seed, blockNum, replicate)
		if err != nil {
			t.Fatal(err)
		}
		if check == false {
			t.Fatal(fmt.Errorf("failed randomness check"))
		}
	}
}

func TestFindUpdateProofs(t *testing.T) {
	key, err := sortition.KeyGen()
	if err != nil {
		t.Fatal(err)
	}

	blockNum, err := rand.Int(rand.Reader, new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil))
	if err != nil {
		t.Fatal(err)
	}
	seed, err := rand.Int(rand.Reader, new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil))
	if err != nil {
		t.Fatal(err)
	}
	cutoff := new(big.Int).Div(bn254.ID.BaseField(), big.NewInt(10))
	weight := uint64(1000)

	updateProofs, err := sortition.FindUpdateProofs(key, seed, cutoff, blockNum, weight)
	if err != nil {
		t.Fatal(err)
	}
	for _, updateProof := range updateProofs {
		if updateProof.Proof.Gamma.X.BigInt(new(big.Int)).Cmp(cutoff) >= 0 {
			t.Fatal("randomness should be smaller than cutoff")
		}
		if updateProof.Replicate.Int64() >= int64(weight) {
			t.Fatal("replicate should be smaller than weight")
		}
		check, err := sortition.VerifyRandomness(*updateProof.Proof, key.Pk, seed, blockNum, updateProof.Replicate)
		if err != nil {
			t.Fatal(err)
		}
		if check == false {
			t.Fatal(fmt.Errorf("failed randomness check"))
		}
	}
}
