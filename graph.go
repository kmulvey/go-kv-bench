package main

import (
	"image/jpeg"
	"os"
	"path/filepath"
	"time"

	"github.com/wcharczuk/go-chart/v2"
)

const outputDir = "./benchresults/graphs"

// 1. line graqh comparing all over time
// 2. line graph for single over time
// 3. bar graph for most recent version of each

type yValueGetter func(benchmarkResult) float64

func drawChart(title string, benchmarks map[string][]benchmarkResult, yGetter yValueGetter) {
	var graph = new(chart.Chart)
	graph.Series = make([]chart.Series, len(benchmarks))
	graph.Title = title

	var i int
	for benchmarkName, benchmarkHistory := range benchmarks {
		var ts = chart.TimeSeries{
			Name:    benchmarkName,
			XValues: make([]time.Time, len(benchmarkHistory)),
			YValues: make([]float64, len(benchmarkHistory)),
		}
		for j, b := range benchmarkHistory {
			ts.XValues[j] = b.Benchtime
			ts.YValues[j] = yGetter(b)
		}
		graph.Series[i] = ts
		i++
	}

	// render the above into a PNG
	var collector = &chart.ImageWriter{}
	handleErr("graph.Render", graph.Render(chart.PNG, collector))

	// get the Go image.Image out
	pngData, err := collector.Image()
	handleErr("get image.Image", err)

	// save it as a jpg
	out, err := os.Create(filepath.Join(outputDir, title+".jpeg"))
	handleErr("jpg create", err)
	handleErr("jpg encode", jpeg.Encode(out, pngData, &jpeg.Options{Quality: 85}))
	handleErr("jpg close", out.Close())
}

// getNsOp conforms to yValueGetter and returns the nanoseconds per op
func getNsOp(b benchmarkResult) float64 {
	return float64(b.Ns_op)
}

// getAllocBytes conforms to yValueGetter and returns the number of bytes allocated
func getAllocBytes(b benchmarkResult) float64 {
	return float64(b.Alloc_bytes)
}

// getAllocOp conforms to yValueGetter and returns number of allocs per op
func getAllocOp(b benchmarkResult) float64 {
	return float64(b.Alloc_op)
}
