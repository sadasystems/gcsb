package cmd

import (
	"fmt"
	"log"

	"github.com/sadasystems/gcsb/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

func init() {
	configDumpCmd.Flags().BoolVarP(&configDumpValidate, "validate", "v", false, "Validate the configuration")

	configCmd.AddCommand(configDumpCmd, configInitCmd)
	rootCmd.AddCommand(configCmd)
}

var (
	configDumpValidate bool // Validate the config before dumping it?

	configCmd = &cobra.Command{
		Use:   "config",
		Short: "Configuration related commands",
		Long:  ``,
	}

	configDumpCmd = &cobra.Command{
		Use:   "dump",
		Short: "Dump the configuration",
		Long:  `Used to help test the configuration package to make sure values, flags, and env variables are parsed correclty`,
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.NewConfig(viper.GetViper())
			if err != nil {
				log.Fatalf("unable to parse configuration: %s", err.Error())
			}

			if configDumpValidate {
				err = cfg.Validate()
				if err != nil {
					log.Fatalf("unable to validate configuration %s", err.Error())
				}
			}

			prettyPrint(cfg)
		},
	}

	configInitCmd = &cobra.Command{
		Use:   "init",
		Short: "initialize a new configuration file",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			v := viper.GetViper()
			cfg, _ := config.NewConfig(v)

			cfg.Project = "YOUR-PROJECT-ID"
			cfg.Instance = "YOUR-SPANNER-INSTANCE"
			cfg.Database = "YOUR-SPANNER-DATABASE"
			cfg.Tables = nil

			bs, err := yaml.Marshal(&cfg)
			if err != nil {
				log.Fatalf("unable to marshal config to YAML: %v", err)
			}
			fmt.Println(string(bs))
		},
	}
)
