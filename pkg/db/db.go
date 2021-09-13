package db

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"

	database "cloud.google.com/go/spanner/admin/database/apiv1"
	"github.com/sadasystems/gcsb/pkg/config"
	adminpb "google.golang.org/genproto/googleapis/spanner/admin/database/v1"
	"gopkg.in/yaml.v2"
)

type Config config.GCSBConfig

var (
	CreateStatement string
	conf            Config
	ctx             context.Context
)

func CreateDatabase(ctx context.Context, db string) (*adminpb.Database, error) {
	matches := regexp.MustCompile("^(.*)/databases/(.*)$").FindStringSubmatch(db)
	if matches == nil || len(matches) != 3 {
		return nil, fmt.Errorf("Invalid database id %s", db)
	}

	adminClient, err := database.NewDatabaseAdminClient(ctx)
	if err != nil {
		return nil, err
	}
	defer adminClient.Close()

	CreateStatement := "CREATE DATABASE `" + matches[2] + "`"
	op, err := adminClient.CreateDatabase(ctx, &adminpb.CreateDatabaseRequest{
		Parent:          matches[1],
		CreateStatement: CreateStatement,
	})

	if err != nil {
		return nil, err
	}
	dbInstance, err := op.Wait(ctx)
	if err != nil {
		return nil, err
	}
	return dbInstance, nil
}

func (c *Config) DBName() string {
	return "projects/" + c.Project + "/instances/" + c.Instance + "/databases/" + c.Database
}

func (c *Config) ReadConfig(configPath string) error {
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) GetCreateStatements() []string {
	statements := []string{}
	for _, t := range c.Tables {
		statements = append(statements, "DROP TABLE "+t.Name)
		var stmt string = "CREATE TABLE " + t.Name + " ( ID INT64 NOT NULL"
		for _, c := range t.Columns {
			stmt += ",\n" + c.Name + " " + c.Type
		}
		stmt += ") PRIMARY KEY (ID)"
		statements = append(statements, stmt)
	}
	return statements
}

func getDatabase(configPath string) (adminpb.Database, error) {
	var ret adminpb.Database
	ctx = context.Background()
	err := conf.ReadConfig(configPath)
	if err != nil {
		return ret, err
	}
	adminClient, err := database.NewDatabaseAdminClient(ctx)
	if err != nil {
		return ret, err
	}
	db, _ := adminClient.GetDatabase(ctx, &adminpb.GetDatabaseRequest{Name: conf.DBName()})

	if db == nil {
		db, err = CreateDatabase(ctx, conf.DBName())
		if err != nil {
			return *db, err
		}
	}

	return *db, nil
}

func CreateTable(configPath string) (adminpb.Database, error) {
	ctx = context.Background()

	DB, _ := getDatabase(configPath)
	adminClient, _ := database.NewDatabaseAdminClient(ctx)

	statements := conf.GetCreateStatements()
	fmt.Println(statements)
	op, err := adminClient.UpdateDatabaseDdl(ctx, &adminpb.UpdateDatabaseDdlRequest{
		Database:   conf.DBName(),
		Statements: statements,
	})
	if err != nil {
		log.Println("Error creating table " + err.Error())
	}
	if err := op.Wait(ctx); err != nil {
		return DB, err
	}
	log.Println("Created Venues table in database " + conf.DBName() + "\n")
	return DB, err
}
