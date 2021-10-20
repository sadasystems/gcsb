package cmd

import (
	"log"

	"github.com/sadasystems/gcsb/pkg/config"
	"github.com/sadasystems/gcsb/pkg/schema"
	"github.com/sadasystems/gcsb/pkg/workload"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	loadCmd.Flags().StringVarP(&loadTable, "table", "t", "", "Table name to load")
	loadCmd.Flags().IntVarP(&loadOperations, "operations", "o", 1000, "Number of records to load")
	viper.BindPFlag("operations.total", loadCmd.Flags().Lookup("operations"))

	rootCmd.AddCommand(loadCmd)
}

var (
	// Flags
	loadTable      string
	loadOperations int

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
			// s, err = schema.LoadSingleTableSchema(ctx, cfg, loadTable) // TODO: Should the load command support multiple targets?
			s, err = schema.LoadSchema(ctx, cfg)
			if err != nil {
				log.Fatalf("unable to infer schema: %s", err.Error())
			}

			// Get a constructor for a workload
			constructor, err := workload.GetWorkloadConstructor("NOTYETSUPPORTED")
			if err != nil {
				log.Fatalf("unable to get workload constructor: %s", err.Error())
			}

			// Create a workload
			log.Println("Creating workload")
			wl, err := constructor(workload.WorkloadConfig{
				Context: ctx,
				Config:  cfg,
				Schema:  s,
			})
			if err != nil {
				log.Fatalf("unable to create workload: %s", err.Error())
			}

			// Execute the load phase
			log.Println("Executing load phase")
			err = wl.Load(loadTable)
			if err != nil {
				log.Fatalf("unable to execute load operation: %s", err.Error())
			}
		},
	}
)
