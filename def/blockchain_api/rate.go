package blockchain_api

import (
	"github.com/sarulabs/di/v2"
	"github.com/sepuka/myza/def"
	cache2 "github.com/sepuka/myza/def/cache"
	"github.com/sepuka/myza/def/http"
	"github.com/sepuka/myza/def/log"
	"github.com/sepuka/myza/domain"
	"github.com/sepuka/myza/internal/blockchain_api"
	"github.com/sepuka/myza/internal/config"
	http3 "github.com/sepuka/myza/internal/http"
	"go.uber.org/zap"
)

const (
	BlockchainApiDef = `blockchain.api.def`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: BlockchainApiDef,
			Build: func(ctx di.Container) (interface{}, error) {
				var (
					logger = ctx.Get(log.LoggerDef).(*zap.SugaredLogger)
					gate   = ctx.Get(http.GateDef).(*http3.Gate)
					cache  = ctx.Get(cache2.CacheDef).(domain.Cache)
				)
				return blockchain_api.NewExchangeConverter(logger, cache, gate), nil
			},
		},
		)
	})
}
