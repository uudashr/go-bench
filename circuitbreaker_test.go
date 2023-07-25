package bench_test

import (
	"context"
	"testing"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	cep21circuit "github.com/cep21/circuit/v3"
	eapachebreaker "github.com/eapache/go-resiliency/breaker"
	rubyistbreaker "github.com/rubyist/circuitbreaker"
	resiliencebreaker "github.com/slok/goresilience/circuitbreaker"
	"github.com/sony/gobreaker"
	"github.com/streadway/handy/breaker"
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
		hystrix.Do("simpleCall", func() error {
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
	cb := rubyistbreaker.NewBreaker()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cb.Call(func() error {
			_, err := simpleCall()
			return err
		}, 1*time.Second)
	}
}

func BenchmarkCircuitBreaker_Cep21Circuit(b *testing.B) {
	cb := cep21circuit.NewCircuitFromConfig("simpleCall", cep21circuit.Config{})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cb.Run(context.Background(), func(ctx context.Context) error {
			_, err := simpleCall()
			return err
		})
	}
}

func BenchmarkCircuitBreaker_GoResiliency(b *testing.B) {
	cb := eapachebreaker.New(5, 10, 5*time.Second)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cb.Run(func() error {
			_, err := simpleCall()
			return err
		})
	}
}

func BenchmarkCircuitBreaker_GoResilience(b *testing.B) {
	runner := resiliencebreaker.New(resiliencebreaker.Config{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		runner.Run(context.Background(), func(context.Context) error {
			_, err := simpleCall()
			return err
		})
	}
}

func simpleCall() (string, error) {
	// do nothing
	return "Hello, World!", nil
}
