package cmd

import (
	"log"
	"os"

	"github.com/rcrowley/go-metrics"
	"github.com/sadasystems/gcsb/pkg/config"
	"github.com/sadasystems/gcsb/pkg/schema"
	"github.com/sadasystems/gcsb/pkg/workload"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	flags := loadCmd.Flags()

	// flags.StringVarP(&loadTable, "table", "t", "", "Table name to load")
	flags.StringSliceVarP(&loadTables, "table", "t", []string{}, "Table name to load")
	flags.IntP("operations", "o", 1000, "Number of records to load")
	flags.Int("threads", 10, "Number of threads")
	flags.BoolVar(&loadDry, "dry", false, "Dry run. Print config and exit.")

	rootCmd.AddCommand(loadCmd)
}

var (
	// Flags
	loadDry    bool
	loadTables []string

	// Command
	loadCmd = &cobra.Command{
		Use:   "load",
		Short: "Load a table with data",
		Long:  ``,
		PreRun: func(cmd *cobra.Command, args []string) {
			flags := cmd.Flags()
			viper.BindPFlag("operations.total", flags.Lookup("operations"))
			viper.BindPFlag("threads", flags.Lookup("threads"))
		},
		Run: func(cmd *cobra.Command, args []string) {
			if len(loadTables) <= 0 {
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

			// Log the configuration
			logConfig(cfg)
			if loadDry {
				log.Println("Exiting (--dry)")
				os.Exit(0)
			}

			// Since we are in the load command, we don't intend to have a lot of READ sessions.
			// Overwrite pool.write_sessions to be 1.0
			cfg.Pool.WriteSessions = 1

			// Get metric registry
			registry := metrics.NewRegistry()

			// Generate a context with cancelation
			log.Println("Creating a context with cancelation")
			ctx, cancel := cfg.Context() // TODO: this is dumb.. be more creative

			// Listen for os signals and cancel the context if we receive them
			log.Println("Listening for OS signals")
			graceful(cancel)

			// Measure how long schema inference takes to run
			schemaTimer := metrics.GetOrRegisterTimer("schema.inference", registry)

			// Infer the table schema from the database
			log.Println("Infering schema from database")
			var s schema.Schema
			schemaTimer.Time(func() {
				s, err = schema.LoadSchema(ctx, cfg)
			})
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
				Context:        ctx,
				Config:         cfg,
				Schema:         s,
				MetricRegistry: registry,
			})
			if err != nil {
				log.Fatalf("unable to create workload: %s", err.Error())
			}

			// measure the run phase
			runTimer := metrics.GetOrRegisterTimer("run", registry)

			// Execute the load phase
			log.Println("Executing load phase")
			runTimer.Time(func() {
				err = wl.Load(loadTables)
			})
			if err != nil {
				log.Fatalf("unable to execute load operation: %s", err.Error())
			}

			summarizeMetricsAsciiTable(registry)
		},
	}
)
