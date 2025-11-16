package workload

func FakeCPU(n int) {
	x := 0
	for i := 0; i < n; i++ {
		// A simple arithmetic chain to keep CPU busy
		x = x*1664525 + 1013904223
	}

	// check to make sure go compiler didn't optimize the loop away, i.e, make the result observable
	if x == 42 {
		println("impossible")
	}
}
