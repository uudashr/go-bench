package bench_test

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkStringConvert_Itoa(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = strconv.Itoa(i)
	}
}

func BenchmarkStringConvert_Sprintf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint(i)
	}
}
