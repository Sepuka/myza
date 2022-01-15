package handler

import (
	"github.com/sepuka/vkbotserver/api"
	"github.com/sepuka/vkbotserver/api/button"
	"github.com/sepuka/vkbotserver/domain"
)

const (
	withdrawMsg = `не реализовано`
)

type (
	withdrawHandler struct {
		api *api.Api
	}
)

func NewWithdrawHandler(api *api.Api) *withdrawHandler {
	return &withdrawHandler{api: api}
}

func (h *withdrawHandler) Handle(req *domain.Request, payload *button.Payload) error {
	var (
		peerId = int(req.Object.Message.FromId)
	)

	return h.api.SendMessage(peerId, withdrawMsg)
}
