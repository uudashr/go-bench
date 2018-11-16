package bench_test

import "testing"

func doNoDefer(t *int) {
	func() {
		*t++
	}()
}
func doDefer(t *int) {
	defer func() {
		*t++
	}()
}

func BenchmarkDeferYes(b *testing.B) {
	t := 0
	for i := 0; i < b.N; i++ {
		doDefer(&t)
	}
}
func BenchmarkDeferNo(b *testing.B) {
	t := 0
	for i := 0; i < b.N; i++ {
		doNoDefer(&t)
	}
}
