package callback

import (
	"github.com/sepuka/myza/internal/btc"
	"github.com/sepuka/vkbotserver/domain"
)

func AssignBtcAddress(assigner *btc.CryptoAddressAssigner, user *domain.User) {
	assigner.AssignBtc(user)
}
