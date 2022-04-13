package cmd

import (
	"github.com/sepuka/myza/def"
	"github.com/sepuka/myza/def/worker"
	worker2 "github.com/sepuka/myza/internal/worker"
	"github.com/spf13/cobra"
)

var (
	balance = &cobra.Command{
		Use: `balance`,
		RunE: func(cmd *cobra.Command, args []string) error {
			instance, err := def.Container.SafeGet(worker.BalanceDef)

			if err != nil {
				return err
			}

			return instance.(*worker2.BalanceUpdater).Work()
		},
	}
)

func init() {
	rootCmd.AddCommand(balance)
}
