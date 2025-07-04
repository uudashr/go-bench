package bench_test

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func BenchmarkStringReader_alloc_bytesReader(b *testing.B) {
	for b.Loop() {
		_ = bytes.NewBufferString(`{"name":"John","age":30,"city":"New York"}`)
	}

}

func BenchmarkStringReader_alloc_stringReader(b *testing.B) {
	for b.Loop() {
		_ = strings.NewReader(`{"name":"John","age":30,"city":"New York"}`)
	}
}

func BenchmarkStringReader_readAll_bytesReader(b *testing.B) {
	for b.Loop() {
		buf := bytes.NewBufferString(`{"name":"John","age":30,"city":"New York"}`)

		_, err := io.ReadAll(buf)
		if err != nil {
			b.Fatal(err)
		}
	}

}

func BenchmarkStringReader_readAll_stringReader(b *testing.B) {
	for b.Loop() {
		r := strings.NewReader(`{"name":"John","age":30,"city":"New York"}`)

		_, err := io.ReadAll(r)
		if err != nil {
			b.Fatal(err)
		}
	}
}
