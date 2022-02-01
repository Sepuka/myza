package income

import (
	"github.com/sarulabs/di/v2"
	"github.com/sepuka/myza/def"
	"github.com/sepuka/myza/def/blockchain_api"
	cache2 "github.com/sepuka/myza/def/cache"
	"github.com/sepuka/myza/def/http"
	"github.com/sepuka/myza/def/log"
	"github.com/sepuka/myza/def/vkapi/method"
	"github.com/sepuka/myza/domain"
	"github.com/sepuka/myza/internal/config"
	msgHandler "github.com/sepuka/myza/internal/income"
	"github.com/sepuka/myza/internal/income/button"
	"github.com/sepuka/myza/internal/income/handler"
	message3 "github.com/sepuka/myza/internal/income/handler/message"
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
					cache          = ctx.Get(cache2.CacheDef).(domain.Cache)
					blockchainApi  = ctx.Get(blockchain_api.BlockchainApiDef).(domain.ExchangeRateConverter)
					buttonHandlers = map[string]message.Handler{
						button.StartIdButton:        handler.NewStartHandler(vkApi),
						button.WithdrawIdButton:     handler.NewWithdrawHandler(vkApi),
						button.GenerateAddrIdButton: handler.NewGenerateAddrHandler(vkApi, cfg.Crypto),
					}
					textHandlers = map[string]domain.TextHandler{
						handler.UnknownIdHandler: message3.NewUnknownRequestHandler(),
						handler.BalanceIdHandler: message3.NewBalanceRequestHandler(logger, httpClient, vkApi, cache, blockchainApi),
					}
					textHandler = handler.NewText(vkApi, logger, textHandlers)
				)
				return msgHandler.NewMessageNew(logger, buttonHandlers, textHandler), nil
			},
		},
		)
	})
}
