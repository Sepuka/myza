package btc

import (
	"fmt"
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

	CryptoAddress struct {
		pub fmt.Stringer
		wif fmt.Stringer
		uid uint32
	}
)

// NewBIP32AddrGenerator creates BIP32 HD key generator
func NewBIP32AddrGenerator(
	cfg config.Crypto,
	net *chaincfg.Params,
) *BIP32AddrGenerator {
	return &BIP32AddrGenerator{
		cfg: cfg,
		net: net,
	}
}

// Generate generates address for user
func (g *BIP32AddrGenerator) Generate(ctx *domain.AddressGeneratorContext) (domain.Address, error) {
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

	return NewCryptoAddress(coinbaseKey, g.net, ctx.UserId)
}

func keyToAddr(key *btcec.PrivateKey, net *chaincfg.Params) (btcutil.Address, error) {
	serializedKey := key.PubKey().SerializeCompressed()
	pubKeyAddr, err := btcutil.NewAddressPubKey(serializedKey, net)
	if err != nil {
		return nil, err
	}

	return pubKeyAddr.AddressPubKeyHash(), nil
}

func NewCryptoAddress(key *btcec.PrivateKey, net *chaincfg.Params, uid uint32) (*CryptoAddress, error) {
	wif, err := btcutil.NewWIF(key, net, false)
	if err != nil {
		return nil, err
	}

	coinbaseAddr, err := keyToAddr(key, net)
	if err != nil {
		return nil, err
	}

	return &CryptoAddress{
		pub: coinbaseAddr,
		wif: wif,
		uid: uid,
	}, nil
}

func (a *CryptoAddress) Pub() string {
	return a.pub.String()
}

func (a *CryptoAddress) Wif() string {
	return a.wif.String()
}

func (a *CryptoAddress) Uid() uint32 {
	return a.uid
}
