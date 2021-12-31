package middleware

import (
	"github.com/sarulabs/di/v2"
	"github.com/sepuka/myza/def"
	"github.com/sepuka/myza/internal/config"
	"github.com/sepuka/vkbotserver/middleware"
)

const (
	BotMiddlewareDef = `middleware.bot.def`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: BotMiddlewareDef,
			Build: func(ctx di.Container) (interface{}, error) {
				var (
					terminalMiddleware = []func(handlerFunc middleware.HandlerFunc) middleware.HandlerFunc{
						middleware.Panic,
					}
				)

				return middleware.BuildHandlerChain(terminalMiddleware), nil
			},
		})
	})
}
