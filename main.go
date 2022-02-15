package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/kmulvey/go-kv-bench/benchbadger"
)

func main() {
	res := testing.Benchmark(benchbadger.BenchmarkBadgerDBPutValue512B)
	fmt.Printf("%s\n%#[1]v\n", res)
	os.RemoveAll("benchdata")
}
