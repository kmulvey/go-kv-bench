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

func goBenchmarkResultToSQLBenchmarkResult(benchName string, tbr testing.BenchmarkResult) benchmarkResult {
	return benchmarkResult{
		Benchname:   benchName,
		Benchtime:   time.Now(),
		Alloc_bytes: tbr.AllocedBytesPerOp(),
		Alloc_op:    tbr.AllocsPerOp(),
		Ns_op:       tbr.NsPerOp(),
	}
}

func insertBenchmarkResult(br benchmarkResult) {
	// open
	var benchResultsDB, err = sql.Open("sqlite3", "./benchresults/bench_results.db")
	handleErr("open sqlite db", err)
	defer func() {
		var err = benchResultsDB.Close()
		handleErr("close sqlite db", err)
	}()

	// insert
	insertStudentSQL := `INSERT INTO bench_results(BENCHNAME, BENCHTIME, ALLOC_BYTES, ALLOC_OP, NS_OP) VALUES (?, ?, ?, ?, ?)`
	statement, err := benchResultsDB.Prepare(insertStudentSQL)
	handleErr("prepare sql", err)
	_, err = statement.Exec(br.Benchname, br.Benchtime, br.Alloc_bytes, br.Alloc_op, br.Ns_op)
	handleErr("insert bench result", err)
}
