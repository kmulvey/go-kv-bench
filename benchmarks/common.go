package benchmarks

import (
	"os"
)

const baseDir = "benchdata"

func init() {
	// we bulldoze right past these errors
	os.RemoveAll(baseDir)

	// start over from scratch
	if err := os.Mkdir(baseDir, 0755); err != nil {
		panic(err)
	}
	if err := os.Mkdir(badgerDir, 0755); err != nil {
		panic(err)
	}
	if err := os.Mkdir(boltDBDir, 0755); err != nil {
		panic(err)
	}
	if err := os.Mkdir(nutsDir, 0755); err != nil {
		panic(err)
	}
}
