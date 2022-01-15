package income

import (
	"github.com/sarulabs/di/v2"
	"github.com/sepuka/myza/def"
	"github.com/sepuka/myza/internal/config"
	"github.com/sepuka/vkbotserver/message"
)

const (
	ConfirmationDef = `def.message.confirmation`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: ConfirmationDef,
			Tags: []di.Tag{
				{
					Name: ExecutorDef,
					Args: nil,
				},
			},
			Build: func(ctx di.Container) (interface{}, error) {
				return message.NewConfirmation(cfg.Server), nil
			},
		})
	})
}
