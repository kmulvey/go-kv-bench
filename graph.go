package main

import (
	"time"
)

// 1. line graqh comparing all over time
// 2. line graph for single over time
// 3. bar graph for most recent version of each

type yValueGetter func(benchmarkResult) float64

type graph struct {
	Title      string
	Benchmarks []benchmarkHistory
}

type benchmarkHistory []benchmarkResult

func drawChart(g graph, yGetter yValueGetter) {
	/*
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
					ts.YValues[j] = yGetter(b)
				}
				graph.Series[i] = ts
			}

			fmt.Printf("%+v\n", graph.Series)

			// render the above into a PNG
			var err error
			var collector = &chart.ImageWriter{}
			graph.Render(chart.PNG, collector)

		//////////////////////////////////////////////////
		p := plot.New()

		p.Title.Text = "Plotutil example"
		p.X.Label.Text = "X"
		p.Y.Label.Text = "Y"

		err := plotutil.AddLinePoints(p,
			"First", randomPoints(15),
			"Second", randomPoints(15),
			"Third", randomPoints(15))
		if err != nil {
			panic(err)
		}

		// Save the plot to a PNG file.
		if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
			panic(err)
		}
		// get the Go image.Image out
		pngData, err := collector.Image()
		handleErr("get image.Image", err)

		// save it as a jpg
		out, err := os.Create("graph.jpeg")
		handleErr("jpg create", err)
		handleErr("jpg encode", jpeg.Encode(out, pngData, &jpeg.Options{Quality: 85}))
		handleErr("jpg close", out.Close())

	*/
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

func getLongestTimeseries(g graph) []time.Time {
	var longest []benchmarkResult
	for _, benchHist := range g.Benchmarks {
		if len(benchHist) > len(longest) {
			longest = benchHist
		}
	}

	var timeseries = make([]time.Time, len(longest))
	for i, b := range longest {
		timeseries[i] = b.Benchtime
	}
	return timeseries
}
