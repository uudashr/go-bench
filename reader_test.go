package bench_test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func BenchmarkReadReader_readAllString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		f, err := os.Open("testdata/loremipsum/lorem_5p.txt")
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

func BenchmarkReadReader_stringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, err := os.Open("testdata/loremipsum/lorem_5p.txt")
		if err != nil {
			b.Fatal(err)
		}

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

func BenchmarkReadReader_bytesBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, err := os.Open("testdata/loremipsum/lorem_5p.txt")
		if err != nil {
			b.Fatal(err)
		}

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
