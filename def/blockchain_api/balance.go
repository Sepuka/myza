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
	"go.uber.org/zap"
	http2 "net/http"
)

const (
	ScoreWalletDef = `blockchain.api.score.def`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: ScoreWalletDef,
			Build: func(ctx di.Container) (interface{}, error) {
				var (
					logger = ctx.Get(log.LoggerDef).(*zap.SugaredLogger)
					cache  = ctx.Get(cache2.CacheDef).(domain.Cache)
					client = ctx.Get(http.ClientDef).(*http2.Client)
				)
				return blockchain_api.NewScoreFetcher(logger, client, cache), nil
			},
		},
		)
	})
}
