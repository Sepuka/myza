package http

import (
	"github.com/sarulabs/di/v2"
	"github.com/sepuka/myza/def"
	"github.com/sepuka/myza/def/log"
	"github.com/sepuka/myza/internal/config"
	http2 "github.com/sepuka/myza/internal/http"
	"go.uber.org/zap"
	"net/http"
)

const GateDef = `http.gate`

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: GateDef,
			Build: func(ctx di.Container) (interface{}, error) {
				var (
					logger     = ctx.Get(log.LoggerDef).(*zap.SugaredLogger)
					httpClient = ctx.Get(ClientDef).(*http.Client)
				)

				return http2.NewGate(httpClient, logger), nil
			},
		})
	})
}
