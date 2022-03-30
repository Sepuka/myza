package domain

import (
	"github.com/btcsuite/btcutil"
	"github.com/sepuka/vkbotserver/domain"
)

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

	Crypto struct {
		Currency CryptoCurrency `pg:",pk"`
		Address  string         `pg:",pk"`
		UserId   uint32         `sql:",fk"`
		User     *domain.User
	}

	// CryptoRepository keeps crypto user`s addresses
	CryptoRepository interface {
		// Assign inserts crypto address to user
		Assign(*Crypto, Address) error
		// Get fetches convenient entity
		Get(user *domain.User, currency CryptoCurrency) *Crypto
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
