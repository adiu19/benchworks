//go:build cgo

package workload

import "time"

/*
#include <unistd.h>

// Sleep for ms milliseconds using usleep.
// usleep takes microseconds.
static void csleep_ms(int ms) {
    usleep((useconds_t)ms * 1000);
}
*/
import "C"

func FakeIO(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func CGOSleepMs(ms int) {
	C.csleep_ms(C.int(ms))
}
