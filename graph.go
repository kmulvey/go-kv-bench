package main

import (
	"image/jpeg"
	"os"
	"time"

	"github.com/wcharczuk/go-chart/v2"
)

// 1. line graqh comparing all over time
// 2. line graph for single over time
// 3. bar graph for most recent version of each

type yValueGetter func(benchmarkResult) float64

type benchSeries struct {
	Name    string
	Results []benchmarkResult
}

func drawChart(data []benchSeries, getter yValueGetter) {
	var graph = new(chart.Chart)
	graph.Series = make([]chart.Series, len(data))

	for i, s := range data {
		var ts = chart.TimeSeries{
			Name:    s.Name,
			XValues: make([]time.Time, len(s.Results)),
			YValues: make([]float64, len(s.Results)),
		}
		for j, b := range s.Results {
			ts.XValues[j] = b.Benchtime
			ts.YValues[j] = getter(b)
		}
		graph.Series[i] = ts
	}

	// render the above into a PNG
	var err error
	var collector = &chart.ImageWriter{}
	graph.Render(chart.PNG, collector)

	// get the Go image.Image out
	pngData, err := collector.Image()
	handleErr("get image.Image", err)

	// save it as a jpg
	out, err := os.Create("graph.jpeg")
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
