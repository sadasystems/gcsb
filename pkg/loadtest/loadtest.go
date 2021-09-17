package loadtest

import (
	"github.com/sadasystems/gcsb/pkg/config"
	"github.com/sadasystems/gcsb/pkg/db"
	adminpb "google.golang.org/genproto/googleapis/spanner/admin/database/v1"
)

type LoadTest struct {
	database      adminpb.Database
	queryExecutor db.QueryExecutor
	config        config.GCSBConfig
}

func NewLoadTest(configPath string) (*LoadTest, error) {
	config, err := config.NewGCSBConfigFromPath(configPath)
	if err != nil {
		return nil, err
	}
	database, err := db.GetDatabase(*config)
	if err != nil {
		return nil, err
	}
	lt := &LoadTest{
		database: database,
		config:   *config,
	}
	return lt, nil
}
