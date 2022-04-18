package worker

import (
	"github.com/sepuka/myza/domain"
	"github.com/sepuka/myza/internal/btc"
	"go.uber.org/zap"
)

type (
	Wif struct {
		generator domain.CryptoAddressGenerator
		logger    *zap.SugaredLogger
	}
)

func NewWif(
	generator domain.CryptoAddressGenerator,
	logger *zap.SugaredLogger,
) *Wif {
	return &Wif{
		generator: generator,
		logger:    logger,
	}
}

func (w *Wif) Print(userId uint32) error {
	var (
		context = btc.NewAddressGeneratorContext(domain.Btc, userId)
		address domain.Address
		err     error
	)

	address, err = w.generator.Generate(context)
	if err != nil {
		return err
	}

	w.
		logger.
		With(
			zap.Uint32(`user_id`, userId),
			zap.String(`public key`, address.Pub()),
			zap.String(`wif`, address.Wif()),
		).
		Info(`Address info has fetched`)

	return nil
}
