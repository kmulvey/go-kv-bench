package main

import (
	"os"

	log "github.com/sirupsen/logrus"
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
	var err = writeChartDataFile(m, "benchresults/benchmark_data.js")
	if err != nil {
		log.Error(err)
	}

	os.RemoveAll("benchdata")
}
