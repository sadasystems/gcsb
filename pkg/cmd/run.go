package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var (
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Execute a load test",
		Long:  ``,
		Run:   doRunCmd,
	}
)

func doRunCmd(cmd *cobra.Command, args []string) {
	log.Println("Load test Run")
}
