package main

import (
	"runtime"
	"testing"
	"time"
)

func BenchmarkFloodSyscalls(b *testing.B) {
	gomax := runtime.GOMAXPROCS(0)
	b.Logf("BenchmarkFloodSyscalls: GOMAXPROCS=%d", gomax)

	cfg := Config{
		NumGoroutines:  10_000,
		SleepPerOpInMs: 5,
		CpuWorkers:     2,
		CpuIters:       1_000_000,
	}

	for i := 0; i < b.N; i++ {
		FloodSyscalls(cfg)
	}
}
