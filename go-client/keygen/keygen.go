package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fast-updates-client/config"
	"fast-updates-client/logger"
	"fast-updates-client/sortition"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fp"
)

var InFlag = flag.String("key", "", "Private key")
var AddressFlag = flag.String("address", "", "Value of the address that needs to be signed.")
var InFileFlag = flag.String("key_file", "", "File to load private and public key")
var KeyOutFlag = flag.String("key_out", "", "File to save a freshly generated private and public key")
var SigOutFlag = flag.String("sig_out", "", "File to save a signature")

type keyStrings struct {
	PublicKeyX string
	PublicKeyY string
	PrivateKey string
}

type sigStrings struct {
	RX string
	RY string
	S  string
}

func main() {
	flag.Parse()
	cfgLog := config.LoggerConfig{Console: true}
	cfg := &config.Config{Logger: cfgLog}

	config.GlobalConfigCallback.Call(cfg)

	var keys *sortition.Key
	var err error
	if *InFileFlag == "" && *InFlag == "" {
		logger.Info("No input specified, generating a new key pair.")
		keys, err = sortition.KeyGen()
		if err != nil {
			log.Fatal(err)
		}
		keyStrings := keyStrings{PrivateKey: "0x" + keys.Sk.Text(16), PublicKeyX: "0x" + keys.Pk.X.Text(16), PublicKeyY: "0x" + keys.Pk.Y.Text(16)}
		keyBytes, err := json.Marshal(keyStrings)
		if err != nil {
			log.Fatal(err)
		}

		if *KeyOutFlag == "" {
			logger.Info("Key generated: " + string(keyBytes))
		} else {
			f, err := os.Create(*KeyOutFlag)
			if err != nil {
				log.Fatal(err)
			}

			_, err = f.Write(keyBytes)
			if err != nil {
				log.Fatal(err)
			}
			err = f.Close()
			if err != nil {
				log.Fatal(err)
			}
			logger.Info("Saved key in file " + *KeyOutFlag)
		}

	} else {
		if *InFileFlag != "" {
			keyBytes, err := os.ReadFile(*InFileFlag)
			if err != nil {
				log.Fatal(err)
			}
			var keyStrings keyStrings
			err = json.Unmarshal(keyBytes, &keyStrings)
			if err != nil {
				log.Fatal(err)
			}

			keys, err = sortition.KeyFromString(keyStrings.PrivateKey)
			if err != nil {
				log.Fatal(err)
			}

			pkCheck := &bn254.G1Affine{}
			pkXCheck, check := new(big.Int).SetString(keyStrings.PublicKeyX, 0)
			if !check {
				log.Fatal(fmt.Errorf("failed to read the key"))
			}
			pkCheck.X = *new(fp.Element).SetBigInt(pkXCheck)
			pkYCheck, check := new(big.Int).SetString(keyStrings.PublicKeyY, 0)
			if !check {
				log.Fatal(fmt.Errorf("failed to read the key"))
			}
			pkCheck.Y = *new(fp.Element).SetBigInt(pkYCheck)

			if !keys.Pk.Equal(pkCheck) {
				log.Fatal(fmt.Errorf("keys deformed"))
			}

			logger.Info("Read the key pair from " + *InFileFlag)
		} else {
			keys, err = sortition.KeyFromString(*InFlag)
			if err != nil {
				log.Fatal(err)
			}

			logger.Info("Read the key from the provided flag")
		}

	}

	if *AddressFlag != "" {
		addressBytes, err := hex.DecodeString((*AddressFlag)[2:])
		if err != nil {
			log.Fatal(err)
		}
		hash := sha256.New()
		hash.Write(addressBytes)
		buf := hash.Sum(nil)

		var msg [32]byte
		copy(msg[:], buf)
		signature, err := sortition.Sign(keys, msg)
		if err != nil {
			log.Fatal(err)
		}
		sigStrings := sigStrings{S: "0x" + signature.S.Text(16), RX: "0x" + signature.R.X.Text(16), RY: "0x" + signature.R.Y.Text(16)}
		sigBytes, err := json.Marshal(sigStrings)
		if err != nil {
			log.Fatal(err)
		}

		if *SigOutFlag == "" {
			logger.Info("Signature generated: " + string(sigBytes))
		} else {
			f, err := os.Create(*SigOutFlag)
			if err != nil {
				log.Fatal(err)
			}

			_, err = f.Write(sigBytes)
			if err != nil {
				log.Fatal(err)
			}
			err = f.Close()
			if err != nil {
				log.Fatal(err)
			}
			logger.Info("Saved the signature in file " + *SigOutFlag)
		}
	}

}
