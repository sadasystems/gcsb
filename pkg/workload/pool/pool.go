package pool

type (
	PoolConfig struct {
		Workers        int
		BufferInput    bool
		InputBufferLen int
	}

	Pool struct {
		Config        PoolConfig
		WorkInput     chan Job // For client to send work to
		End           chan bool
		Workers       []Worker
		WorkerChannel chan chan Job
	}
)

// NewPool will return a configured Pool. You must call Start() to begin processing jobs
func NewPool(cfg PoolConfig) *Pool {

	p := &Pool{
		Config:        cfg,
		End:           make(chan bool),     // channel to spin down workers
		WorkerChannel: make(chan chan Job), // WorkerChannel is a channel of work worker channels (lol)
	}

	var input chan Job
	if cfg.BufferInput {
		input = make(chan Job, cfg.InputBufferLen) // channel to recieve work
	} else {
		input = make(chan Job)
	}
	p.WorkInput = input

	var i int
	for i < cfg.Workers {
		i++
		worker := Worker{
			ID:            i,
			JobInput:      make(chan Job),
			WorkerChannel: p.WorkerChannel,
			End:           make(chan bool),
		}

		worker.Start()
		p.Workers = append(p.Workers, worker) // store worker
	}

	return p
}

// Start will start the job dispatcher
func (p *Pool) Start() {
	// start collector
	go func() {
		for {
			select {
			case <-p.End:
				for _, w := range p.Workers {
					w.Stop() // stop worker
				}
				return
			case work := <-p.WorkInput:
				worker := <-p.WorkerChannel // wait for available channel
				worker <- work              // dispatch work to worker
			}
		}
	}()
}

// Submit is syntax sugar for pushing a job into the input channel
func (p *Pool) Submit(job Job) {
	p.WorkInput <- job
}

// Stop workers and then the workerpool dispatcher
func (p *Pool) Stop() {
	p.End <- true
}
