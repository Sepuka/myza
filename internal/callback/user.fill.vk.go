package callback

import (
	"github.com/sepuka/vkbotserver/api/users"
	"github.com/sepuka/vkbotserver/domain"
)

func FillVkUserName(get *users.Get, user *domain.User) {
	if !user.IsFilledPersonalData() {
		get.FillUser(user)
	}
}
