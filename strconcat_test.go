package bench_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func BenchmarkStringConcat_Literal(b *testing.B) {
	lines := make([]string, 1000)
	for i := range lines {
		lines[i] = fmt.Sprintf("Hello... World... #%d !!!", i)
	}

	concat := func(s []string) string {
		var temp string
		for i := range s {
			temp += s[i]
		}
		return temp
	}

	for b.Loop() {
		_ = concat(lines)
	}
}

func BenchmarkStringConcat_BytesBuffer(b *testing.B) {
	lines := make([]string, 1000)
	for i := range lines {
		lines[i] = fmt.Sprintf("Hello... World... #%d !!!", i)
	}

	concat := func(s []string) string {
		var bb bytes.Buffer
		for i := range s {
			if _, err := bb.WriteString(s[i]); err != nil {
				panic(err)
			}
		}
		return bb.String()
	}

	for b.Loop() {
		_ = concat(lines)
	}
}

func BenchmarkStringConcat_StringBuilder(b *testing.B) {
	lines := make([]string, 1000)
	for i := range lines {
		lines[i] = fmt.Sprintf("Hello... World... #%d !!!", i)
	}

	concat := func(s []string) string {
		var sb strings.Builder
		for i := range s {
			if _, err := sb.WriteString(s[i]); err != nil {
				panic(err)
			}
		}
		return sb.String()
	}

	for b.Loop() {
		_ = concat(lines)
	}
}
