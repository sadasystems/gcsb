package cmd

import (
	"log"

	"github.com/rcrowley/go-metrics"

	"github.com/sadasystems/gcsb/pkg/config"
	"github.com/sadasystems/gcsb/pkg/schema"
	"github.com/sadasystems/gcsb/pkg/workload"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	runCmd.Flags().StringVarP(&runTable, "table", "t", "", "Table name to load")
	runCmd.Flags().IntVarP(&runOperations, "operations", "o", 1000, "Number of operations to perform")
	viper.BindPFlag("operations.total", runCmd.Flags().Lookup("operations"))

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
			// l := log.Default()
			// metrics.WriteOnce(registry, l.Writer())

			// b := &bytes.Buffer{}
			// metrics.WriteJSONOnce(registry, b)
			// log.Println(b.String())
		},
	}
)
