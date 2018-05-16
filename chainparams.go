package main

import (
	"github.com/machinecoin-project/lnd/keychain"
	litecoinCfg "github.com/ltcsuite/ltcd/chaincfg"
	litecoinWire "github.com/ltcsuite/ltcd/wire"
	machinecoinCfg "github.com/macsuite/macd/chaincfg"
	machinecoinWire "github.com/macsuite/macd/wire"
	"github.com/roasbeef/btcd/chaincfg"
	bitcoinCfg "github.com/roasbeef/btcd/chaincfg"
	"github.com/roasbeef/btcd/chaincfg/chainhash"
	bitcoinWire "github.com/roasbeef/btcd/wire"
)

// activeNetParams is a pointer to the parameters specific to the currently
// active bitcoin network.
var activeNetParams = bitcoinTestNetParams

// bitcoinNetParams couples the p2p parameters of a network with the
// corresponding RPC port of a daemon running on the particular network.
type bitcoinNetParams struct {
	*bitcoinCfg.Params
	rpcPort  string
	CoinType uint32
}

// litecoinNetParams couples the p2p parameters of a network with the
// corresponding RPC port of a daemon running on the particular network.
type litecoinNetParams struct {
	*litecoinCfg.Params
	rpcPort  string
	CoinType uint32
}

// machinecoinNetParams couples the p2p parameters of a network with the
// corresponding RPC port of a daemon running on the particular network.
type machinecoinNetParams struct {
	*machinecoinCfg.Params
	rpcPort  string
	CoinType uint32
}

// bitcoinTestNetParams contains parameters specific to the 3rd version of the
// test network.
var bitcoinTestNetParams = bitcoinNetParams{
	Params:   &bitcoinCfg.TestNet3Params,
	rpcPort:  "18334",
	CoinType: keychain.CoinTypeTestnet,
}

// bitcoinMainNetParams contains parameters specific to the current Bitcoin
// mainnet.
var bitcoinMainNetParams = bitcoinNetParams{
	Params:   &bitcoinCfg.MainNetParams,
	rpcPort:  "8334",
	CoinType: keychain.CoinTypeBitcoin,
}

// bitcoinSimNetParams contains parameters specific to the simulation test
// network.
var bitcoinSimNetParams = bitcoinNetParams{
	Params:   &bitcoinCfg.SimNetParams,
	rpcPort:  "18556",
	CoinType: keychain.CoinTypeTestnet,
}

// litecoinTestNetParams contains parameters specific to the 4th version of the
// test network.
var litecoinTestNetParams = litecoinNetParams{
	Params:   &litecoinCfg.TestNet4Params,
	rpcPort:  "19334",
	CoinType: keychain.CoinTypeTestnet,
}

// litecoinMainNetParams contains the parameters specific to the current
// Litecoin mainnet.
var litecoinMainNetParams = litecoinNetParams{
	Params:   &litecoinCfg.MainNetParams,
	rpcPort:  "9334",
	CoinType: keychain.CoinTypeLitecoin,
}

// machinecoinTestNetParams contains parameters specific to the 4th version of the
// test network.
var machinecoinTestNetParams = machinecoinNetParams{
	Params:   &machinecoinCfg.TestNet4Params,
	rpcPort:  "50332",
	CoinType: keychain.CoinTypeTestnet,
}

// machinecoinMainNetParams contains the parameters specific to the current
// Machinecoin mainnet.
var machinecoinMainNetParams = machinecoinNetParams{
	Params:   &machinecoinCfg.MainNetParams,
	rpcPort:  "40332",
	CoinType: keychain.CoinTypeMachinecoin,
}

// regTestNetParams contains parameters specific to a local regtest network.
var regTestNetParams = bitcoinNetParams{
	Params:   &bitcoinCfg.RegressionNetParams,
	rpcPort:  "18334",
	CoinType: keychain.CoinTypeTestnet,
}

