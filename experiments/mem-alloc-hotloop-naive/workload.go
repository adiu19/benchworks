package main

func main() {
	const N = 1_000_000
	for i := 0; i < N; i++ {
		b := make([]byte, 1<<10) // 1 KB allocation
		_ = b
	}
}
