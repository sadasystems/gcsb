package cmd

import (
	"github.com/sadasystems/gcsb/pkg/config"
	"github.com/sadasystems/gcsb/pkg/db"
	"github.com/spf13/cobra"
)

func init() {
	doCreateFlags := createCmd.Flags()
	doCreateFlags.StringVarP(&configPath, "config", "c", "gcsb.yaml", "yaml config path")

	dbCmd.AddCommand(createCmd)
	rootCmd.AddCommand(dbCmd)

}

var (
	configPath string

	dbCmd = &cobra.Command{
		Use:   "db",
		Short: "DB CLI",
	}
	createCmd = &cobra.Command{
		Use: "create",
		Run: doCreate,
	}
)

func doCreate(cmd *cobra.Command, args []string) {
	cfg, err := config.NewGCSBConfigFromPath(configPath)
	if err != nil {
		panic(err)
	}
	DB, err := db.NewDB(*cfg)
	if err != nil {
		panic(err)
	}
	err = DB.GetDatabase()
	if err != nil {
		panic(err)
	}
}
