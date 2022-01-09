package cache

import (
	"github.com/go-redis/redis/v8"
	"github.com/sarulabs/di/v2"
	"github.com/sepuka/myza/def"
	"github.com/sepuka/myza/internal/config"
)

const CacheDef = `cache.redis.def`

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: CacheDef,
			Build: func(ctx di.Container) (interface{}, error) {
				var (
					client = redis.NewClient(&redis.Options{})
				)

				return client, nil
			},
			Close: func(obj interface{}) error {
				return obj.(*redis.Client).Close()
			},
		})
	})
}
