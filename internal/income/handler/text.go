package handler

import (
	domain2 "github.com/sepuka/myza/domain"
	"github.com/sepuka/vkbotserver/api"
	"github.com/sepuka/vkbotserver/domain"
	"go.uber.org/zap"
)

const (
	addrLen = 34
)

type (
	// Text encapsulates users' requests handler
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

// Handle is a custom user's requests factory which creates concrete handler
func (h *Text) Handle(vkReq *domain.Request) error {
	var (
		msg         = vkReq.Object.Message.Text
		peerId      = int(vkReq.Object.Message.FromId)
		userRequest = domain2.NewTextRequest(peerId, msg)
		handler     domain2.TextHandler
	)

	if len(msg) == addrLen {
		handler = h.textHandlers[BalanceIdHandler]
	} else {
		handler = h.textHandlers[UnknownIdHandler]
	}

	return handler.Handle(userRequest)
}
