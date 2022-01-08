package domain

type (
	TextRequest struct {
		peerId int
		msg    string
	}

	TextHandler interface {
		Handle(request TextRequest) error
	}
)

func NewTextRequest(peerId int, msg string) TextRequest {
	return TextRequest{
		peerId: peerId,
		msg:    msg,
	}
}

func (r *TextRequest) GetMessage() string {
	return r.msg
}

func (r *TextRequest) GetPeerId() int {
	return r.peerId
}
