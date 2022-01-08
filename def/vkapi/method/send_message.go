package method

import (
	"github.com/sarulabs/di/v2"
	"github.com/sepuka/myza/def"
	"github.com/sepuka/myza/def/http"
	"github.com/sepuka/myza/def/log"
	"github.com/sepuka/myza/internal/config"
	"github.com/sepuka/vkbotserver/api"
	"go.uber.org/zap"
	http2 "net/http"
)

const (
	ApiDef = `def.api.send_message`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: ApiDef,
			Build: func(ctx di.Container) (interface{}, error) {
				var (
					client = ctx.Get(http.ClientDef).(*http2.Client)
					logger = ctx.Get(log.LoggerDef).(*zap.SugaredLogger)
					rnd    = api.NewRnder()
				)
				return api.NewApi(logger, cfg.Server, client, rnd), nil
			},
		})
	})
}
