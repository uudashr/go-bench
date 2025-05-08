package bench_test

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkStringConvert_Itoa(b *testing.B) {
	for i := 0; b.Loop(); i++ {
		_ = strconv.Itoa(i)
	}
}

func BenchmarkStringConvert_SprintInt(b *testing.B) {
	for i := 0; b.Loop(); i++ {
		_ = fmt.Sprint(i)
	}
}

func BenchmarkStringConvert_SprintInt64(b *testing.B) {
	for i := int64(0); b.Loop(); i++ {
		_ = fmt.Sprint(i)
	}
}

func BenchmarkStringConvert_SprintUint64(b *testing.B) {
	for i := uint64(0); b.Loop(); i++ {
		_ = fmt.Sprint(i)
	}
}

func BenchmarkStringConvert_FormatInt_non64(b *testing.B) {
	for i := 0; b.Loop(); i++ {
		_ = strconv.FormatInt(int64(i), 10)
	}
}

func BenchmarkStringConvert_FormatInt_indirection(b *testing.B) {
	for i := 0; b.Loop(); i++ {
		_ = formatInt(i)
	}
}

func BenchmarkStringConvert_FormatInt(b *testing.B) {
	for i := int64(0); b.Loop(); i++ {
		_ = strconv.FormatInt(i, 10)
	}
}

func BenchmarkStringConvert_FormatUint(b *testing.B) {
	for i := uint64(0); b.Loop(); i++ {
		_ = strconv.FormatUint(i, 10)
	}
}

func formatInt(i int) string {
	return strconv.FormatInt(int64(i), 10)
}
