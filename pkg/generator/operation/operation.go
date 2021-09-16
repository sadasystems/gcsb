package operation

import (
	"github.com/sadasystems/gcsb/pkg/config"
	"github.com/sadasystems/gcsb/pkg/generator/selector"
	"math/rand"
	"time"
)

/*
 * Here we will be doing an 'operation generator'
 *
 * An operation generator returns a SQL operation (READ|WRITE)
 * with a configurable 'weighted probability'
 *
 * For example, I want to generate operations with 80% reads and 20% writes
 */

type (
	Operation int
)

const (
	READ Operation = 1 + iota
	WRITE
)

func NewOperationGenerator(operations *config.TableConfigOperations) selector.Selector {
	rs := rand.New(rand.NewSource(time.Now().UnixNano()))
	var reads, writes uint
	if operations == nil {
		reads = 50
		writes = 50
	} else {
		reads = operations.Read
		writes = operations.Write
	}

	wrs, _ := selector.NewWeightedRandomSelector(rs, selector.NewWeightedChoice(READ, reads), selector.NewWeightedChoice(WRITE, writes))

	return wrs
}
