package btc

import (
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/sepuka/myza/domain"
	"github.com/sepuka/myza/internal/config"
)

type (
	BIP32AddrGenerator struct {
		cfg config.Crypto
		net *chaincfg.Params
	}
)

func NewBIP32AddrGenerator(
	cfg config.Crypto,
	net *chaincfg.Params,
) *BIP32AddrGenerator {
	return &BIP32AddrGenerator{
		cfg: cfg,
		net: net,
	}
}

func (g *BIP32AddrGenerator) Generate(ctx domain.AddressGeneratorContext) (domain.Address, error) {
	hdRoot, err := hdkeychain.NewMaster([]byte(g.cfg.Seed), g.net)
	if err != nil {
		return nil, err
	}

	coinbaseChild, err := hdRoot.Derive(ctx.UserId)
	if err != nil {
		return nil, err
	}
	coinbaseKey, err := coinbaseChild.ECPrivKey()
	if err != nil {
		return nil, err
	}
	coinbaseAddr, err := keyToAddr(coinbaseKey, g.net)
	if err != nil {
		return nil, err
	}

	return coinbaseAddr, nil
}

func keyToAddr(key *btcec.PrivateKey, net *chaincfg.Params) (btcutil.Address, error) {
	serializedKey := key.PubKey().SerializeCompressed()
	pubKeyAddr, err := btcutil.NewAddressPubKey(serializedKey, net)
	if err != nil {
		return nil, err
	}

	return pubKeyAddr.AddressPubKeyHash(), nil
}
