package cmd

import (
	"github.com/sepuka/myza/def"
	"github.com/sepuka/myza/def/worker"
	worker2 "github.com/sepuka/myza/internal/worker"
	"github.com/spf13/cobra"
	"strconv"
)

var (
	wif = &cobra.Command{
		Use: `wif`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				userId int64
				err    error
			)
			userId, err = strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}

			instance, err := def.Container.SafeGet(worker.WifDef)

			if err != nil {
				return err
			}

			return instance.(*worker2.Wif).Print(uint32(userId))
		},
	}
)

func init() {
	rootCmd.AddCommand(wif)
}
