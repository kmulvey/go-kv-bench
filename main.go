package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	/*
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

		result = testing.Benchmark(benchmarks.BenchmarkGoMapPutValue512B)
		br = goBenchmarkResultToBenchmarkResult("BenchmarkGoMapPutValue512B", result)
		insertBenchmarkResult(br)
		fmt.Println(result)
	*/
	var m = getAllBenchmarkResults()
	var c = marshalChartjs(m)
	js, _ := json.Marshal(c)
	fmt.Println(string(js))

	//drawChart("all-ns.jpeg", getAllBenchmarkResults(), "ns")
	//drawChart("all-alloc-op.jpeg", getAllBenchmarkResults(), "alloc-op")
	//drawChart("all-alloc-bytes.jpeg", getAllBenchmarkResults(), "alloc-bytes")
	os.RemoveAll("benchdata")
}
