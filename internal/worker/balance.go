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
	updatedLimit = 10
)

type (
	BalanceUpdater struct {
		date       time.Time
		cryptoRepo domain.CryptoRepository
		score      *blockchain_api.Score
		logger     *zap.SugaredLogger
	}
)

func NewBalanceUpdater(
	cryptoRepo domain.CryptoRepository,
	balanceChecker *blockchain_api.Score,
	logger *zap.SugaredLogger,
) *BalanceUpdater {
	return &BalanceUpdater{
		cryptoRepo: cryptoRepo,
		score:      balanceChecker,
		logger:     logger,
		date:       time.Now().Add(-time.Hour),
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
		time.Sleep(time.Second * 60)
	}

	return nil
}

func (w *BalanceUpdater) update() {
	var (
		crypto  *domain.Crypto
		cryptos []*domain.Crypto
		err     error
		balance float64
	)

	if cryptos, err = w.cryptoRepo.FindOutdated(w.date, updatedLimit); err != nil {
		w.
			logger.
			With(
				zap.Error(err),
				zap.Time(`from`, w.date),
			).
			Error(`Could not find any outdated cryptos`)

		return
	}

	for _, crypto = range cryptos {
		if balance, err = w.score.GetBalance(crypto); err == nil {
			if err = w.cryptoRepo.UpdateBalance(crypto, balance); err != nil {
				w.
					logger.
					With(
						zap.Error(err),
					).
					Error(`Could not update wallet balance`)
			}
		} else {
			return
		}
	}

}
