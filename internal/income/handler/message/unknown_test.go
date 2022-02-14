package message

import (
	"github.com/sepuka/myza/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnknownRequestHandler(t *testing.T) {
	var (
		handler = NewUnknownRequestHandler()
		req     domain.TextRequest
	)
	req = domain.NewTextRequest(0, ``)
	assert.Nil(t, handler.Handle(req))
}
