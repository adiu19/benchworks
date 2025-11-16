package main

import (
	"benchworks/lib/workload"
	"os"
	"runtime"
	"sync"
	"syscall"
)

type Config struct {
	NumGoroutines  int
	SleepPerOpInMs int
	CpuWorkers     int
	CpuIters       int
}

func FloodSysCalls(cfg Config) {
	gomax := runtime.GOMAXPROCS(0)
	_ = gomax

	var wg sync.WaitGroup
	wg.Add(cfg.CpuWorkers)
	for i := 0; i < cfg.CpuWorkers; i++ {
		go func() {
			defer wg.Done()
			workload.FakeCPU(cfg.CpuIters)
		}()
	}

	wg.Add(cfg.NumGoroutines)
	for i := 0; i < cfg.NumGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j <= 3; j++ {
				workload.FakeCPU(10_000)
			}
			workload.CGOSleepMs(cfg.SleepPerOpInMs)

		}()
	}
	wg.Wait()

}

func FloodSyscallReads(cfg Config) {
	gomax := runtime.GOMAXPROCS(0)
	_ = gomax

	// Create a pipe and never write to it
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	// We keep both open during the trace so the read blocks forever.
	defer r.Close()
	defer w.Close()

	var wg sync.WaitGroup

	wg.Add(cfg.CpuWorkers)
	for i := 0; i < cfg.CpuWorkers; i++ {
		go func() {
			defer wg.Done()
			workload.FakeCPU(cfg.CpuIters)
		}()
	}

	// Blockers: each does a syscall.Read that never returns.
	fd := int(r.Fd())
	for i := 0; i < cfg.NumGoroutines; i++ {
		go func() {
			runtime.LockOSThread()
			defer runtime.UnlockOSThread()
			buf := make([]byte, 1)
			// This will block in the kernel; runtime should treat it as a syscall
			_, _ = syscall.Read(fd, buf)
		}()
	}

	wg.Wait()

	workload.FakeIO(500)
}
