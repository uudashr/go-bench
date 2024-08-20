package bench_test

import (
	"fmt"
	"testing"
)

var sampleID = "12345"

func BenchmarkSimpleStringConcat_Sprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("/rides/%s", sampleID)
	}
}

func BenchmarkSimpleStringConcat_Concat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = "/rides/" + sampleID
	}
}
