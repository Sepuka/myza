package domain

import "github.com/btcsuite/btcutil"

const (
	MainNet = `main`
	TestNet = `test`
)

type (
	CryptoAddress struct {
		addr Address
	}

	AddressGeneratorContext struct {
		Currency CryptoCurrency
		UserId   uint32
	}

	Address interface {
		String() string
	}

	CryptoAddressGenerator interface {
		Generate(context AddressGeneratorContext) (Address, error)
	}
)

func (a *CryptoAddress) String() string {
	switch a.addr.(type) {
	case btcutil.Address:
		return a.String()
	default:
		return ``
	}
}
