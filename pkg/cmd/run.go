package cmd

import (
	"log"

	"github.com/sadasystems/gcsb/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var (
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Execute a load test",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.NewConfig(viper.GetViper())
			if err != nil {
				log.Fatalf("unable to parse configuration: %s", err.Error())
			}

			err = cfg.Validate()
			if err != nil {
				log.Fatalf("unable to validate configuration %s", err.Error())
			}
		},
	}
)
