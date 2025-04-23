package bench_test

import (
	"encoding/json"
	"testing"

	"github.com/go-viper/mapstructure/v2"
)

func BenchmarkMapToStruct_mapstructure(b *testing.B) {
	m := map[string]interface{}{
		"name":         "John Doe",
		"age_in_years": float64(30),
	}

	b.ResetTimer()
	for b.Loop() {
		var p Person

		config := mapstructure.DecoderConfig{
			TagName: "json",
			Result:  &p,
		}

		decoder, err := mapstructure.NewDecoder(&config)
		if err != nil {
			b.Fatalf("failed to create decoder: %v", err)
		}

		err = decoder.Decode(m)
		if err != nil {
			b.Fatalf("failed to decode map: %v", err)
		}

		_ = p
	}

}

func BenchmarkMapToStruct_json(b *testing.B) {
	m := map[string]interface{}{
		"name":         "John Doe",
		"age_in_years": float64(30),
	}

	b.ResetTimer()
	for b.Loop() {
		var p Person

		data, err := json.Marshal(m)
		if err != nil {
			b.Fatalf("failed to marshal map: %v", err)
		}

		err = json.Unmarshal(data, &p)
		if err != nil {
			b.Fatalf("failed to unmarshal map: %v", err)
		}

		_ = p
	}
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age_in_years"`
}
