package workload

import (
    "math/rand"
)

var rng = rand.New(rand.NewSource(1009)) // seed

// RandomBytes returns n deterministic pseudo-random bytes.
func RandomBytes(n int) []byte {
	b := make([]byte, n)
    for i := 0; i < n; i++ {
        b[i] = byte(rng.Intn(256))
    }
    return b
}