# Memory Allocation Hotloop (1 MB per iteration)

With GOMAXPROCS=1, all allocs and GC happens on a single P/M, producing a clean view of heap growth and GC triggers

The experiment is run using the generic Makefile driver:

```
make trace EXPERIMENT=mem-alloc-hotloop-naive
```
Below are the three runs demonstrating how Go’s GC responds under different GOGC settings.

## 1. GOGC=off - GC disabled completely

```
GOGC=off GOMAXPROCS=1 make trace EXPERIMENT=mem-alloc-hotloop-naive
```

### Expected behavior
- No GC cycles occur (NextGC = 0).
- The heap grows linearly and unbounded as each iteration allocates 1 MB.
- Trace shows this,
    - an empty GC actvitity
    - a steadily rising heap graph

### Interpretation
With GC disabled, Go will never reclaim memory, the process grows until OOM or termination. 

## 2. GOGC=10 - aggressive GC

```
GOGC=10 GOMAXPROCS=1 make trace EXPERIMENT=mem-alloc-hotloop-naive
```

### Expected behavior

- GC triggers after only 10% heap growth past the last collection. 
- Heap remains close to its minimum post-GC size.
- Trace shows this,
    - frequent GC activity
    - almost no heap growth between cycles

### Interpretation
At low GOGC values, Go keeps heap size minimal but spends significant CPU time on GC. 

## 3. Default (GOGC=100) - standard Go behavior

```
GOMAXPROCS=1 make trace EXPERIMENT=mem-alloc-hotloop-naive
```

### Expected behavior

- GC triggers when the heap grows 100% beyond the previous post-GC size (default behavior).
- Heap grows steadily as we allocate 1 MB at a time and drops when GC runs
- GC cycles run regularly but not continuously.
- Trace shows this,
    - heap increasing significantly before each GC
    - GC slices that correlate precisely with drops in the “allocated” graph