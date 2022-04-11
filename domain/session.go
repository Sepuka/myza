package domain

import (
	"github.com/sepuka/vkbotserver/domain"
	"time"
)

type (
	Session struct {
		UserId   int
		Token    string
		DateTime time.Time    `pg:"default:now(),notnull"`
		Oauth    domain.Oauth `pg:"notnull"`
	}
)
