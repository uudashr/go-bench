package bench_test

import (
	"bytes"
	_ "embed"
	"io"
	"os"
	"strings"
	"testing"
)

func BenchmarkReadReader_readAllJSONString(b *testing.B) {
	for b.Loop() {
		b.StopTimer()
		f, err := os.Open("testdata/json/error.json")
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()

		out, err := io.ReadAll(f)
		if err != nil {
			b.Fatal(err)
		}

		s := string(out)
		_ = s

		b.StopTimer()
		f.Close()
		b.StartTimer()
	}
}

func BenchmarkReadReader_stringsBuilderJSON(b *testing.B) {
	for b.Loop() {
		b.StopTimer()
		f, err := os.Open("testdata/json/error.json")
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()

		var sb strings.Builder
		_, err = io.Copy(&sb, f)
		if err != nil {
			b.Fatal(err)
		}

		s := sb.String()
		_ = s

		b.StopTimer()
		f.Close()
		b.StartTimer()
	}
}

func BenchmarkReadReader_bytesBufferJSONString(b *testing.B) {
	for b.Loop() {
		b.StopTimer()
		f, err := os.Open("testdata/json/error.json")
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()

		var buf bytes.Buffer
		_, err = buf.ReadFrom(f)
		if err != nil {
			b.Fatal(err)
		}

		s := buf.String()
		_ = s

		b.StopTimer()
		f.Close()
		b.StartTimer()
	}
}
