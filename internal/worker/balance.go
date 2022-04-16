package worker

import (
	"github.com/sepuka/myza/domain"
	"github.com/sepuka/myza/internal/blockchain_api"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	updatedLimit   = 10
	outdatedPeriod = time.Hour * 2
)

type (
	BalanceUpdater struct {
		cryptoRepo domain.CryptoRepository
		score      *blockchain_api.Score
		converter  domain.ExchangeRateConverter
		logger     *zap.SugaredLogger
	}
)

func NewBalanceUpdater(
	cryptoRepo domain.CryptoRepository,
	balanceChecker *blockchain_api.Score,
	converter domain.ExchangeRateConverter,
	logger *zap.SugaredLogger,
) *BalanceUpdater {
	return &BalanceUpdater{
		cryptoRepo: cryptoRepo,
		score:      balanceChecker,
		converter:  converter,
		logger:     logger,
	}
}

func (w *BalanceUpdater) Work() error {
	var (
		stop    bool
		signals = make(chan os.Signal, 1)
	)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signals
		stop = true
	}()

	for !stop {
		go w.update()
		time.Sleep(outdatedPeriod)
	}

	return nil
}

func (w *BalanceUpdater) update() {
	var (
		crypto      *domain.Crypto
		cryptos     []*domain.Crypto
		err         error
		wallet      *domain.Wallet
		fiatBalance float64
		date        = time.Now().Add(-outdatedPeriod)
	)

	if cryptos, err = w.cryptoRepo.FindOutdated(date, updatedLimit); err != nil {
		w.
			logger.
			With(
				zap.Error(err),
				zap.Time(`from`, date),
			).
			Error(`Could not find any outdated cryptos.`)

		return
	}

	for _, crypto = range cryptos {
		if crypto.Address == `` {
			continue
		}

		if wallet, err = w.score.GetBalance(crypto); err == nil {
			if fiatBalance, err = w.converter.Convert(wallet, domain.Rub); err == nil {
				crypto.Balance = wallet.FinalBalance
				crypto.Fiat = fiatBalance
				if err = w.cryptoRepo.UpdateBalance(crypto); err != nil {
					w.
						logger.
						With(
							zap.Error(err),
							zap.Float64(`fiat balance`, fiatBalance),
							zap.Float64(`crypto balance`, wallet.FinalBalance),
						).
						Error(`Could not update wallet balance.`)
				}
			} else {
				w.
					logger.
					With(
						zap.Error(err),
						zap.Float64(`fiat balance`, fiatBalance),
						zap.Float64(`crypto balance`, wallet.FinalBalance),
					).
					Error(`Could not convert wallet value to fiat currency.`)

				return
			}
		} else {
			return
		}
	}

}
