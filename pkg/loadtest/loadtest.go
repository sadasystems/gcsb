package loadtest

import (
	"github.com/sadasystems/gcsb/pkg/config"
	"github.com/sadasystems/gcsb/pkg/db"
)

type LoadTest struct {
	database      db.DB
	queryExecutor db.QueryExecutor
	config        config.GCSBConfig
}

func NewLoadTest(configPath string) (*LoadTest, error) {
	config, err := config.NewGCSBConfigFromPath(configPath)
	if err != nil {
		return nil, err
	}
	database, err := db.NewDB(*config)
	if err != nil {
		return nil, err
	}
	queryExecutor, err := db.NewQueryExecutor(*config)
	if err != nil {
		return nil, err
	}
	lt := &LoadTest{
		database:      *database,
		config:        *config,
		queryExecutor: *queryExecutor,
	}
	return lt, nil
}

func (lt *LoadTest) Execute() {
	lt.database.GetDatabase()
	for _, table := range lt.config.Tables {
		lt.queryExecutor.Execute(table)
	}
}
