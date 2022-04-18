package worker

import (
	"github.com/sarulabs/di/v2"
	"github.com/sepuka/myza/def"
	"github.com/sepuka/myza/def/btc"
	"github.com/sepuka/myza/def/log"
	"github.com/sepuka/myza/domain"
	"github.com/sepuka/myza/internal/config"
	"github.com/sepuka/myza/internal/worker"
	"go.uber.org/zap"
)

const (
	WifDef = `worker.wif.ef`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: WifDef,
			Build: func(ctx di.Container) (interface{}, error) {
				var (
					generator = ctx.Get(btc.Bip32GeneratorDef).(domain.CryptoAddressGenerator)
					logger    = ctx.Get(log.LoggerDef).(*zap.SugaredLogger)
				)

				return worker.NewWif(generator, logger), nil
			},
		})
	})
}
