package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	loadtestCmd.AddCommand(loadtestLoadCmd, loadtestRunCmd)
	rootCmd.AddCommand(loadtestCmd)
}

var (
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

	loadtestRunCmd = &cobra.Command{
		Use:   "run",
		Short: "Execute a load test",
		Long:  ``,
		Run:   doLoadTestRunCmd,
	}
)

func doLoadTestLoadCmd(cmd *cobra.Command, args []string) {
	log.Println("Load test Load")
}

func doLoadTestRunCmd(cmd *cobra.Command, args []string) {
	log.Println("Load test Run")
}
