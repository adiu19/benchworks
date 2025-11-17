package main

import (
	"sync"
	"testing"
)

// a per-P object cache with a shared global pool as fallback
var pool = sync.Pool{
	New: func() any {
		return make([]byte, 1<<20) // 1 MB allocation, cost paid per P
	},
}

func BenchmarkAllocPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := pool.Get().([]byte) // the first call to Get() will allocate a new 1 MB object via New()
		buf[0] = byte(i)
		pool.Put(buf)
	}
}
