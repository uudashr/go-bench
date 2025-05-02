package bench_test

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func BenchmarkStringConvert_Itoa(b *testing.B) {

	for i := 0; b.Loop(); i++ {
		_ = strconv.Itoa(i)
	}
}

func BenchmarkStringConvert_Sprint(b *testing.B) {

	for i := 0; b.Loop(); i++ {
		_ = fmt.Sprint(i)
	}
}

func BenchmarkStringConvert_Int64Itoa(b *testing.B) {

	for b.Loop() {
		_ = strconv.FormatUint(uint64(math.MaxInt64)+1, 10)
	}
}

func BenchmarkStringConvert_Int64Sprint(b *testing.B) {

	for b.Loop() {
		_ = fmt.Sprint(uint64(math.MaxInt64) + 1)
	}
}
