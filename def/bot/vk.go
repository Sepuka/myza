package bot

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/focalism/def/middleware"
	"github.com/sepuka/myza/def"
	"github.com/sepuka/myza/def/log"
	messageDef "github.com/sepuka/myza/def/message"
	"github.com/sepuka/myza/internal/config"
	"github.com/sepuka/vkbotserver/message"
	middleware2 "github.com/sepuka/vkbotserver/middleware"
	"github.com/sepuka/vkbotserver/server"
	"go.uber.org/zap"
)

const (
	Bot = `def.bot.vk`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Build: func(container di.Container) (interface{}, error) {
				var (
					handlerMap     = container.Get(messageDef.HandlerMapDef).(message.HandlerMap)
					middlewareList = container.Get(middleware.BotMiddlewareDef).(middleware2.HandlerFunc)
					logger         = container.Get(log.LoggerDef).(*zap.SugaredLogger)
				)

				return server.NewSocketServer(cfg.Server, handlerMap, middlewareList, logger), nil
			},
			Close: nil,
			Name:  Bot,
			Scope: di.App,
		})
	})
}
