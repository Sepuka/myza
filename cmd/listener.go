package cmd

import (
	"github.com/sepuka/myza/def"
	defBot "github.com/sepuka/myza/def/bot"
	"github.com/sepuka/vkbotserver/server"
	"github.com/spf13/cobra"
)

var (
	vkBot = &cobra.Command{
		Use: `vkbot`,
		RunE: func(cmd *cobra.Command, args []string) error {
			instance, err := def.Container.SafeGet(defBot.Bot)

			if err != nil {
				return err
			}

			return instance.(*server.SocketServer).Listen()
		},
	}
)

func init() {
	rootCmd.AddCommand(vkBot)
}
