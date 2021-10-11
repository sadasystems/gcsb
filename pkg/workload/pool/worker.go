package pool

type (
	// Worker is a goroutine that listens for work and processes incoming requests (Jobs)
	Worker struct {
		ID            int
		WorkerChannel chan chan Job
		JobOutput     chan Job // Job Output Channel
		JobInput      chan Job // Job input channel
		End           chan bool
	}
)

// Start worker
func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerChannel <- w.JobInput
			select {
			case job := <-w.JobInput:
				job.Execute()
			case <-w.End:
				return
			}
		}
	}()
}

// Stop worker
func (w *Worker) Stop() {
	w.End <- true
}
