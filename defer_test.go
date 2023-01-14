package bench_test

import (
	"fmt"
	"io"
	"testing"
)

func doNoDefer(t *int) {
	fmt.Fprintln(io.Discard, "Main logic")

	fmt.Fprintln(io.Discard, "Closure")
}

func doNoDeferClosureFunc(t *int) {
	fmt.Fprintln(io.Discard, "Main logic")

	func() {
		fmt.Fprintln(io.Discard, "Closure")
	}()
}

func doDefer(t *int) {
	defer func() {
		*t++
		fmt.Fprintln(io.Discard, "Closure")
	}()
	fmt.Fprintln(io.Discard, "Main logic")
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

func BenchmarkDeferNo_ClosureFunc(b *testing.B) {
	t := 0
	for i := 0; i < b.N; i++ {
		doNoDeferClosureFunc(&t)
	}
}
