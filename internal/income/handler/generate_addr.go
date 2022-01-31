package handler

import (
	"fmt"
	"github.com/btcsuite/btcutil"
	"github.com/sepuka/myza/internal/btc"
	"github.com/sepuka/vkbotserver/api"
	"github.com/sepuka/vkbotserver/api/button"
	"github.com/sepuka/vkbotserver/domain"
)

type (
	generateAddrHandler struct {
		api *api.Api
	}
)

func NewGenerateAddrHandler(api *api.Api) *generateAddrHandler {
	return &generateAddrHandler{api: api}
}

func (h *generateAddrHandler) Handle(req *domain.Request, payload *button.Payload) error {
	var (
		peerId    = int(req.Object.Message.FromId)
		btcUserId = uint32(peerId)
		addr      btcutil.Address
		err       error
	)

	if addr, err = btc.NewAddr(btcUserId); err != nil {
		return err
	}

	return h.api.SendMessage(peerId, fmt.Sprintf(`your address is %s`, addr.String()))
}
