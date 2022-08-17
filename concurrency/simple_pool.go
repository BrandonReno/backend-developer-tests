// Package concurrency implements worker pool interfaces, one simple and one a
// bit more complex.
package concurrency

import "log"

var AvailableWorkers = make(chan chan func())

// SimplePool is a simple worker pool that does not support cancellation or
// closing. All functions are safe to call from multiple goroutines.
type SimplePool interface {
	// Submit a task to be executed asynchronously. This function will return as
	// soon as the task is submitted. If the pool does not have an available slot
	// for the task, this blocks until it can submit.
	Submit(func())
}

type Collector struct {
	Work chan func() // Jobs come in through this channel
}

func (c *Collector) Submit(f func()) {
	c.Work <- f
}

// NewSimplePool creates a new SimplePool that only allows the given maximum
// concurrent tasks to run at any one time. maxConcurrent must be greater than
// zero.
func NewSimplePool(maxConcurrent int) SimplePool {
	incomingJobs := make(chan func(), maxConcurrent) // set buffer to maxconcurrent so the system blocks if more
	collector := Collector{Work: incomingJobs}
	var workers []Worker
	l := log.Default()
	for i := 0; i < maxConcurrent; i++ { // add all the workers to match maxconcurrent
		w := Worker{
			ID:              i,
			DispatchChannel: AvailableWorkers,
			WorkerChannel:   make(chan func()),
			Log:             l,
		}

		workers = append(workers, w)

		w.Start() // start all workers
	}

	// Start the dispatcher
	go func() { // listen forever for incoming tasks until an end is recieved
		for job := range incomingJobs {
			nextAvailableWorker := <-AvailableWorkers // check who the next available worker in the Available worker channel
			nextAvailableWorker <- job                // send that next available worker the job to process
		}
	}()
	return &collector
}
