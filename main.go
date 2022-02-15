package main

import (
	"os"
	"testing"

	"github.com/kmulvey/go-kv-bench/benchbadger"
)

func main() {
	var result = testing.Benchmark(benchbadger.BenchmarkBadgerDBPutValue512B)
	insertBenchmarkResult(goBenchmarkResultToSQLBenchmarkResult("BenchmarkBadgerDBPutValue512B", result))
	os.RemoveAll("benchdata")
}
