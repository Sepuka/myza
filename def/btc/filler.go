package btc

import (
	"github.com/sarulabs/di/v2"
	"github.com/sepuka/myza/def"
	"github.com/sepuka/myza/def/db"
	"github.com/sepuka/myza/def/log"
	"github.com/sepuka/myza/domain"
	"github.com/sepuka/myza/internal/btc"
	"github.com/sepuka/myza/internal/config"
	"go.uber.org/zap"
)

const (
	CryptoFillerDef = `btc.filler.def`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: CryptoFillerDef,
			Build: func(ctx di.Container) (interface{}, error) {
				var (
					logger     = ctx.Get(log.LoggerDef).(*zap.SugaredLogger)
					generator  = ctx.Get(Bip32GeneratorDef).(domain.CryptoAddressGenerator)
					cryptoRepo = ctx.Get(db.CryptoRepoDef).(domain.CryptoRepository)
				)

				return btc.NewCryptoAddressAssigner(generator, logger, cryptoRepo), nil
			},
		})
	})
}
