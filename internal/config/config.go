package config

import (
	"github.com/sepuka/vkbotserver/config"
	"github.com/stevenroose/gonfig"
)

type (
	Log struct {
		Prod bool
	}

	Database struct {
		User     string
		Password string
		Name     string
		Host     string
		Port     int
	}

	Config struct {
		Server config.Config
		Log    Log
		Db     Database
	}
)

func GetConfig(path string) (*Config, error) {
	var (
		cfg = &Config{}
		err = gonfig.Load(cfg, gonfig.Conf{
			FileDefaultFilename: path,
			FlagIgnoreUnknown:   true,
			FlagDisable:         true,
			EnvDisable:          true,
		})
	)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
