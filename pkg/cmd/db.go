package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/sadasystems/gcsb/pkg/db"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2/google"
)

func init() {
	createDbFlags := dbCreateDbCmd.Flags()
	createDbFlags.StringVarP(&instanceName, "instance", "i", "test", "spanner instance name")
	createDbFlags.StringVarP(&dbName, "db", "d", "test", "Database name")

	createTableFlags := dbCreateTableCmd.Flags()
	createTableFlags.StringVarP(&configPath, "config", "c", "gcsb.yaml", "yaml config path")

	dbCreateCmd.AddCommand(dbCreateDbCmd)
	dbCreateCmd.AddCommand(dbCreateTableCmd)
	dbCmd.AddCommand(dbCreateCmd)
	rootCmd.AddCommand(dbCmd)
}

var (
	instanceName string
	dbName       string
	configPath   string

	dbCmd = &cobra.Command{
		Use:   "db",
		Short: "DB CLI",
	}

	dbCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create something",
	}

	dbCreateDbCmd = &cobra.Command{
		Use:   "db",
		Short: "Create a database",
		Run:   doCreateDb,
	}

	dbCreateTableCmd = &cobra.Command{
		Use:   "table",
		Short: "Create a table",
		Run:   doCreateTable,
	}
)

func doCreateDb(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	var sb strings.Builder
	credentials, _ := google.FindDefaultCredentials(ctx)

	sb.WriteString("projects/")
	sb.WriteString(credentials.ProjectID)
	sb.WriteString("/instances/")
	sb.WriteString(instanceName)
	sb.WriteString("/databases/")
	sb.WriteString(dbName)

	_, err := db.CreateDatabase(ctx, sb.String())
	if err != nil {
		fmt.Println(err)
	}
}

func doCreateTable(cmd *cobra.Command, args []string) {
	c, err := db.CreateTable(configPath)
	if err != nil {
		log.Println(err)
	}
	log.Println(c)
}
