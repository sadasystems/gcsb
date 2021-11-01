package operation

import (
	"math/rand"
	"time"

	"github.com/sadasystems/gcsb/pkg/config"
	"github.com/sadasystems/gcsb/pkg/generator/selector"
)

type (
	Operation int
)

const (
	READ Operation = 1 + iota
	WRITE
)

func NewOperationSelector(cfg *config.Config) (selector.Selector, error) {
	return selector.NewWeightedRandomSelector(
		rand.New(rand.NewSource(time.Now().UnixNano())),
		selector.NewWeightedChoice(READ, uint(cfg.Operations.Read)),
		selector.NewWeightedChoice(WRITE, uint(cfg.Operations.Write)),
	)
}
