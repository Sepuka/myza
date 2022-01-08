package handler

import (
	domain2 "github.com/sepuka/myza/domain"
	"github.com/sepuka/vkbotserver/api"
	"github.com/sepuka/vkbotserver/domain"
	"go.uber.org/zap"
)

type (
	Text struct {
		api          *api.Api
		log          *zap.SugaredLogger
		textHandlers map[string]domain2.TextHandler
	}
)

// NewText parses text requests
func NewText(
	api *api.Api,
	log *zap.SugaredLogger,
	textHandlers map[string]domain2.TextHandler,
) *Text {
	return &Text{
		api:          api,
		log:          log,
		textHandlers: textHandlers,
	}
}

func (h *Text) Handle(vkReq *domain.Request) error {
	var (
		msg                = vkReq.Object.Message.Text
		peerId             = int(vkReq.Object.Message.FromId)
		userRequest        = domain2.NewTextRequest(peerId, msg)
		userRequestHandler domain2.TextHandler
	)

	if len(msg) == 34 {
		userRequestHandler = h.textHandlers[BalanceIdHandler]
	} else {
		userRequestHandler = h.textHandlers[UnknownIdHandler]
	}

	return userRequestHandler.Handle(userRequest)
}