// applyLitecoinParams applies the relevant chain configuration parameters that
// differ for litecoin to the chain parameters typed for btcsuite derivation.
// This function is used in place of using something like interface{} to
// abstract over _which_ chain (or fork) the parameters are for.
func applyLitecoinParams(params *bitcoinNetParams, litecoinParams *litecoinNetParams) {
	params.Name = litecoinParams.Name
	params.Net = bitcoinWire.BitcoinNet(litecoinParams.Net)
	params.DefaultPort = litecoinParams.DefaultPort
	params.CoinbaseMaturity = litecoinParams.CoinbaseMaturity

	copy(params.GenesisHash[:], litecoinParams.GenesisHash[:])

	// Address encoding magics
	params.PubKeyHashAddrID = litecoinParams.PubKeyHashAddrID
	params.ScriptHashAddrID = litecoinParams.ScriptHashAddrID
	params.PrivateKeyID = litecoinParams.PrivateKeyID
	params.WitnessPubKeyHashAddrID = litecoinParams.WitnessPubKeyHashAddrID
	params.WitnessScriptHashAddrID = litecoinParams.WitnessScriptHashAddrID
	params.Bech32HRPSegwit = litecoinParams.Bech32HRPSegwit

	copy(params.HDPrivateKeyID[:], litecoinParams.HDPrivateKeyID[:])
	copy(params.HDPublicKeyID[:], litecoinParams.HDPublicKeyID[:])

	params.HDCoinType = litecoinParams.HDCoinType

	checkPoints := make([]chaincfg.Checkpoint, len(litecoinParams.Checkpoints))
	for i := 0; i < len(litecoinParams.Checkpoints); i++ {
		var chainHash chainhash.Hash
		copy(chainHash[:], litecoinParams.Checkpoints[i].Hash[:])

		checkPoints[i] = chaincfg.Checkpoint{
			Height: litecoinParams.Checkpoints[i].Height,
			Hash:   &chainHash,
		}
	}
	params.Checkpoints = checkPoints

	params.rpcPort = litecoinParams.rpcPort
	params.CoinType = litecoinParams.CoinType
}

// applyMachinecoinParams applies the relevant chain configuration parameters that
// differ for machinecoin to the chain parameters typed for btcsuite derivation.
// This function is used in place of using something like interface{} to
// abstract over _which_ chain (or fork) the parameters are for.
func applyMachinecoinParams(params *bitcoinNetParams, machinecoinParams *machinecoinNetParams) {
	params.Name = machinecoinParams.Name
	params.Net = bitcoinWire.BitcoinNet(machinecoinParams.Net)
	params.DefaultPort = machinecoinParams.DefaultPort
	params.CoinbaseMaturity = machinecoinParams.CoinbaseMaturity

	copy(params.GenesisHash[:], machinecoinParams.GenesisHash[:])

	// Address encoding magics
	params.PubKeyHashAddrID = machinecoinParams.PubKeyHashAddrID
	params.ScriptHashAddrID = machinecoinParams.ScriptHashAddrID
	params.PrivateKeyID = machinecoinParams.PrivateKeyID
	params.WitnessPubKeyHashAddrID = machinecoinParams.WitnessPubKeyHashAddrID
	params.WitnessScriptHashAddrID = machinecoinParams.WitnessScriptHashAddrID
	params.Bech32HRPSegwit = machinecoinParams.Bech32HRPSegwit

	copy(params.HDPrivateKeyID[:], machinecoinParams.HDPrivateKeyID[:])
	copy(params.HDPublicKeyID[:], machinecoinParams.HDPublicKeyID[:])

	params.HDCoinType = machinecoinParams.HDCoinType

	checkPoints := make([]chaincfg.Checkpoint, len(machinecoinParams.Checkpoints))
	for i := 0; i < len(machinecoinParams.Checkpoints); i++ {
		var chainHash chainhash.Hash
		copy(chainHash[:], machinecoinParams.Checkpoints[i].Hash[:])

		checkPoints[i] = chaincfg.Checkpoint{
			Height: machinecoinParams.Checkpoints[i].Height,
			Hash:   &chainHash,
		}
	}
	params.Checkpoints = checkPoints

	params.rpcPort = machinecoinParams.rpcPort
	params.CoinType = machinecoinParams.CoinType
}

// isTestnet tests if the given params correspond to a testnet
// parameter configuration.
func isTestnet(params *bitcoinNetParams) bool {
	switch params.Params.Net {
	case bitcoinWire.TestNet3, bitcoinWire.BitcoinNet(litecoinWire.TestNet4), bitcoinWire.BitcoinNet(machinecoinWire.TestNet4):
		return true
	default:
		return false
	}
}
