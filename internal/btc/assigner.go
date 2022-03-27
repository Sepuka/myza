package btc

import (
	"github.com/sepuka/myza/domain"
	domain2 "github.com/sepuka/vkbotserver/domain"
	"go.uber.org/zap"
)

type (
	CryptoAddressAssigner struct {
		generator  domain.CryptoAddressGenerator
		logger     *zap.SugaredLogger
		cryptoRepo domain.CryptoRepository
	}
)

func NewCryptoAddressAssigner(
	generator domain.CryptoAddressGenerator,
	log *zap.SugaredLogger,
	cryptoRepo domain.CryptoRepository,
) *CryptoAddressAssigner {
	return &CryptoAddressAssigner{
		generator:  generator,
		logger:     log,
		cryptoRepo: cryptoRepo,
	}
}

func (c *CryptoAddressAssigner) AssignBtc(user *domain2.User) {
	var (
		address domain.Address
		crypto  *domain.Crypto
		context = domain.AddressGeneratorContext{
			Currency: domain.Btc,
			UserId:   uint32(user.UserId),
		}
		err error
	)

	crypto = c.cryptoRepo.Get(user, crypto.Currency)
	// address has already assigned
	if crypto != nil {
		return
	}

	if address, err = c.generator.Generate(context); err != nil {
		c.
			logger.
			With(
				zap.Int(`user_id`, user.UserId),
				zap.String(`currency`, string(crypto.Currency)),
			).
			Error(`Could not generate crypto address`)

		return
	}

	crypto = &domain.Crypto{
		Currency: domain.Btc,
		UserId:   uint32(user.UserId),
		Address:  address.String(),
	}

	err = c.cryptoRepo.Assign(crypto, address)

	if err != nil {
		c.
			logger.
			With(
				zap.Int(`user_id`, user.UserId),
				zap.String(`currency`, string(crypto.Currency)),
			).
			Error(`Could not assign crypto address`)
	}
}
