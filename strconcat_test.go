package bench_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func BenchmarkStringConcat_Literal(b *testing.B) {
	lines := make([]string, 1000)
	for i := 0; i < len(lines); i++ {
		lines[i] = fmt.Sprintf("Hello... World... #%d !!!", i)
	}

	concat := func(s []string) string {
		var temp string
		for i := 0; i < len(s); i++ {
			temp += s[i]
		}
		return temp
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = concat(lines)
	}
}

func BenchmarkStringConcat_BytesBuffer(b *testing.B) {
	lines := make([]string, 1000)
	for i := 0; i < len(lines); i++ {
		lines[i] = fmt.Sprintf("Hello... World... #%d !!!", i)
	}

	concat := func(s []string) string {
		var bb bytes.Buffer
		for i := 0; i < len(s); i++ {
			if _, err := bb.WriteString(s[i]); err != nil {
				panic(err)
			}
		}
		return bb.String()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = concat(lines)
	}
}

func BenchmarkStringConcat_StringBuilder(b *testing.B) {
	lines := make([]string, 1000)
	for i := 0; i < len(lines); i++ {
		lines[i] = fmt.Sprintf("Hello... World... #%d !!!", i)
	}

	concat := func(s []string) string {
		var sb strings.Builder
		for i := 0; i < len(s); i++ {
			if _, err := sb.WriteString(s[i]); err != nil {
				panic(err)
			}
		}
		return sb.String()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = concat(lines)
	}
}
