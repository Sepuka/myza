package income

import (
	"github.com/sarulabs/di/v2"
	"github.com/sepuka/myza/def"
	"github.com/sepuka/myza/def/db"
	"github.com/sepuka/myza/def/http"
	"github.com/sepuka/myza/def/log"
	"github.com/sepuka/myza/internal/config"
	"github.com/sepuka/vkbotserver/domain"
	"github.com/sepuka/vkbotserver/message"
	"go.uber.org/zap"
	http2 "net/http"
)

const (
	AuthVkDef = `def.auth.vk`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: AuthVkDef,
			Tags: []di.Tag{
				{
					Name: ExecutorDef,
					Args: nil,
				},
			},
			Build: func(ctx di.Container) (interface{}, error) {
				var (
					logger     = ctx.Get(log.LoggerDef).(*zap.SugaredLogger)
					httpClient = ctx.Get(http.ClientDef).(*http2.Client)
					userRepo   = ctx.Get(db.UserRepoDef).(domain.UserRepository)
				)
				return message.NewAuthVk(cfg.Server.VkOauth, httpClient, logger, userRepo), nil
			},
		})
	})
}
