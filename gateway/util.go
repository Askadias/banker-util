package gateway

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

func NewMnemonic(length int) (string, error) {
	entropy := make([]byte, length)
	_, err := rand.Read(entropy)
	if err != nil {
		return "", err
	}
	return bip39.NewMnemonic(entropy)
}

func hdWallet(mnemonic string) (*ecdsa.PrivateKey, error) {
	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed(mnemonic, "")
	// Generate a new master node using the seed.
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}

	// This gives the path: m/44H
	acc44H, err := masterKey.Child(hdkeychain.HardenedKeyStart + 44)
	if err != nil {
		return nil, err
	}

	// This gives the path: m/44H/60H
	acc44H60H, err := acc44H.Child(hdkeychain.HardenedKeyStart + 60)
	if err != nil {
		return nil, err
	}

	// This gives the path: m/44H/60H/0H
	acc44H60H0H, err := acc44H60H.Child(hdkeychain.HardenedKeyStart + 0)
	if err != nil {
		return nil, err
	}

	// This gives the path: m/44H/60H/0H/0
	acc44H60H0H0, err := acc44H60H0H.Child(0)
	if err != nil {
		return nil, err
	}

	// This gives the path: m/44H/60H/0H/0/0
	acc44H60H0H00, err := acc44H60H0H0.Child(0)
	if err != nil {
		return nil, err
	}

	btcecPrivKey, err := acc44H60H0H00.ECPrivKey()
	if err != nil {
		return nil, err
	}

	privateKey := btcecPrivKey.ToECDSA()

	return privateKey, nil
}

// Get private key from seed.
func PrivateKeyBySeed(seed []byte) (string, error) {
	// Generate a new master node using the seed.
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return "", err
	}

	// This gives the path: m/44H
	acc44H, err := masterKey.NewChildKey(bip32.FirstHardenedChild + 44)
	if err != nil {
		return "", err
	}

	// This gives the path: m/44H/60H
	acc44H60H, err := acc44H.NewChildKey(bip32.FirstHardenedChild + 60)
	if err != nil {
		return "", err
	}

	// This gives the path: m/44H/60H/0H
	acc44H60H0H, err := acc44H60H.NewChildKey(bip32.FirstHardenedChild + 0)
	if err != nil {
		return "", err
	}

	// This gives the path: m/44H/60H/0H/0
	acc44H60H0H0, err := acc44H60H0H.NewChildKey(0)
	if err != nil {
		return "", err
	}

	// This gives the path: m/44H/60H/0H/0/0
	acc44H60H0H00, err := acc44H60H0H0.NewChildKey(0)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(acc44H60H0H00.Key), nil
}