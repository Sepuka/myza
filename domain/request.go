package domain

type (
	TextRequest struct {
		peerId         int
		msg            string
		CryptoCurrency CryptoCurrency
		FiatCurrency   FiatCurrency
	}

	TextHandler interface {
		Handle(request TextRequest) error
	}
)

// NewTextRequest handles variety texts
func NewTextRequest(peerId int, msg string) TextRequest {
	return TextRequest{
		peerId:         peerId,
		msg:            msg,
		CryptoCurrency: Btc,
		FiatCurrency:   Rub,
	}
}

func (r *TextRequest) GetMessage() string {
	return r.msg
}

func (r *TextRequest) GetPeerId() int {
	return r.peerId
}
