package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/sadasystems/gcsb/pkg/config"
	"github.com/sadasystems/gcsb/pkg/schema"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	plumbingSchemaInferCmd.Flags().StringVarP(&plumbingSchemaInferTableName, "table", "t", "", "table name")

	plumbingConfigDumpCmd.Flags().BoolVarP(&plumbingConfigDumpCmdValidate, "validate", "v", false, "Validate the configuration")

	plumbingConfigCmd.AddCommand(plumbingConfigDumpCmd)
	plumbingSchemaCmd.AddCommand(plumbingSchemaInferCmd)
	plumbingCmd.AddCommand(plumbingConfigCmd, plumbingSchemaCmd)
	rootCmd.AddCommand(plumbingCmd)
}

var (
	plumbingConfigDumpCmdValidate bool   // Validate the configuration
	plumbingSchemaInferTableName  string // table name to inspect

	plumbingCmd = &cobra.Command{
		Use:    "plumbing",
		Short:  "Plumbing commands used during development",
		Long:   `These commands are not a part of --help messages. Test things here. `,
		Hidden: true,
	}

	plumbingConfigCmd = &cobra.Command{
		Use:   "config",
		Short: "Configuration related commands",
		Long:  ``,
	}

	plumbingConfigDumpCmd = &cobra.Command{
		Use:   "dump",
		Short: "Dump the configuration",
		Long:  `Used to help test the configuration package to make sure values and flags are parsed correclty`,
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.NewConfig(viper.GetViper())
			if err != nil {
				log.Fatalf("unable to parse configuration: %s", err.Error())
			}

			if plumbingConfigDumpCmdValidate {
				err = cfg.Validate()
				if err != nil {
					log.Fatalf("unable to validate configuration %s", err.Error())
				}
			}

			prettyPrint(cfg)
		},
	}

	plumbingSchemaCmd = &cobra.Command{
		Use:   "schema",
		Short: "Schema related commands",
		Long:  ``,
	}

	plumbingSchemaInferCmd = &cobra.Command{
		Use:   "infer",
		Short: "Connect to the database and infer the schema",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			cfg, err := config.NewConfig(viper.GetViper())
			if err != nil {
				log.Fatalf("unable to parse configuration: %s", err.Error())
			}

			if plumbingConfigDumpCmdValidate {
				err = cfg.Validate()
				if err != nil {
					log.Fatalf("unable to validate configuration %s", err.Error())
				}
			}

			var s schema.Schema
			if plumbingSchemaInferTableName == "" {
				s, err = schema.LoadSchema(ctx, cfg)
			} else {
				s, err = schema.LoadSingleTableSchema(ctx, cfg, plumbingSchemaInferTableName)
			}
			if err != nil {
				log.Fatalf("unable to load schema: %s", err.Error())
			}
			spew.Dump(s)
		},
	}
)

// I'm too lazy to format output of plumbing commands
func prettyPrint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
}
