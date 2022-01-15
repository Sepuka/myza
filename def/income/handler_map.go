package income

import (
	"fmt"
	"github.com/sarulabs/di/v2"
	"github.com/sepuka/myza/def"
	"github.com/sepuka/myza/internal/config"
	"github.com/sepuka/vkbotserver/message"
)

const (
	HandlerMapDef = `handler.map.def`
	ExecutorDef   = `handler.command.def`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: HandlerMapDef,
			Build: func(ctx di.Container) (interface{}, error) {
				var (
					handlers   = def.GetByTag(ExecutorDef)
					handlerMap = make(message.HandlerMap, len(handlers))
					msgName    string
				)

				for _, cmd := range handlers {
					msgName = cmd.(fmt.Stringer).String()
					handlerMap[msgName] = cmd.(message.Executor)
				}

				return handlerMap, nil
			},
		})
	})
}
