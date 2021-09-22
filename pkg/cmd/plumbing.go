package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/sadasystems/gcsb/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	plumbingCmd.AddCommand(plumbingConfigCmd)

	rootCmd.AddCommand(plumbingCmd)
}

var (
	plumbingCmd = &cobra.Command{
		Use:    "plumbing",
		Short:  "Plumbing commands used during development",
		Long:   `These commands are not a part of --help messages. Test things here. `,
		Hidden: true,
	}

	plumbingConfigCmd = &cobra.Command{
		Use:   "config-dump",
		Short: "Dump the configuration",
		Long:  `Used to help test the configuration package to make sure values and flags are parsed correclty`,
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.NewConfig(viper.GetViper())
			if err != nil {
				log.Fatalf("unable to parse configuration: %s", err.Error())
			}

			prettyPrint(cfg)
		},
	}
)

// I'm too lazy to format output of plumbing commands
func prettyPrint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
}
