package main

import (
	"fmt"
	"testing"

	"github.com/kmulvey/go-kv-bench/benchbadger"
)

/*
func init() {
	var err = os.Mkdir("testdata", 0755)
	fmt.Println("err:", err)
	if err != nil {
		panic(err)
	}
}
*/

func main() {
	res := testing.Benchmark(benchbadger.BenchmarkBadgerDBPutValue64B)
	fmt.Printf("%s\n%#[1]v\n", res)
}
