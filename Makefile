
EXPERIMENT ?= sched-gomaxprocs-cpu
EXP_PKG := ./experiments/$(EXPERIMENT)

GOMAXPROCS ?= 0


.PHONY: help
help:
	@echo "Usage:"
	@echo "  make run           EXPERIMENT=<name> [GOMAXPROCS=N]"
	@echo "  make bench         EXPERIMENT=<name> [GOMAXPROCS=N]"
	@echo "  make profile-cpu   EXPERIMENT=<name> [GOMAXPROCS=N]"
	@echo "  make profile-mem   EXPERIMENT=<name> [GOMAXPROCS=N]"
	@echo "  make trace         EXPERIMENT=<name> [GOMAXPROCS=N]"
	@echo "  make list-experiments"
	@echo ""
	@echo "Current EXPERIMENT: $(EXPERIMENT)"

.PHONY: list-experiments
list-experiments:
	@echo "--------------------------------------------------------------------------------------------- \n"
	@echo "----------------------------------- Available experiments ----------------------------------- \n"
	@echo "--------------------------------------------------------------------------------------------- \n"

	@find experiments -maxdepth 1 -mindepth 1 -type d -exec basename {} \;



.PHONY: run
run:
	@echo "Running $(EXPERIMENT) with GOMAXPROCS=$(GOMAXPROCS)..."
	GOMAXPROCS=$(GOMAXPROCS) go run $(EXP_PKG)

.PHONY: bench
bench:
	@echo "Benchmarking $(EXPERIMENT) with GOMAXPROCS=$(GOMAXPROCS)..."
	GOMAXPROCS=$(GOMAXPROCS) go test -run=^$$ -bench=. $(EXP_PKG)


.PHONY: profile-cpu
profile-cpu:
	@echo "CPU profiling $(EXPERIMENT) (cpu.pprof), GOMAXPROCS=$(GOMAXPROCS)..."
	GOMAXPROCS=$(GOMAXPROCS) go test -run=^$$ -bench=. -cpuprofile=cpu.pprof $(EXP_PKG)
	@echo "CPU profile written to cpu.pprof"
	@echo "Inspect with: go tool pprof cpu.pprof"

.PHONY: profile-mem
profile-mem:
	@echo "Heap profiling $(EXPERIMENT) (mem.pprof), GOMAXPROCS=$(GOMAXPROCS)..."
	GOMAXPROCS=$(GOMAXPROCS) go test -run=^$$ -bench=. -memprofile=mem.pprof $(EXP_PKG)
	@echo "Heap profile written to mem.pprof"
	@echo "Inspect with: go tool pprof mem.pprof"

.PHONY: trace
trace:
	@echo "Tracing $(EXPERIMENT) (trace.out), GOMAXPROCS=$(GOMAXPROCS)..."
	GOMAXPROCS=$(GOMAXPROCS) go test -run=^$$ -bench=. -trace=trace.out $(EXP_PKG)
	@echo "Trace written to trace.out"
	@echo "Inspect with: go tool trace trace.out"

# ---------- Misc ----------


.PHONY: tidy
tidy:
	go mod tidy