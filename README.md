# Go-Bench

Collection of go benchmarks.

## Usage
Run all benchmark
```
go test -benchmem -run="^$" -bench=.
```

Run specific benchmark
```
go test -benchmem -run="^$" -bench="^BenchmarkStringConvert_"
```