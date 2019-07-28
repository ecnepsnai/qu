/*
Package qu is a executor service in golang.

You add your jobs to a list, then run them in parallel with a configurable amount of concurrency.
*/
package qu

import (
	"sync"
	"time"

	"github.com/ecnepsnai/qu/atomic"
)

// Queue describes a queue of jobs
type Queue struct {
	jobs     []func(payload interface{})
	payloads []interface{}
}

// Add add a new job to the queue. Will be called with the given payload.
func (q *Queue) Add(job func(payload interface{}), payload interface{}) {
	q.jobs = append(q.jobs, job)
	q.payloads = append(q.payloads, payload)
}

// Run execute all jobs in the queue with the specified number of threads.
// Will block until all jobs have completed.
func (q *Queue) Run(threads int) {
	runningJobs := atomic.NewInteger(0)
	remainingJobs := len(q.jobs)

	for remainingJobs > 0 {
		if runningJobs.Get() < threads {
			wg := &sync.WaitGroup{}
			wg.Add(1)
			runningJobs.IncrementAndGet()

			job := q.jobs[0]
			payload := q.payloads[0]
			q.jobs = append(q.jobs[:0], q.jobs[1:]...)
			q.payloads = append(q.payloads[:0], q.payloads[1:]...)
			remainingJobs = len(q.jobs)

			go func() {
				job(payload)
				runningJobs.DecrementAndGet()
				wg.Done()
			}()

			if remainingJobs == 0 {
				wg.Wait()
			}
		} else {
			time.Sleep(1 * time.Millisecond)
		}
	}
}
