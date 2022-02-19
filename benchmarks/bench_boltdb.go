package benchmarks

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/boltdb/bolt"
	"github.com/kmulvey/go-kv-bench/testdata"
	"github.com/stretchr/testify/assert"
)

var boltDBDir = filepath.Join(baseDir, "boltDB")
var boltDBFile = filepath.Join(boltDBDir, "bolt.db")

func deferBoltClose(b *testing.B, boltDB *bolt.DB) {
	var err = boltDB.Close()
	assert.NoError(b, err)
}

func BenchmarkBoltDBPutValue64B(b *testing.B) {
	b.ReportAllocs()
	var boltDB, err = bolt.Open(boltDBFile, 0666, nil)
	assert.NoError(b, err)
	defer deferBoltClose(b, boltDB)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		key := testdata.GetKey(n)
		val := testdata.GetValue64B()
		if err = boltDB.Update(func(tx *bolt.Tx) error {
			bucket, err := tx.CreateBucketIfNotExists([]byte("bucket1"))
			if err != nil {
				return err
			}

			return bucket.Put([]byte(key), val)
		}); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBoltDBPutValue128B(b *testing.B) {
	b.ReportAllocs()
	var boltDB, err = bolt.Open(boltDBFile, 0666, nil)
	assert.NoError(b, err)
	defer deferBoltClose(b, boltDB)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		key := testdata.GetKey(n)
		val := testdata.GetValue128B()
		var err = boltDB.Update(func(tx *bolt.Tx) error {
			bucket, err := tx.CreateBucketIfNotExists([]byte("bucket1"))
			if err != nil {
				return err
			}

			return bucket.Put([]byte(key), val)
		})
		assert.NoError(b, err)
	}
}

func BenchmarkBoltDBPutValue256B(b *testing.B) {
	b.ReportAllocs()
	var boltDB, err = bolt.Open(boltDBFile, 0666, nil)
	assert.NoError(b, err)
	defer deferBoltClose(b, boltDB)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		key := testdata.GetKey(n)
		val := testdata.GetValue256B()
		var err = boltDB.Update(func(tx *bolt.Tx) error {
			bucket, err := tx.CreateBucketIfNotExists([]byte("bucket1"))
			if err != nil {
				return err
			}

			return bucket.Put([]byte(key), val)
		})
		assert.NoError(b, err)
	}
}

func BenchmarkBoltDBPutValue512B(b *testing.B) {
	b.ReportAllocs()
	var boltDB, err = bolt.Open(boltDBFile, 0666, nil)
	assert.NoError(b, err)
	defer deferBoltClose(b, boltDB)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		key := testdata.GetKey(n)
		val := testdata.GetValue512B()
		var err = boltDB.Update(func(tx *bolt.Tx) error {
			bucket, err := tx.CreateBucketIfNotExists([]byte("bucket1"))
			if err != nil {
				return err
			}

			return bucket.Put([]byte(key), val)
		})
		assert.NoError(b, err)
	}
}

func BenchmarkBoltDBGet(b *testing.B) {
	b.ReportAllocs()
	var boltDB, err = bolt.Open(boltDBFile, 0666, nil)
	assert.NoError(b, err)
	defer deferBoltClose(b, boltDB)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		var err = boltDB.View(func(tx *bolt.Tx) error {
			key := []byte("key_" + fmt.Sprintf("%07d", 99))
			tx.Bucket([]byte("bucket1")).Get([]byte(key))
			return nil
		})
		assert.NoError(b, err)
	}
}
