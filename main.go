package main

import (
	"os"
	"testing"

	"github.com/kmulvey/go-kv-bench/benchbadger"
)

func main() {
	var result = testing.Benchmark(benchbadger.BenchmarkBadgerDBPutValue512B)
	var br = goBenchmarkResultToBenchmarkResult("BenchmarkBadgerDBPutValue512B", result)
	insertBenchmarkResult(br)

	var bs = benchSeries{
		Name:    "BenchmarkBadgerDBPutValue512B",
		Results: []benchmarkResult{br},
	}
	drawChart([]benchSeries{bs}, getNsOp)
	os.RemoveAll("benchdata")
}
