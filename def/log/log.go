package log

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/myza/def"
	"github.com/sepuka/myza/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	errPkg "github.com/pkg/errors"
)

const LoggerDef = `logger.def`

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: LoggerDef,
			Build: func(container di.Container) (interface{}, error) {
				var (
					err               error
					logger            *zap.Logger
					sugar             *zap.SugaredLogger
					zapCfg            zap.Config
					core              zapcore.Core
					fileEncoder       zapcore.Encoder
					fileEncoderConfig zapcore.EncoderConfig
				)

				fileSynchronizer, closeOut, err := zap.Open(`stdout`)
				if err != nil {
					return nil, errPkg.Wrap(err, `unable to open output files`)
				}

				writeSyncer := zapcore.AddSync(fileSynchronizer)

				consoleMsgLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
					if cfg.Log.Prod {
						return lvl >= zapcore.InfoLevel
					}

					return true
				})

				if cfg.Log.Prod {
					zapCfg = zap.NewProductionConfig()
					fileEncoderConfig = zap.NewProductionEncoderConfig()
				} else {
					zapCfg = zap.NewDevelopmentConfig()
					fileEncoderConfig = zap.NewDevelopmentEncoderConfig()
				}

				zapCfg.OutputPaths = []string{`stdout`}

				fileEncoder = zapcore.NewJSONEncoder(fileEncoderConfig)
				core = zapcore.NewTee(
					zapcore.NewCore(fileEncoder, writeSyncer, consoleMsgLevel),
				)

				logger = zap.New(core)
				sugar = logger.Sugar()
				if sugar == nil {
					closeOut()
					return nil, errPkg.New(`unable build sugar logger`)
				}

				return sugar, err
			},
			Close: func(obj interface{}) error {
				logger := obj.(*zap.SugaredLogger)
				return logger.Sync()
			},
		})
	})
}
