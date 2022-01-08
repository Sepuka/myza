package http

import (
	"github.com/sarulabs/di/v2"
	"github.com/sepuka/myza/def"
	"github.com/sepuka/myza/internal/config"
	"net/http"
)

const ClientDef = `http.client`

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: ClientDef,
			Build: func(ctx di.Container) (interface{}, error) {
				return &http.Client{}, nil
			},
		})
	})
}
