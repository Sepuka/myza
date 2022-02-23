package db

import (
	"github.com/go-pg/pg"
	"github.com/sarulabs/di/v2"
	"github.com/sepuka/myza/def"
	"github.com/sepuka/myza/internal/config"
	db2 "github.com/sepuka/myza/internal/db"
	"net"
	"strconv"
)

const (
	DatabasePgDef = `postgres.db.def`
	UserRepoDef   = `repo.user.def`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: DatabasePgDef,
			Build: func(ctx di.Container) (interface{}, error) {
				var (
					db *pg.DB
				)

				db = pg.Connect(&pg.Options{
					User:     cfg.Db.User,
					Password: cfg.Db.Password,
					Addr:     net.JoinHostPort(cfg.Db.Host, strconv.Itoa(cfg.Db.Port)),
					Database: cfg.Db.Name,
				})

				_, err := db.Exec(`SET timezone TO 'UTC'`)

				return db, err
			},
			Close: func(obj interface{}) error {
				return obj.(*pg.DB).Close()
			},
		},
		)
	})

	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: UserRepoDef,
			Build: func(ctx di.Container) (interface{}, error) {
				var (
					db = ctx.Get(DatabasePgDef).(*pg.DB)
				)

				return db2.NewUserRepository(db), nil
			},
		})
	})
}
