package handler

import (
	"fmt"
	domain2 "github.com/sepuka/myza/domain"
	"github.com/sepuka/myza/internal/btc"
	"github.com/sepuka/myza/internal/config"
	"github.com/sepuka/vkbotserver/api"
	"github.com/sepuka/vkbotserver/api/button"
	"github.com/sepuka/vkbotserver/domain"
)

type (
	generateAddrHandler struct {
		api              *api.Api
		btcAddrGenerator domain2.CryptoAddressGenerator
		cfg              config.Crypto
	}
)

// NewGenerateAddrHandler handles requests in order to return user's crypto address
func NewGenerateAddrHandler(
	api *api.Api,
	cfg config.Crypto,
) *generateAddrHandler {
	return &generateAddrHandler{
		api: api,
		cfg: cfg,
	}
}

func (h *generateAddrHandler) Handle(req *domain.Request, payload *button.Payload) error {
	var (
		peerId    = int(req.Object.Message.FromId)
		addr      domain2.Address
		err       error
		generator domain2.CryptoAddressGenerator
	)

	if generator, err = h.generatorFactoryMethod(req); err != nil {
		return err
	}

	if addr, err = generator.Generate(); err != nil {
		return err
	}

	return h.api.SendMessage(peerId, fmt.Sprintf(`your address is %s`, addr.String()))
}

func (h *generateAddrHandler) generatorFactoryMethod(req *domain.Request) (domain2.CryptoAddressGenerator, error) {
	return btc.NewBIP32AddrGenerator(h.cfg, uint32(req.Object.Message.FromId)), nil
}
