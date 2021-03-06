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
		context = NewAddressGeneratorContext(domain.Btc, uint32(user.UserId))
		err     error
	)

	crypto = c.cryptoRepo.Get(user, context.Currency)
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
		Address:  address.Pub(),
	}

	err = c.cryptoRepo.AssignAddress(crypto, address)

	if err != nil {
		c.
			logger.
			With(
				zap.Int(`user_id`, user.UserId),
				zap.String(`currency`, string(crypto.Currency)),
				zap.Error(err),
			).
			Error(`Could not assign crypto address`)
	}
}
