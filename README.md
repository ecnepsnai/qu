# QU

[![Go Report Card](https://goreportcard.com/badge/github.com/ecnepsnai/qu?style=flat-square)](https://goreportcard.com/report/github.com/ecnepsnai/qu)
[![Godoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/ecnepsnai/qu)
[![Releases](https://img.shields.io/github/release/ecnepsnai/qu/all.svg?style=flat-square)](https://github.com/ecnepsnai/qu/releases)
[![LICENSE](https://img.shields.io/github/license/ecnepsnai/qu.svg?style=flat-square)](https://github.com/ecnepsnai/qu/blob/master/LICENSE)

Package qu is a simple executor service. You add jobs to a queue, then run them concurrently with a configurable
amount of concurrency.

# Usage

```golang
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
```
