package handler

import (
	"fmt"
	domain2 "github.com/sepuka/myza/domain"
	"github.com/sepuka/myza/internal/btc"
	"github.com/sepuka/vkbotserver/api"
	"github.com/sepuka/vkbotserver/api/button"
	"github.com/sepuka/vkbotserver/domain"
)

type (
	generateAddrHandler struct {
		api              *api.Api
		btcAddrGenerator domain2.CryptoAddressGenerator
	}
)

// NewGenerateBtcAddrHandler handles requests in order to return user's crypto address
func NewGenerateBtcAddrHandler(
	api *api.Api,
	generator domain2.CryptoAddressGenerator,
) *generateAddrHandler {
	return &generateAddrHandler{
		api:              api,
		btcAddrGenerator: generator,
	}
}

func (h *generateAddrHandler) Handle(req *domain.Request, payload *button.Payload) error {
	var (
		peerId  = int(req.Object.Message.FromId)
		addr    domain2.Address
		err     error
		context = btc.NewAddressGeneratorContext(domain2.Btc, uint32(req.Object.Message.FromId))
	)

	if addr, err = h.btcAddrGenerator.Generate(context); err != nil {
		return err
	}

	return h.api.SendMessage(peerId, fmt.Sprintf(`your address is %s`, addr.Pub()))
}
