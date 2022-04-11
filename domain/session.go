package domain

import (
	"github.com/sepuka/vkbotserver/domain"
	"time"
)

type (
	Session struct {
		UserId   int
		Token    string
		Datetime time.Time    `pg:"default:now(),notnull"`
		OAuth    domain.Oauth `pg:"notnull"`
	}
)
