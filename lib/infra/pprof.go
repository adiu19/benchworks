package infra

import (
    "os"
    "runtime/pprof"
)

// StartCPUProfile starts a CPU profile and returns a stop function.
func StartCPUProfile(path string) func() {
    f, err := os.Create(path)
    if err != nil {
        panic(err)
    }
    if err := pprof.StartCPUProfile(f); err != nil {
        panic(err)
    }
    return func() {
        pprof.StopCPUProfile()
        f.Close()
    }
}