package users

import (
	"github.com/sarulabs/di/v2"
	"github.com/sepuka/myza/def"
	"github.com/sepuka/myza/def/db"
	"github.com/sepuka/myza/def/http"
	"github.com/sepuka/myza/def/log"
	"github.com/sepuka/myza/internal/config"
	"github.com/sepuka/vkbotserver/api/users"
	"github.com/sepuka/vkbotserver/domain"
	"go.uber.org/zap"
	http2 "net/http"
)

const (
	ApiUsersGetDef = `api.users.get.def`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: ApiUsersGetDef,
			Build: func(ctx di.Container) (interface{}, error) {
				var (
					apiUsersGet *users.Get
					logger      = ctx.Get(log.LoggerDef).(*zap.SugaredLogger)
					httpClient  = ctx.Get(http.ClientDef).(*http2.Client)
					userRepo    = ctx.Get(db.UserRepoDef).(domain.UserRepository)
				)

				apiUsersGet = users.NewGet(httpClient, logger, userRepo)

				return apiUsersGet, nil
			},
		},
		)
	})
}
