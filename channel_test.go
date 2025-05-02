package bench_test

import (
	"fmt"
	"testing"
)

func BenchmarkChannel_Close(b *testing.B) {
	for i := 0; b.Loop(); i++ {
		out := formatAsyncClose("Hello World %d", i)
		<-out
	}
}

func BenchmarkChannel_NoClose(b *testing.B) {
	for i := 0; b.Loop(); i++ {
		out := formatAsyncNoClose("Hello World %d", i)
		<-out
	}
}

func formatAsyncClose(format string, a ...interface{}) <-chan string {
	out := make(chan string)
	go func() {
		out <- fmt.Sprintf(format, a...)
		close(out)
	}()
	return out
}

func formatAsyncNoClose(format string, a ...interface{}) <-chan string {
	out := make(chan string)
	go func() {
		out <- fmt.Sprintf(format, a...)
	}()
	return out
}
