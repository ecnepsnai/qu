package qu_test

import (
	"testing"
	"time"

	"github.com/ecnepsnai/qu"
)

func TestQueueMoreJobsThanThreads(t *testing.T) {
	queue := &qu.Queue{}

	job := func(payload interface{}) {
		time.Sleep(1 * time.Millisecond)
	}

	i := 0
	for i < 50 {
		queue.Add(job, i)
		i++
	}

	queue.Run(2)
}

func TestQueueMoreThreadsThanJobs(t *testing.T) {
	queue := &qu.Queue{}

	job := func(payload interface{}) {
		time.Sleep(1 * time.Millisecond)
	}

	i := 0
	for i < 5 {
		queue.Add(job, i)
		i++
	}

	queue.Run(50)
}

func TestQueueCheckLater(t *testing.T) {
	queue := &qu.Queue{}

	job := func(payload interface{}) {
		time.Sleep(1 * time.Millisecond)
	}

	i := 0
	for i < 5 {
		queue.Add(job, i)
		i++
	}

	go queue.Run(50)
	time.Sleep(10 * time.Millisecond)
	if !queue.Done {
		t.Errorf("Queue not done")
	}
}
