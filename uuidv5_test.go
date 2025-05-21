package bench_test

import (
	"testing"

	"github.com/google/uuid"
)

func BenchmarkUUIDv5_noPredefinedNS(b *testing.B) {
	for b.Loop() {
		ns := uuid.NewSHA1(uuid.NameSpaceDNS, []byte("example.com"))
		res := uuid.NewSHA1(ns, []byte("test"))
		_ = res.String()
	}
}

func BenchmarkUUIDv5_predefinedNS(b *testing.B) {
	ns := uuid.NewSHA1(uuid.NameSpaceDNS, []byte("example.com"))
	for b.Loop() {
		res := uuid.NewSHA1(ns, []byte("test"))
		_ = res.String()
	}
}
