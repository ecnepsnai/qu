# DS

[![Go Report Card](https://goreportcard.com/badge/github.com/ecnepsnai/qu?style=flat-square)](https://goreportcard.com/report/github.com/ecnepsnai/qu)
[![Godoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/ecnepsnai/qu)
[![Releases](https://img.shields.io/github/release/ecnepsnai/qu/all.svg?style=flat-square)](https://github.com/ecnepsnai/qu/releases)
[![LICENSE](https://img.shields.io/github/license/ecnepsnai/qu.svg?style=flat-square)](https://github.com/ecnepsnai/qu/blob/master/LICENSE)

Package qu is a executor service in golang.

You add your jobs to a list, then run them in parallel with a configurable amount of concurrency.

# Usage

```golang
queue := &qu.Queue{}

job := func(payload interface{}) {
    // do something with your payload
}

// Add a job to the queue
queue.Add(job, "your payload")

// Run through the queue with 2 threads
queue.Run(2)
```
