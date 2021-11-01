package cmd

import (
	"log"
	"time"

	"github.com/rcrowley/go-metrics"

	"github.com/sadasystems/gcsb/pkg/config"
	"github.com/sadasystems/gcsb/pkg/schema"
	"github.com/sadasystems/gcsb/pkg/workload"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	flags := runCmd.Flags()
	flags.StringVarP(&runTable, "table", "t", "", "Table name to load")

	flags.IntP("operations", "o", 1000, "Number of operations to perform")
	viper.BindPFlag("operations.total", flags.Lookup("operations"))

	flags.Int("threads", 10, "Number of threads")
	viper.BindPFlag("threads", flags.Lookup("threads"))

	flags.Int("num-conns", 10, "Number of spanner connections")
	viper.BindPFlag("num_conns", flags.Lookup("num-conns"))

	flags.IntP("reads", "r", 50, "Read weight")
	viper.BindPFlag("operations.read", flags.Lookup("reads"))

	flags.IntP("writes", "w", 50, "Write weight")
	viper.BindPFlag("operations.write", flags.Lookup("writes"))

	flags.Float64P("sample-size", "s", 10, "Percentage of table to sample")
	viper.BindPFlag("operations.sample_size", flags.Lookup("sample-size"))

	flags.Bool("read-stale", false, "Perform stale reads")
	viper.BindPFlag("operations.read_stale", flags.Lookup("read-stale"))

	flags.Duration("staleness", time.Duration(15*time.Second), "Exact staleness timestamp bound")
	viper.BindPFlag("operations.staleness", flags.Lookup("staleness"))

	rootCmd.AddCommand(runCmd)
}

var (
	// Flags
	runTable string

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

			// Execute the run phase
			log.Println("Executing run phase")
			runTimer.Time(func() {
				err = wl.Run(runTable)
			})
			if err != nil {
				log.Fatalf("unable to execute run operation: %s", err.Error())
			}

			summarizeMetricsAsciiTable(registry)
		},
	}
)
