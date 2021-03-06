package qu_test

import (
	"fmt"
	"time"

	"github.com/ecnepsnai/qu"
)

func ExampleQueue_Run() {
	queue := &qu.Queue{}

	job := func(payload interface{}) {
		i := payload.(int)
		fmt.Printf("Job %d\n", i)
		time.Sleep(1 * time.Millisecond)
	}

	// Add 50 jobs to the queue
	i := 0
	for i < 50 {
		queue.Add(job, i)
		i++
	}

	// Go through those 50 jobs across 2 threads
	queue.Run(2)
}
