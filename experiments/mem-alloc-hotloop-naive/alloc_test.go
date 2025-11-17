package main

import "testing"

func BenchmarkAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := make([]byte, 1<<20) // 1 MB allocation, every time
		_ = buf
	}
}
