package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/kmulvey/go-kv-bench/benchmarks"
)

func main() {
	var result = testing.Benchmark(benchmarks.BenchmarkBadgerDBPutValue512B)
	var br = goBenchmarkResultToBenchmarkResult("BenchmarkBadgerDBPutValue512B", result)
	insertBenchmarkResult(br)
	fmt.Println(result)

	result = testing.Benchmark(benchmarks.BenchmarkBoltDBPutValue512B)
	br = goBenchmarkResultToBenchmarkResult("BenchmarkBoltDBPutValue512B", result)
	insertBenchmarkResult(br)
	fmt.Println(result)

	result = testing.Benchmark(benchmarks.BenchmarkNutsDBPutValue512B)
	br = goBenchmarkResultToBenchmarkResult("BenchmarkNutsDBPutValue512B", result)
	insertBenchmarkResult(br)
	fmt.Println(result)

	//var bs = benchSeries{
	//	Name:    "BenchmarkBadgerDBPutValue512B",
	//	Results: []benchmarkResult{br},
	//}
	//drawChart([]benchSeries{bs}, getNsOp)
	os.RemoveAll("benchdata")
}
