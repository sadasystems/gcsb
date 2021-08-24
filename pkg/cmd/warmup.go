package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(warmUpCmd)
}

var (
	warmUpCmd = &cobra.Command{
		Use:   "warmup",
		Short: "Warmup a Spanner table",
		Long: `
THIS
IS 
LONG
		`,
		Run: doWarmupCmd,
	}
)

func doWarmupCmd(cmd *cobra.Command, args []string) {
	log.Println("Warmup Command")
}
