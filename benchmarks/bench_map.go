package benchmarks

import (
	"strconv"
	"testing"

	"github.com/kmulvey/go-kv-bench/testdata"
	"github.com/stretchr/testify/assert"
)

func BenchmarkGoMapPutValue64B(b *testing.B) {
	b.ReportAllocs()
	var m = make(map[string]string)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		m[string(testdata.GetKey(n))] = string(testdata.GetValue64B())
	}
}

func BenchmarkGoMapPutValue128B(b *testing.B) {
	b.ReportAllocs()
	var m = make(map[string]string)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		m[string(testdata.GetKey(n))] = string(testdata.GetValue128B())
	}
}

func BenchmarkGoMapPutValue256B(b *testing.B) {
	b.ReportAllocs()
	var m = make(map[string]string)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		m[string(testdata.GetKey(n))] = string(testdata.GetValue256B())
	}
}

func BenchmarkGoMapPutValue512B(b *testing.B) {
	b.ReportAllocs()
	var m = make(map[string]string)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		m[string(testdata.GetKey(n))] = string(testdata.GetValue512B())
	}
}

func BenchmarkGoMapGet(b *testing.B) {
	b.ReportAllocs()
	var m = make(map[string]string)
	for n := 0; n < b.N; n++ {
		m[string(testdata.GetKey(n))] = string(testdata.GetValue512B())
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		var val, ok = m[strconv.Itoa(n)]
		assert.True(b, ok)
		assert.NotEmpty(b, val)
	}
}
