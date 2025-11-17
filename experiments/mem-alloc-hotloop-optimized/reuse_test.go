package main

import (
	"testing"
)

func BenchmarkAllocOptimized(b *testing.B) {
	buf := make([]byte, 1<<20) // allocate 1 MB once

	for i := 0; i < b.N; i++ {
		buf[0] = byte(i) // pretend work
	}
}
