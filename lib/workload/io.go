package workload

import "time"

func FakeIO(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}