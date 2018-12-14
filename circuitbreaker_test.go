package bench_test

import (
	"testing"
	"time"

	circuit "github.com/rubyist/circuitbreaker"

	"github.com/streadway/handy/breaker"

	"github.com/sony/gobreaker"

	"github.com/afex/hystrix-go/hystrix"
)

func BenchmarkCircuitBreaker_None(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simpleCall()
	}
}

func BenchmarkCircuitBreaker_Hystrix(b *testing.B) {
	hystrix.ConfigureCommand("simpleCall", hystrix.CommandConfig{})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hystrix.Go("simpleCall", func() error {
			_, err := simpleCall()
			return err
		}, nil)
	}
}

func BenchmarkCircuitBreaker_GoBreaker(b *testing.B) {
	cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name: "simpleCall",
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cb.Execute(func() (interface{}, error) {
			return simpleCall()
		})
	}
}

func BenchmarkCircuitBreaker_HandyBreaker(b *testing.B) {
	cb := breaker.NewBreaker(0.05)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !cb.Allow() {
			continue
		}

		begin := time.Now()
		_, err := simpleCall()
		if err != nil {
			cb.Failure(time.Since(begin))
			continue
		}

		cb.Success(time.Since(begin))
	}
}

func BenchmarkCircuitBreaker_RubyistBreaker(b *testing.B) {
	cb := circuit.NewBreaker()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cb.Call(func() error {
			_, err := simpleCall()
			return err
		}, 1*time.Second)
	}
}

func simpleCall() (string, error) {
	// do nothing
	return "Hello, World!", nil
}
