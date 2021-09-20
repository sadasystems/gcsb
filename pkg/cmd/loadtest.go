package cmd

import (
	"github.com/sadasystems/gcsb/pkg/loadtest"
	"github.com/spf13/cobra"
)

func init() {
	doLoadTestFlags := loadtestLoadCmd.Flags()
	doLoadTestFlags.StringVarP(&path, "config", "c", "gcsb.yaml", "yaml config path")

	loadtestCmd.AddCommand(loadtestLoadCmd)
	rootCmd.AddCommand(loadtestCmd)
}

var (
	path string

	loadtestCmd = &cobra.Command{
		Use:   "loadtest",
		Short: "Load test commands",
		Long:  ``,
	}

	loadtestLoadCmd = &cobra.Command{
		Use:   "load",
		Short: "Load a table with data",
		Long:  ``,
		Run:   doLoadTestLoadCmd,
	}
)

func doLoadTestLoadCmd(cmd *cobra.Command, args []string) {
	lt, err := loadtest.NewLoadTest(path)
	if err != nil {
		panic(err)
	}
	lt.Execute()
}
