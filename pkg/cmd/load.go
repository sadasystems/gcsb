package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(loadCmd)
}

var (
	loadCmd = &cobra.Command{
		Use:   "load",
		Short: "Load a table with data",
		Long:  ``,
		Run:   doLoadCmd,
	}
)

func doLoadCmd(cmd *cobra.Command, args []string) {
	log.Println("Load test Load")
}
