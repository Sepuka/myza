package btc

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/sarulabs/di/v2"
	"github.com/sepuka/myza/def"
	"github.com/sepuka/myza/domain"
	"github.com/sepuka/myza/errors"
	"github.com/sepuka/myza/internal/btc"
	"github.com/sepuka/myza/internal/config"
)

const (
	Bip32GeneratorDef = `btc.bip32.def`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: Bip32GeneratorDef,
			Build: func(ctx di.Container) (interface{}, error) {
				var (
					net *chaincfg.Params
				)

				switch cfg.Crypto.Net {
				case domain.MainNet:
					net = &chaincfg.MainNetParams
				case domain.TestNet:
					net = &chaincfg.TestNet3Params
				default:
					return nil, errors.NewInvalidBlockchainNetError(`check crypto.net config`)
				}

				return btc.NewBIP32AddrGenerator(cfg.Crypto, net), nil
			},
		},
		)
	})
}
