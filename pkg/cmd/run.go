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
	runCmd.Flags().StringVarP(&runTable, "table", "t", "", "Table name to load")
	runCmd.Flags().IntVarP(&runOperations, "operations", "o", 1000, "Number of operations to perform")
	viper.BindPFlag("operations.total", loadCmd.Flags().Lookup("operations"))

	rootCmd.AddCommand(runCmd)
}

var (
	// Flags
	runTable      string
	runOperations int

	// Command
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Execute a load test",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			if runTable == "" {
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

			// Execute the run phase
			log.Println("Executing run phase")
			err = wl.Run(runTable)
			if err != nil {
				log.Fatalf("unable to execute run operation: %s", err.Error())
			}
		},
	}
)
