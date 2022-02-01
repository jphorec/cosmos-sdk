package testutil

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
<<<<<<< HEAD:testutil/key_test.go
=======
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/simapp"
>>>>>>> fred/allow_multiple_futures_for_sim:server/init_test.go
	"github.com/cosmos/cosmos-sdk/types"
)

func TestGenerateCoinKey(t *testing.T) {
	t.Parallel()
<<<<<<< HEAD:testutil/key_test.go
	addr, mnemonic, err := GenerateCoinKey(hd.Secp256k1)
=======
	cdc := simapp.MakeTestEncodingConfig().Codec
	addr, mnemonic, err := server.GenerateCoinKey(hd.Secp256k1, cdc)
>>>>>>> fred/allow_multiple_futures_for_sim:server/init_test.go
	require.NoError(t, err)

	// Test creation
	k, err := keyring.NewInMemory(cdc).NewAccount("xxx", mnemonic, "", hd.NewFundraiserParams(0, types.GetConfig().GetCoinType(), 0).String(), hd.Secp256k1)
	require.NoError(t, err)
	addr1, err := k.GetAddress()
	require.NoError(t, err)
	require.Equal(t, addr, addr1)
}

func TestGenerateSaveCoinKey(t *testing.T) {
	t.Parallel()

	encCfg := simapp.MakeTestEncodingConfig()
	kb, err := keyring.New(t.Name(), "test", t.TempDir(), nil, encCfg.Codec)
	require.NoError(t, err)

	addr, mnemonic, err := GenerateSaveCoinKey(kb, "keyname", "", false, hd.Secp256k1)
	require.NoError(t, err)

	// Test key was actually saved
	k, err := kb.Key("keyname")
	require.NoError(t, err)
	addr1, err := k.GetAddress()
	require.NoError(t, err)
	require.Equal(t, addr, addr1)

	// Test in-memory recovery
	k, err = keyring.NewInMemory(encCfg.Codec).NewAccount("xxx", mnemonic, "", hd.NewFundraiserParams(0, types.GetConfig().GetCoinType(), 0).String(), hd.Secp256k1)
	require.NoError(t, err)
	addr1, err = k.GetAddress()
	require.NoError(t, err)
	require.Equal(t, addr, addr1)
}

func TestGenerateSaveCoinKeyOverwriteFlag(t *testing.T) {
	t.Parallel()

	encCfg := simapp.MakeTestEncodingConfig()
	kb, err := keyring.New(t.Name(), "test", t.TempDir(), nil, encCfg.Codec)
	require.NoError(t, err)

	keyname := "justakey"
	addr1, _, err := GenerateSaveCoinKey(kb, keyname, "", false, hd.Secp256k1)
	require.NoError(t, err)

	// Test overwrite with overwrite=false
	_, _, err = GenerateSaveCoinKey(kb, keyname, "", false, hd.Secp256k1)
	require.Error(t, err)

	// Test overwrite with overwrite=true
	addr2, _, err := GenerateSaveCoinKey(kb, keyname, "", true, hd.Secp256k1)
	require.NoError(t, err)

	require.NotEqual(t, addr1, addr2)
}
