package bench_test

import (
	"reflect"
	"testing"
)

type UserProfile struct {
	FirstName string
	LastName  string
	Sex       string
	Address   string
}

func BenchmarkZeroStructCheck_emptyStructVar(b *testing.B) {
	var userProfile UserProfile
	empty := UserProfile{}
	for b.Loop() {
		var zero bool
		if userProfile == empty {
			zero = true
		}
		_ = zero
	}
}

func BenchmarkZeroStructCheck_directEmptyStruct(b *testing.B) {
	var userProfile UserProfile
	for b.Loop() {
		var zero bool
		if userProfile == (UserProfile{}) {
			zero = true
		}
		_ = zero
	}
}

func BenchmarkZeroStructCheck_reflection(b *testing.B) {
	var userProfile UserProfile
	for b.Loop() {
		var zero bool
		if reflect.ValueOf(userProfile).IsZero() {
			zero = true
		}
		_ = zero
	}
}

func BenchmarkZeroStructCheck_fieldCheck(b *testing.B) {
	var userProfile UserProfile
	for b.Loop() {
		var zero bool
		if userProfile.FirstName == "" { // use FirstName as zero check
			zero = true
		}
		_ = zero
	}
}
