package cmd

import (
	"fmt"
	"github.com/sepuka/myza/def"
	"github.com/spf13/cobra"
	"os"
)

var (
	configFile string

	rootCmd = &cobra.Command{
		Use:  `myza`,
		Args: cobra.MinimumNArgs(1),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return def.Build(configFile)
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "/path/to/config.yml")
	_ = rootCmd.MarkPersistentFlagRequired("config")
}
