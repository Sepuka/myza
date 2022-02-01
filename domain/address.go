package domain

import "github.com/btcsuite/btcutil"

type (
	CryptoAddress struct {
		addr Address
	}

	Address interface {
		String() string
	}

	CryptoAddressGenerator interface {
		Generate() (Address, error)
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
