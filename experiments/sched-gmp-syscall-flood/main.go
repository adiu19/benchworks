package main

import (
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

func main() {
	gomax := runtime.GOMAXPROCS(0)
	log.Printf("GOMAXPROCS=%d", gomax)

	f, err := os.Create("trace-sched-gmp-syscall-flood.out")
	if err != nil {
		log.Fatalf("failed to create trace file: %v", err)
	}
	defer f.Close()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	cfg := Config{
		NumGoroutines:  64,         // G >> P. In trace, we should see tons of goroutines mostly blocked on syscall
		SleepPerOpInMs: 5000,       // enough time to keep the goroutines blocked on syscall, but not too long that the trace window takes forever
		CpuWorkers:     4,          // Ps always have something to schedule
		CpuIters:       20_000_000, // enough work to keep CPU workers running for the whole trace window
	}

	FloodSyscallReads(cfg)

	time.Sleep(100 * time.Millisecond)

}
