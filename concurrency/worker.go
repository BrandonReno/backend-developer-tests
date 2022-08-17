package concurrency

import "log"

type Worker struct {
	ID              int              // used to differentiate workers
	DispatchChannel chan chan func() // Channel of channel of referenced jobs, shared between the workers and dispatcher,
	WorkerChannel   chan func()      // Channel of job reference, personal to each worker
	Log             *log.Logger      // logger instance to log results
}

func (w *Worker) Start() {
	go func() {
		for {
			w.DispatchChannel <- w.WorkerChannel // place the worker channel on the dispatch channel
			select {
			case f := <-w.WorkerChannel: // In the case that dispatcher send a job
				w.Log.Printf("worker %d recieved task", w.ID)
				f()
				w.Log.Printf("worker %d completed task", w.ID)
			}
		}
	}()
}
