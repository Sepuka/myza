package income

import (
	"encoding/json"
	"fmt"
	"github.com/sepuka/myza/internal/income/handler"
	"github.com/sepuka/vkbotserver/api"
	"github.com/sepuka/vkbotserver/api/button"
	"github.com/sepuka/vkbotserver/domain"
	"github.com/sepuka/vkbotserver/errors"
	"github.com/sepuka/vkbotserver/message"
	"go.uber.org/zap"
	"net/http"
)

type (
	// MessageNew handles VK message "message_new"
	MessageNew struct {
		logger      *zap.SugaredLogger
		btnHandlers map[string]message.Handler
		textHandler *handler.Text
	}
)

// NewMessageNew is a MessageNew's constructor
func NewMessageNew(
	logger *zap.SugaredLogger,
	btnHandlers map[string]message.Handler,
	answerHandler *handler.Text,
) *MessageNew {
	return &MessageNew{
		logger:      logger,
		btnHandlers: btnHandlers,
		textHandler: answerHandler,
	}
}

func (o *MessageNew) Exec(req *domain.Request, resp http.ResponseWriter) error {
	var (
		err            error
		rawPayload     = req.Object.Message.Payload
		payloadData    button.Payload
		isHandlerKnown bool
		buttonHandler  message.Handler
	)

	defer func() {
		_, err = resp.Write(api.DefaultResponseBody())
	}()

	if req.IsKeyboardButton() {
		if err = json.Unmarshal([]byte(rawPayload), &payloadData); err != nil {
			return o.buildPayloadError(req.Object.Message, err, `invalid payload JSON`)
		}

		if buttonHandler, isHandlerKnown = o.btnHandlers[payloadData.Command]; isHandlerKnown {
			o.logger.Debugf(`Got "%s" command`, payloadData.Command)
			if err = buttonHandler.Handle(req, &payloadData); err != nil {
				o.logger.Error(err)
			}
		}
	} else {
		if err = o.textHandler.Handle(req); err != nil {
			return err
		}
	}

	return err
}

func (o *MessageNew) buildPayloadError(msg domain.Message, err error, text string) error {
	var (
		rawPayload = msg.Payload
		userId     = msg.FromId
	)

	o.
		logger.
		With(
			zap.String(`json`, rawPayload),
			zap.Int32(`user_id`, userId),
			zap.Error(err),
		).
		Error(text)

	return errors.NewInvalidJsonError(fmt.Sprintf(`%v`, msg), err)
}

func (o *MessageNew) String() string {
	return `message_new`
}
