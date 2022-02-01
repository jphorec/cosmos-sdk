package server

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

<<<<<<< HEAD
// Deprecated: GenerateCoinKey generates a new key mnemonic along with its addrress.
// Please use testutils.GenerateCoinKey instead.
func GenerateCoinKey(algo keyring.SignatureAlgo) (sdk.AccAddress, string, error) {
	// generate a private key, with mnemonic
	info, secret, err := keyring.NewInMemory().NewMnemonic(
		"name",
		keyring.English,
		sdk.GetConfig().GetFullBIP44Path(),
		keyring.DefaultBIP39Passphrase,
		algo,
	)
=======
// GenerateCoinKey returns the address of a public key, along with the secret
// phrase to recover the private key.
func GenerateCoinKey(algo keyring.SignatureAlgo, cdc codec.Codec) (sdk.AccAddress, string, error) {
	// generate a private key, with recovery phrase
	k, secret, err := keyring.NewInMemory(cdc).NewMnemonic("name", keyring.English, sdk.GetConfig().GetFullBIP44Path(), keyring.DefaultBIP39Passphrase, algo)
>>>>>>> fred/allow_multiple_futures_for_sim
	if err != nil {
		return sdk.AccAddress{}, "", err
	}
<<<<<<< HEAD

	return sdk.AccAddress(info.GetPubKey().Address()), secret, nil
=======
	addr, err := k.GetAddress()
	if err != nil {
		return nil, "", err
	}
	return addr, secret, nil
>>>>>>> fred/allow_multiple_futures_for_sim
}

// Deprecated: GenerateSaveCoinKey generates a new key mnemonic with its addrress.
// If mnemonic is provided then it's used for key generation.
// The key is saved in the keyring. The function returns error if overwrite=true and the key
// already exists.
// Please use testutils.GenerateSaveCoinKey instead.
func GenerateSaveCoinKey(
	keybase keyring.Keyring,
	keyName string,
	overwrite bool,
	algo keyring.SignatureAlgo,
) (sdk.AccAddress, string, error) {
	exists := false
	_, err := keybase.Key(keyName)
	if err == nil {
		exists = true
	}

	// ensure no overwrite
	if !overwrite && exists {
		return sdk.AccAddress{}, "", fmt.Errorf("key already exists, overwrite is disabled")
	}

	// remove the old key by name if it exists
	if exists {
		if err := keybase.Delete(keyName); err != nil {
			return sdk.AccAddress{}, "", fmt.Errorf("failed to overwrite key")
		}
	}

<<<<<<< HEAD
	k, mnemonic, err := keybase.NewMnemonic(keyName, keyring.English, sdk.GetConfig().GetFullBIP44Path(), keyring.DefaultBIP39Passphrase, algo)
=======
	k, secret, err := keybase.NewMnemonic(keyName, keyring.English, sdk.GetConfig().GetFullBIP44Path(), keyring.DefaultBIP39Passphrase, algo)
>>>>>>> fred/allow_multiple_futures_for_sim
	if err != nil {
		return sdk.AccAddress{}, "", err
	}

<<<<<<< HEAD
	return sdk.AccAddress(k.GetAddress()), mnemonic, nil
=======
	addr, err := k.GetAddress()
	if err != nil {
		return nil, "", err
	}

	return addr, secret, nil
>>>>>>> fred/allow_multiple_futures_for_sim
}
