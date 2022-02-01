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
		cfg   config.Crypto
		child uint32
	}
)

func NewBIP32AddrGenerator(
	cfg config.Crypto,
	child uint32,
) *BIP32AddrGenerator {
	return &BIP32AddrGenerator{
		cfg:   cfg,
		child: child,
	}
}

func (g *BIP32AddrGenerator) Generate() (domain.Address, error) {
	var net = &chaincfg.MainNetParams
	hdRoot, err := hdkeychain.NewMaster([]byte(g.cfg.Seed), net)
	if err != nil {
		return nil, err
	}

	// The first child key from the hd root is reserved as the coinbase
	// generation address.
	coinbaseChild, err := hdRoot.Derive(g.child)
	if err != nil {
		return nil, err
	}
	coinbaseKey, err := coinbaseChild.ECPrivKey()
	if err != nil {
		return nil, err
	}
	coinbaseAddr, err := keyToAddr(coinbaseKey, net)
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
