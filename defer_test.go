package bench_test

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func doNoDefer(t *int) {
	fmt.Fprintln(ioutil.Discard, "Main logic")

	fmt.Fprintln(ioutil.Discard, "Closure")
}

func doNoDeferClosureFunc(t *int) {
	fmt.Fprintln(ioutil.Discard, "Main logic")

	func() {
		fmt.Fprintln(ioutil.Discard, "Closure")
	}()
}

func doDefer(t *int) {
	defer func() {
		*t++
		fmt.Fprintln(ioutil.Discard, "Closure")
	}()
	fmt.Fprintln(ioutil.Discard, "Main logic")
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
