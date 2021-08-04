/*
Package qu is a simple executor service. You add jobs to a queue, then run them concurrently with a configurable
amount of concurrency.
*/
package qu

import (
	"sync"
	"time"

	"github.com/ecnepsnai/qu/atomic"
)

// Queue describes a queue of jobs
type Queue struct {
	Done     bool
	jobs     []func(payload interface{})
	payloads []interface{}
}

// Add will add a new job to the queue. When the job is run it will be called with the value of payload.
// The job will not be invoked until queue.Run is called.
func (q *Queue) Add(job func(payload interface{}), payload interface{}) {
	q.jobs = append(q.jobs, job)
	q.payloads = append(q.payloads, payload)
}

// Run will begin to execute all of the jobs in the queue, running each job concurrently up-to the specified number of
// threads. Run will block until all jobs have completed. After this, the Done property on the queue will be true.
//
// Jobs may not be executed in the same order that they were added. If any jobs panics, the panic will bubble up to
// here.
func (q *Queue) Run(threads int) {
	runningJobs := atomic.NewInteger(0)
	remainingJobs := len(q.jobs)
	wg := &sync.WaitGroup{}

	for remainingJobs > 0 {
		if runningJobs.Get() < threads {
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

	q.Done = true
}
