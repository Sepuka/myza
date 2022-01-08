package text

import "github.com/sepuka/myza/domain"

type (
	UnknownRequestHandler struct {
	}
)

func NewUnknownRequestHandler() *UnknownRequestHandler {
	return &UnknownRequestHandler{}
}

func (b *UnknownRequestHandler) Handle(req domain.TextRequest) error {
	return nil
}
