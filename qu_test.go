package qu_test

import (
	"testing"
	"time"

	"github.com/ecnepsnai/qu"
)

func TestQu(t *testing.T) {
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
