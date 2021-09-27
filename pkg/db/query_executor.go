package db

import (
	"cloud.google.com/go/spanner"
	"context"
	"fmt"
	"github.com/sadasystems/gcsb/pkg/config"
)

type (
	QueryExecutor struct {
		config        config.GCSBConfig
		queryBuilders map[string]*QueryBuilder
		dbClient      *spanner.Client
		ctx           context.Context
	}
)

func NewQueryExecutor(config config.GCSBConfig) (*QueryExecutor, error) {

	ctx := context.Background()

	client, err := spanner.NewClient(ctx, config.DBName())
	if err != nil {
		return nil, err
	}
	q := &QueryExecutor{config: config, queryBuilders: make(map[string]*QueryBuilder), dbClient: client}

	for _, table := range config.Tables {
		q.queryBuilders[table.Name] = NewQueryBuilder(table)
	}

	return q, nil
}

func (qe *QueryExecutor) Execute(table config.TableConfigTable) {

	for i := 0; i < table.RowCount; i++ {
		qe.Next(table)
	}
}

func (qe *QueryExecutor) Next(table config.TableConfigTable) {
	ctx := context.Background()
	tx, err := qe.dbClient.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		stmt := qe.queryBuilders[table.Name].Next()
		fmt.Println(stmt)
		txn.Query(ctx, stmt)
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(tx.Clock())
}
