package domain

import (
	"github.com/sepuka/vkbotserver/domain"
	"time"
)

const (
	MainNet = `main`
	TestNet = `test`
)

type (
	AddressGeneratorContext struct {
		Currency CryptoCurrency
		UserId   uint32
	}

	Address interface {
		Pub() string
		Wif() string
		Uid() uint32
	}

	CryptoAddressGenerator interface {
		Generate(context *AddressGeneratorContext) (Address, error)
	}

	// Crypto represents the last state of user`s balance
	Crypto struct {
		Currency  CryptoCurrency `pg:",pk"`
		Address   string         `pg:",pk"`
		UserId    uint32         `sql:",fk"`
		Balance   float64
		UpdatedAt time.Time
		Fiat      float64
		User      *domain.User
	}

	// CryptoRepository keeps crypto user`s addresses
	CryptoRepository interface {
		// AssignAddress inserts crypto address to user
		AssignAddress(*Crypto, Address) error
		// Get fetches convenient entity
		Get(user *domain.User, currency CryptoCurrency) *Crypto
		// UpdateBalance updates balance and fiat fields only
		UpdateBalance(user *Crypto) error
		// FindOutdated looking for outdated rows
		FindOutdated(date time.Time, limit int) ([]*Crypto, error)
	}
)
