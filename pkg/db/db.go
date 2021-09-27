package db

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
	database "cloud.google.com/go/spanner/admin/database/apiv1"
	"github.com/sadasystems/gcsb/pkg/config"
	databasepb "google.golang.org/genproto/googleapis/spanner/admin/database/v1"
	"google.golang.org/grpc/codes"
)

type (
	DB struct {
		admin  *database.DatabaseAdminClient
		client *spanner.Client
		ctx    context.Context
		config config.GCSBConfig
	}
)

func NewDB(config config.GCSBConfig) (*DB, error) {
	ctx := context.Background()
	admin, err := database.NewDatabaseAdminClient(ctx)
	if err != nil {
		return nil, err
	}
	client, err := spanner.NewClient(ctx, config.DBName())
	if err != nil {
		return nil, err
	}
	db := &DB{
		admin:  admin,
		client: client,
		ctx:    ctx,
		config: config,
	}
	return db, nil
}

func (db *DB) CreateDatabase() error {
	_, err := db.admin.CreateDatabase(db.ctx, &databasepb.CreateDatabaseRequest{
		Parent:          db.config.ParentName(),
		CreateStatement: "CREATE DATABASE `" + db.config.Database + "`",
		ExtraStatements: db.config.GetCreateStatements(),
	})
	if err == nil {
		fmt.Println("Created database " + db.config.DBName())
	}
	return err
}

func (db *DB) GetDatabase() error {
	fmt.Println(db.config.DBName())
	_, err := db.admin.GetDatabase(db.ctx, &databasepb.GetDatabaseRequest{Name: db.config.DBName()})
	if err != nil && spanner.ErrCode(err) == codes.NotFound {
		fmt.Println("Database not found, creating")
		err = db.CreateDatabase()
	} else if err != nil {
		return err
	}
	return err
}
