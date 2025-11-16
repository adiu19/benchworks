package infra

import "time"

func Time(fn func()) time.Duration {
	start := time.Now()
	fn()
	return time.Since(start)
}
