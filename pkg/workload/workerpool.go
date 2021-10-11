package workload

import "github.com/sadasystems/gcsb/pkg/workload/pool"

var (
	// Assert that WorkerPool implements Workload
	_ Workload = (*WorkerPool)(nil)
)

type (
	WorkerPool struct { // Implement Workload
		Pool *pool.Pool
	}
)

// NewPoolWorkload initializes a "worker pool" type workload
func NewPoolWorkload(cfg WorkloadConfig) (Workload, error) {
	w := &WorkerPool{}
	w.Pool = pool.NewPool(pool.PoolConfig{
		Workers:        cfg.Config.Threads,
		BufferInput:    true,
		InputBufferLen: 100, // TODO: don't hardcode this
	})

	return w, nil
}

func (w *WorkerPool) Initialize() error {
	w.Pool.Start()

	return nil
}

func (w *WorkerPool) Load(tableName string) error {

	return nil
}

func (w *WorkerPool) Run() error {
	return nil
}
