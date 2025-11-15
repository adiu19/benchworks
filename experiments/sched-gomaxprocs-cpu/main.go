package main

import (
	"flag"
	"fmt"
	"runtime"
	"sync"
	"time"

	"benchworks/lib/workload"
)

var (
	itersFlag   = flag.Int("iters", 50_000_000, "iterations per worker")
	workersFlag = flag.Int("workers", 0, "number of workers (0 = use NumCPU())")
)

func main() {
	flag.Parse()

	gomax := runtime.GOMAXPROCS(0)
	cpus := runtime.NumCPU()

	workers := *workersFlag
	if workers <= 0 {
		workers = cpus
	}

	fmt.Printf("sched-gomaxprocs-cpu\n")
	fmt.Printf("  GOMAXPROCS = %d\n", gomax)
	fmt.Printf("  NumCPU     = %d\n", cpus)
	fmt.Printf("  workers    = %d\n", workers)
	fmt.Printf("  iters/worker = %d\n", *itersFlag)

	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			workload.FakeCPU(*itersFlag)
		}()
	}
	wg.Wait()

	dur := time.Since(start)
	fmt.Printf("total duration:      %s\n", dur)
	fmt.Printf("approx per-worker:   %s\n", dur/time.Duration(workers))
}
