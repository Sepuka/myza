package def

import (
	"github.com/sarulabs/di/v2"
	"github.com/sepuka/myza/internal/config"
)

var (
	defs      []containerFnc
	Container di.Container
	err       error
)

type (
	containerFnc func(builder *di.Builder, cfg *config.Config) error
)

func Register(fnc containerFnc) {
	defs = append(defs, fnc)
}

func Build(cfgPath string) error {
	var (
		builder *di.Builder
		cfg     *config.Config
		fnc     containerFnc
	)

	builder, err = di.NewBuilder(di.App, di.Request)
	if err != nil {
		return err
	}

	cfg, err = config.GetConfig(cfgPath)
	if err != nil {
		return err
	}

	for _, fnc = range defs {
		if err = fnc(builder, cfg); err != nil {
			return err
		}
	}

	Container = builder.Build()

	return nil
}

func GetByTag(tag string) []interface{} {
	var defs []interface{}

	for _, def := range Container.Definitions() {
		for _, defTag := range def.Tags {
			if defTag.Name == tag {
				var content interface{}
				if err := Container.Fill(def.Name, &content); err != nil {
					panic(err)
				}
				defs = append(defs, content)
			}
		}
	}

	return defs
}
