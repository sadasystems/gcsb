package cmd

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/sadasystems/gcsb/pkg/config"
	"github.com/sadasystems/gcsb/pkg/schema"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	loadCmd.Flags().StringVarP(&loadTable, "table", "t", "", "Table name to load")

	rootCmd.AddCommand(loadCmd)
}

var (
	// Flags
	loadTable string

	// Command
	loadCmd = &cobra.Command{
		Use:   "load",
		Short: "Load a table with data",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			if loadTable == "" {
				log.Fatal("missing table name (-t)")
			}

			// Load configuration
			log.Println("Loading configuration")
			cfg, err := config.NewConfig(viper.GetViper())
			if err != nil {
				log.Fatalf("unable to parse configuration: %s", err.Error())
			}

			// Validate the configuration
			log.Println("Validating configuration")
			err = cfg.Validate()
			if err != nil {
				log.Fatalf("unable to validate configuration %s", err.Error())
			}

			// Generate a context with cancelation
			log.Println("Creating a context with cancelation")
			ctx, cancel := cfg.Context() // TODO: this is dumb.. be more creative

			// Listen for os signals and cancel the context if we receive them
			log.Println("Listening for OS signals")
			graceful(cancel)

			// Infer the table schema from the database
			log.Println("Infering schema from database")
			var s schema.Schema
			s, err = schema.LoadSingleTableSchema(ctx, cfg, loadTable) // TODO: Should the load command support multiple targets?
			if err != nil {
				log.Fatalf("unable to infer schema: %s", err.Error())
			}

			//

			spew.Dump(s)
		},
	}
)
