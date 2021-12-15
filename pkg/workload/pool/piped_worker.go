package pool

// Worker is a goroutine that listens for work and processes incoming requests
type PipedWorker struct {
	ID            int
	WorkerChannel chan chan Job

	WriteToOutChannel bool
	JobOutput         chan Job // Job Output Channel
	JobInput          chan Job // Job input channel
	End               chan bool
}

// Start worker
func (w *PipedWorker) Start() {
	go func() {
		for {
			w.WorkerChannel <- w.JobInput
			select {
			case job := <-w.JobInput:
				job.Execute()
				if w.WriteToOutChannel {
					w.JobOutput <- job
				}
			case <-w.End:
				return
			}
		}
	}()
}

// Stop worker
func (w *PipedWorker) Stop() {
	w.End <- true
}
