package token_store

import (
	"github.com/sarulabs/di/v2"
	"github.com/sepuka/myza/def"
	cache2 "github.com/sepuka/myza/def/cache"
	"github.com/sepuka/myza/def/log"
	"github.com/sepuka/myza/domain"
	"github.com/sepuka/myza/internal/config"
	"github.com/sepuka/myza/internal/token_store"
	"go.uber.org/zap"
)

const (
	VkTokenStoreDef = `def.store_token.vk`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: VkTokenStoreDef,
			Build: func(ctx di.Container) (interface{}, error) {
				var (
					cache  = ctx.Get(cache2.CacheDef).(domain.Cache)
					logger = ctx.Get(log.LoggerDef).(*zap.SugaredLogger)
				)

				return token_store.NewTokenStore(cache, logger), nil
			},
		})
	})
}
