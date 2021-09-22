package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var (
	cfgFile string

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "gcsb",
		Short: "Like YCSB but for spanner",
		Long:  ``,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./gcsb.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in "." directory with name gcsb.yaml
		viper.AddConfigPath(".")
		viper.SetConfigName("gcsb")
	}

	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvPrefix("GCSB")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// err := viper.ReadInConfig()
	// if err != nil {
	// 	log.Fatalf("error reading config: %s", err.Error())
	// }

	viper.ReadInConfig() // Ignore errors here. We don't want to exit if no config file is found
}
