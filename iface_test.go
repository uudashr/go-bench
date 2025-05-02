package bench_test

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkFuncCall(b *testing.B) {
	for b.Loop() {
		stat := Binary(200)
		_ = stringCall(stat)
	}
}

func BenchmarkFuncCallIface(b *testing.B) {
	for b.Loop() {
		b := Binary(200)
		_ = stringCallIface(b)
	}
}

type Binary uint64

func (b Binary) String() string {
	return strconv.FormatUint(uint64(b), 2)
}

func stringCall(b Binary) string {
	return b.String()
}

func stringCallIface(stat fmt.Stringer) string {
	return stat.String()
}
