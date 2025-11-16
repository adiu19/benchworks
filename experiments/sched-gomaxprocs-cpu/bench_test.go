// experiments/sched-gomaxprocs-cpu/bench_test.go
package main

import (
	"runtime"
	"testing"

	"benchworks/lib/workload"
)

// BenchmarkCpuBurnSingle runs a CPU-bound loop in a single goroutine.
func BenchmarkCpuBurnSingle(b *testing.B) {
	gomax := runtime.GOMAXPROCS(0)
	b.Logf("BenchmarkCpuBurnSingle: GOMAXPROCS=%d", gomax)

	for i := 0; i < b.N; i++ {
		workload.FakeCPU(10_000_000)
	}
}

// BenchmarkCpuBurnParallel uses RunParallel to let the benchmark harness
func BenchmarkCpuBurnParallel(b *testing.B) {
	gomax := runtime.GOMAXPROCS(0)
	b.Logf("BenchmarkCpuBurnParallel: GOMAXPROCS=%d", gomax)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			workload.FakeCPU(10_000_000)
		}
	})
}
