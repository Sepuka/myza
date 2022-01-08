package message

import (
	"github.com/sarulabs/di/v2"
	"github.com/sepuka/myza/def"
	"github.com/sepuka/myza/def/http"
	"github.com/sepuka/myza/def/log"
	"github.com/sepuka/myza/def/vkapi/method"
	"github.com/sepuka/myza/domain"
	"github.com/sepuka/myza/internal/config"
	msgHandler "github.com/sepuka/myza/internal/message"
	"github.com/sepuka/myza/internal/message/button"
	"github.com/sepuka/myza/internal/message/handler"
	"github.com/sepuka/myza/internal/text"
	api2 "github.com/sepuka/vkbotserver/api"
	"github.com/sepuka/vkbotserver/message"
	"go.uber.org/zap"
	http2 "net/http"
)

const (
	MsgNewDef = `def.message.new`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: MsgNewDef,
			Tags: []di.Tag{
				{
					Name: ExecutorDef,
					Args: nil,
				},
			},
			Build: func(ctx di.Container) (interface{}, error) {
				var (
					logger         = ctx.Get(log.LoggerDef).(*zap.SugaredLogger)
					vkApi          = ctx.Get(method.ApiDef).(*api2.Api)
					httpClient     = ctx.Get(http.ClientDef).(*http2.Client)
					buttonHandlers = map[string]message.Handler{
						button.StartIdButton: handler.NewStartHandler(vkApi),
					}
					textHandlers = map[string]domain.TextHandler{
						handler.UnknownIdHandler: text.NewUnknownRequestHandler(),
						handler.BalanceIdHandler: text.NewBalanceRequestHandler(logger, httpClient, vkApi),
					}
					textHandler = handler.NewText(vkApi, logger, textHandlers)
				)
				return msgHandler.NewMessageNew(logger, buttonHandlers, textHandler), nil
			},
		},
		)
	})
}
