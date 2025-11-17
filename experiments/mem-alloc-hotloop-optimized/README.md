# Goal

How fast we can do heap allocation, which is especially useful for high-performance go applications avoid GC pressure.

In the naive version of this experiment, we had something like

```
b := make([]byte, 1024 * 1024) // 1 MB allocation
```

and this comes with several caveats.

1. fresh heap span
2. GC pressure
3. mutator assist if the go runtime decides the goroutine is allocating too fast too soon.


The goal is to achieve the same use-case with ZERO GC traffic and ZERO heap growth which should give us

1. stable latency
2. minimal GC
3. predictable throughput
