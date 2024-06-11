package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
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
	"github.com/ethereum/go-ethereum/accounts/abi"
	"golang.org/x/crypto/scrypt"
)

var InFlag = flag.String("key", "", "Private key")
var AddressFlag = flag.String("address", "", "Value of the address that needs to be signed.")
var InFileFlag = flag.String("key_file", "", "File to load private and public key")
var KeyOutFlag = flag.String("key_out", "", "File to save a freshly generated private and public key")
var PassFlag = flag.String("pass", "", "Password for encrypting/decrypting private key")
var SigOutFlag = flag.String("sig_out", "", "File to save a signature")
var uint256Ty, _ = abi.NewType("uint256", "uint256", nil)

type keyStrings struct {
	PublicKeyX string
	PublicKeyY string
	PrivateKey string
}

type keyEncrypted struct {
	PublicKeyX          string
	PublicKeyY          string
	EncryptedPrivateKey []byte
}

type sigStrings struct {
	Signature string
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

		if *KeyOutFlag == "" {
			keysString := keyStrings{PrivateKey: "0x" + keys.Sk.Text(16), PublicKeyX: "0x" + keys.Pk.X.Text(16), PublicKeyY: "0x" + keys.Pk.Y.Text(16)}
			keyBytes, err := json.Marshal(keysString)
			if err != nil {
				log.Fatal(err)
			}
			logger.Info("Key generated: " + string(keyBytes))
		} else {
			if *PassFlag == "" {
				log.Fatal("password should be specified to encrypt the private key")
			}

			encryptedPrivateKey, err := Encrypt([]byte(*PassFlag), keys.Sk.Bytes())
			if err != nil {
				log.Fatal(err)
			}

			keyEncrypted := keyEncrypted{EncryptedPrivateKey: encryptedPrivateKey, PublicKeyX: "0x" + keys.Pk.X.Text(16), PublicKeyY: "0x" + keys.Pk.Y.Text(16)}
			keyBytes, err := json.Marshal(keyEncrypted)
			if err != nil {
				log.Fatal(err)
			}

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
			if *PassFlag == "" {
				log.Fatal("password should be specified to decrypt the private key")
			}

			keyBytes, err := os.ReadFile(*InFileFlag)
			if err != nil {
				log.Fatal(err)
			}
			var keyEncrypted keyEncrypted
			err = json.Unmarshal(keyBytes, &keyEncrypted)
			if err != nil {
				log.Fatal(err)
			}

			privateKeyBytes, err := Decrypt([]byte(*PassFlag), keyEncrypted.EncryptedPrivateKey)
			if err != nil {
				log.Fatal(err)
			}

			keys, err = sortition.KeyFromString("0x" + new(big.Int).SetBytes(privateKeyBytes).Text(16))
			if err != nil {
				log.Fatal(err)
			}

			pkCheck := &bn254.G1Affine{}
			pkXCheck, check := new(big.Int).SetString(keyEncrypted.PublicKeyX, 0)
			if !check {
				log.Fatal(fmt.Errorf("failed to read the key"))
			}
			pkCheck.X = *new(fp.Element).SetBigInt(pkXCheck)
			pkYCheck, check := new(big.Int).SetString(keyEncrypted.PublicKeyY, 0)
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

		arguments := abi.Arguments{
			{Type: uint256Ty}, {Type: uint256Ty}, {Type: uint256Ty},
		}
		sigBytes, err := arguments.Pack(
			signature.S,
			signature.R.X.BigInt(new(big.Int)),
			signature.R.Y.BigInt(new(big.Int)),
		)
		if err != nil {
			log.Fatal(err)
		}

		if *SigOutFlag == "" {
			logger.Info("Signature generated: 0x" + hex.EncodeToString(sigBytes))
		} else {
			f, err := os.Create(*SigOutFlag)
			if err != nil {
				log.Fatal(err)
			}

			sigStruct := sigStrings{Signature: "0x" + hex.EncodeToString(sigBytes)}
			sigToWrite, err := json.Marshal(sigStruct)
			if err != nil {
				log.Fatal(err)
			}
			_, err = f.Write(sigToWrite)
			if err != nil {
				log.Fatal(err)
			}
			err = f.Close()
			if err != nil {
				log.Fatal(err)
			}
			logger.Info("Saved the signature in file " + *SigOutFlag)
		}
	} else {
		// in case that key was read and decrypted from a file, but not used for the signature, just print it
		if *InFileFlag != "" {
			keysString := keyStrings{PrivateKey: "0x" + keys.Sk.Text(16), PublicKeyX: "0x" + keys.Pk.X.Text(16), PublicKeyY: "0x" + keys.Pk.Y.Text(16)}
			keyBytes, err := json.Marshal(keysString)
			if err != nil {
				log.Fatal(err)
			}
			logger.Info("Key: " + string(keyBytes))
		}
	}
}

func Encrypt(password, data []byte) ([]byte, error) {
	key, salt, err := DeriveKey(password, nil)
	if err != nil {
		return nil, err
	}

	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	ciphertext = append(ciphertext, salt...)

	return ciphertext, nil
}

func Decrypt(password, data []byte) ([]byte, error) {
	salt, data := data[len(data)-32:], data[:len(data)-32]

	key, _, err := DeriveKey(password, salt)
	if err != nil {
		return nil, err
	}

	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}

	nonce, ciphertext := data[:gcm.NonceSize()], data[gcm.NonceSize():]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func DeriveKey(password, salt []byte) ([]byte, []byte, error) {
	if salt == nil {
		salt = make([]byte, 32)
		if _, err := rand.Read(salt); err != nil {
			return nil, nil, err
		}
	}

	key, err := scrypt.Key(password, salt, 1048576, 8, 1, 32)
	if err != nil {
		return nil, nil, err
	}

	return key, salt, nil
}
