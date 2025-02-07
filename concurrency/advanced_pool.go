package concurrency

import (
	"context"
	"errors"
	"log"
)

var AdvancedAvailableWorkers = make(chan chan func(context.Context))

// ErrPoolClosed is returned from AdvancedPool.Submit when the pool is closed
// before submission can be sent.
var ErrPoolClosed = errors.New("pool closed")

// AdvancedPool is a more advanced worker pool that supports cancelling the
// submission and closing the pool. All functions are safe to call from multiple
// goroutines.
type AdvancedPool interface {
	// Submit submits the given task to the pool, blocking until a slot becomes
	// available or the context is closed. The given context and its lifetime only
	// affects this function and is not the context passed to the callback. If the
	// context is closed before a slot becomes available, the context error is
	// returned. If the pool is closed before a slot becomes available,
	// ErrPoolClosed is returned. Otherwise the task is submitted to the pool and
	// no error is returned. The context passed to the callback will be closed
	// when the pool is closed.
	Submit(context.Context, func(context.Context)) error

	// Close closes the pool and waits until all submitted tasks have completed
	// before returning. If the pool is already closed, ErrPoolClosed is returned.
	// If the given context is closed before all tasks have finished, the context
	// error is returned. Otherwise, no error is returned.
	Close(context.Context) error
}

type AdvancedCollector struct {
	Work chan func(context.Context) // Jobs come in through this channel
	End chan bool
}

func (c AdvancedCollector) Submit(ctx context.Context,f func(context.Context)) error {
	c.Work <- f
	return nil
}

func (c AdvancedCollector) Close(context.Context) error{
	c.End <- true
	return nil
}

// NewAdvancedPool creates a new AdvancedPool. maxSlots is the maximum total
// submitted tasks, running or waiting, that can be submitted before Submit
// blocks waiting for more room. maxConcurrent is the maximum tasks that can be
// running at any one time. An error is returned if maxSlots is less than
// maxConcurrent or if either value is not greater than zero.
func NewAdvancedPool(maxSlots, maxConcurrent int) (AdvancedPool, error) {
	if maxSlots < maxConcurrent || maxConcurrent < 0{
		return nil, errors.New("invalid configuration of maxSlots, maxConcurrent")
	}
	incomingJobs := make(chan func(context.Context), maxSlots)
	endWork := make(chan bool) 
	collector := AdvancedCollector{Work: incomingJobs, End: endWork}
	var workers []Worker
	l := log.Default()

	//Initialize the workers

	for i := 0; i < maxConcurrent; i++{
		w := Worker{
			ID: i,
			DispatchChannel: AdvancedAvailableWorkers,
			WorkerChannel: make(chan func()),
			End: make(chan bool),
			Log: l,
		}
		
		workers = append(workers, w)

		w.Start() //start all the workers
	}

	// Start the dispatcher
	go func(){
		for {
			select{
				case <-endWork:
					for _,w := range workers{
						l.Println("Worker channels closing.")
						w.Stop()
					}
					return
				case job := <-incomingJobs: //when an incoming job comes in
					nextAvailableWorker := <- AvailableWorkers // check who the next available worker in the Available worker channel
					nextAvailableWorker <- job // send that next available worker the job to process
			}

				

		}
	}()
	
	return collector, nil
}
