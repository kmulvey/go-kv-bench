package benchmarks

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xujiajun/nutsdb"
)

var nutsDir = filepath.Join(baseDir, "nutsDB")

func InitNutsDBByDefaultOpt(b *testing.B) *nutsdb.DB {
	opt := nutsdb.DefaultOptions
	opt.Dir = nutsDir
	opt.SegmentSize = 64 * 1024 * 1024
	var nutsDB, err = nutsdb.Open(opt)
	assert.NoError(b, err)
	return nutsDB
}

func InitNutsDBByByHintKeyOpt(b *testing.B) *nutsdb.DB {
	opt := nutsdb.DefaultOptions
	opt.Dir = "testdata/nutsDB"
	opt.SegmentSize = 64 * 1024 * 1024
	opt.EntryIdxMode = nutsdb.HintKeyAndRAMIdxMode
	var nutsDB, err = nutsdb.Open(opt)
	assert.NoError(b, err)
	return nutsDB
}

func deferNutsClose(b *testing.B, nutsDB *nutsdb.DB) {
	var err = nutsDB.Close()
	assert.NoError(b, err)
}

func InitNutsDBData(b *testing.B, nutsDB *nutsdb.DB) {
	for n := 0; n < 10000; n++ {
		key := GetKey(n)
		val := GetValue64B()

		var err = nutsDB.Update(
			func(tx *nutsdb.Tx) error {
				return tx.Put("bucket1", key, val, 0)
			})
		assert.NoError(b, err)
	}
}

func BenchmarkNutsDBPutValue64B(b *testing.B) {
	var db = InitNutsDBByDefaultOpt(b)
	defer deferNutsClose(b, db)

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		key := GetKey(n)
		val := GetValue64B()
		var err = db.Update(
			func(tx *nutsdb.Tx) error {
				return tx.Put("bucket1", key, val, 0)
			})
		assert.NoError(b, err)
	}
}

func BenchmarkNutsDBPutValue128B(b *testing.B) {
	var db = InitNutsDBByDefaultOpt(b)
	defer deferNutsClose(b, db)

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		key := GetKey(n)
		val := GetValue128B()
		var err = db.Update(
			func(tx *nutsdb.Tx) error {
				return tx.Put("bucket1", key, val, 0)
			})
		assert.NoError(b, err)
	}
}

func BenchmarkNutsDBPutValue256B(b *testing.B) {
	var db = InitNutsDBByDefaultOpt(b)
	defer deferNutsClose(b, db)

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		key := GetKey(n)
		val := GetValue256B()
		var err = db.Update(
			func(tx *nutsdb.Tx) error {
				return tx.Put("bucket1", key, val, 0)
			})
		assert.NoError(b, err)
	}
}

func BenchmarkNutsDBPutValue512B(b *testing.B) {
	var db = InitNutsDBByDefaultOpt(b)
	defer deferNutsClose(b, db)

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		key := GetKey(n)
		val := GetValue512B()
		var err = db.Update(
			func(tx *nutsdb.Tx) error {
				return tx.Put("bucket1", key, val, 0)
			})
		assert.NoError(b, err)
	}
}

func BenchmarkNutsDBGet(b *testing.B) {
	var db = InitNutsDBByDefaultOpt(b)
	defer deferNutsClose(b, db)
	InitNutsDBData(b, db)

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		var err = db.View(
			func(tx *nutsdb.Tx) error {
				key := []byte("key_" + fmt.Sprintf("%07d", 99))
				if _, err := tx.Get("bucket1", key); err != nil {
					return err
				}
				return nil
			})
		assert.NoError(b, err)
	}
}

func BenchmarkNutsDBGetByHintKey(b *testing.B) {
	var db = InitNutsDBByDefaultOpt(b)
	defer deferNutsClose(b, db)
	InitNutsDBData(b, db)

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		var err = db.View(
			func(tx *nutsdb.Tx) error {
				key := []byte("key_" + fmt.Sprintf("%07d", 99))
				if _, err := tx.Get("bucket1", key); err != nil {
					return err
				}
				return nil
			})
		assert.NoError(b, err)
	}
}
