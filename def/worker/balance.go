package worker

import (
	"github.com/sarulabs/di/v2"
	"github.com/sepuka/myza/def"
	"github.com/sepuka/myza/def/blockchain_api"
	"github.com/sepuka/myza/def/db"
	"github.com/sepuka/myza/def/log"
	"github.com/sepuka/myza/domain"
	blockchain_api2 "github.com/sepuka/myza/internal/blockchain_api"
	"github.com/sepuka/myza/internal/config"
	"github.com/sepuka/myza/internal/worker"
	"go.uber.org/zap"
)

const (
	BalanceDef = `worker.balance.def`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: BalanceDef,
			Build: func(ctx di.Container) (interface{}, error) {
				var (
					cryptoRepo   = ctx.Get(db.CryptoRepoDef).(domain.CryptoRepository)
					scoreFetcher = ctx.Get(blockchain_api.ScoreWalletDef).(*blockchain_api2.Score)
					logger       = ctx.Get(log.LoggerDef).(*zap.SugaredLogger)
				)

				return worker.NewBalanceUpdater(cryptoRepo, scoreFetcher, logger), nil
			},
		})
	})
}
