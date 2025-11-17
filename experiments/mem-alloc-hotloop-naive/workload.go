package main

func main() {
	const N = 1_000_000
	for i := 0; i < N; i++ {
		b := make([]byte, 1024)
		_ = b
	}
}
