package main

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type benchmarkResult struct {
	Benchname   string
	Benchtime   time.Time
	Alloc_bytes int64
	Alloc_op    int64
	Ns_op       int64
}

func goBenchmarkResultToBenchmarkResult(benchName string, tbr testing.BenchmarkResult) benchmarkResult {
	return benchmarkResult{
		Benchname:   benchName,
		Benchtime:   time.Now(),
		Alloc_bytes: tbr.AllocedBytesPerOp(),
		Alloc_op:    tbr.AllocsPerOp(),
		Ns_op:       tbr.NsPerOp(),
	}
}

func closeDB(db *sql.DB) {
	var err = db.Close()
	handleErr("close sqlite db", err)
}

func insertBenchmarkResult(br benchmarkResult) {
	// open
	var benchResultsDB, err = sql.Open("sqlite3", "./benchresults/bench_results.db")
	handleErr("open sqlite db", err)
	defer closeDB(benchResultsDB)

	// insert
	insertStudentSQL := `INSERT INTO bench_results(BENCHNAME, BENCHTIME, ALLOC_BYTES, ALLOC_OP, NS_OP) VALUES (?, ?, ?, ?, ?)`
	statement, err := benchResultsDB.Prepare(insertStudentSQL)
	handleErr("prepare sql", err)
	_, err = statement.Exec(br.Benchname, br.Benchtime, br.Alloc_bytes, br.Alloc_op, br.Ns_op)
	handleErr("insert bench result", err)
}

func getAllBenchmarkResults() map[string][]benchmarkResult {
	// open
	var benchResultsDB, err = sql.Open("sqlite3", "./benchresults/bench_results.db")
	handleErr("open sqlite db", err)
	defer closeDB(benchResultsDB)

	rows, err := benchResultsDB.Query("select BENCHNAME, BENCHTIME, ALLOC_BYTES, ALLOC_OP, NS_OP from bench_results order by BENCHNAME")
	handleErr("select bench result", err)

	// marshal data
	var allBenchmarkResults = make(map[string][]benchmarkResult) // map of benchmark name => slice of individual benchmark results
	var lastBenchName string                                     // used to seperate benchmarks by name
	var oneBenchmarkHistory []benchmarkResult                    // the history of a single benchmark
	for rows.Next() {
		var br = new(benchmarkResult)
		var err = rows.Scan(&br.Benchname, &br.Benchtime, &br.Alloc_bytes, &br.Alloc_op, &br.Ns_op)
		handleErr("scan bench result", err)

		if br.Benchname != lastBenchName {
			allBenchmarkResults[br.Benchname] = oneBenchmarkHistory
			oneBenchmarkHistory = oneBenchmarkHistory[:0] // reset this for nect benchmark
			lastBenchName = br.Benchname
		}
		oneBenchmarkHistory = append(oneBenchmarkHistory, *br)
	}
	return allBenchmarkResults
}
