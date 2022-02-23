package cmd

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/sepuka/myza/def"
	db2 "github.com/sepuka/myza/def/db"
	domain2 "github.com/sepuka/vkbotserver/domain"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateTables)
}

var (
	generateTables = &cobra.Command{
		Use:     `db generate`,
		Example: `db generate -c /config/path`,

		RunE: func(cmd *cobra.Command, args []string) error {
			var db *pg.DB
			if err := def.Container.Fill(db2.DatabasePgDef, &db); err != nil {
				return err
			}

			for _, model := range []interface{}{&domain2.User{}} {
				err := db.CreateTable(model, &orm.CreateTableOptions{
					FKConstraints: true,
				})
				if err != nil {
					return err
				}
			}

			return nil
		},
	}
)
