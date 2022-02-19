package benchmarks

import (
	"fmt"
	"os"
	"testing"

	"github.com/dgraph-io/badger/v3"
	"github.com/kmulvey/go-kv-bench/testdata"
)

const badgerDir = baseDir + "/badgerDB"

var badgerOpts = badger.DefaultOptions(badgerDir).WithLoggingLevel(badger.ERROR)

func init() {
	// we bulldoze right past these errors
	os.RemoveAll(badgerDir)
	os.Mkdir(badgerDir, 0755)
}

func deferClose(badgerDB *badger.DB) {
	var err = badgerDB.Close()
	if err != nil {
		panic(err)
	}
}

func BenchmarkBadgerDBPutValue64B(b *testing.B) {
	b.ReportAllocs()
	badgerDB, err := badger.Open(badgerOpts)
	if err != nil {
		panic(err)
	}
	defer deferClose(badgerDB)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		key := testdata.GetKey(n)
		val := testdata.GetValue64B()

		if err = badgerDB.Update(
			func(txn *badger.Txn) error {
				return txn.Set(key, val)
			}); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBadgerDBPutValue128B(b *testing.B) {
	b.ReportAllocs()
	badgerDB, err := badger.Open(badgerOpts)
	if err != nil {
		panic(err)
	}
	defer deferClose(badgerDB)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		key := testdata.GetKey(n)
		val := testdata.GetValue128B()

		if err = badgerDB.Update(
			func(txn *badger.Txn) error {
				return txn.Set(key, val)
			}); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBadgerDBPutValue256B(b *testing.B) {
	b.ReportAllocs()
	badgerDB, err := badger.Open(badgerOpts)
	if err != nil {
		panic(err)
	}
	defer deferClose(badgerDB)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		key := testdata.GetKey(n)
		val := testdata.GetValue256B()

		if err = badgerDB.Update(
			func(txn *badger.Txn) error {
				return txn.Set(key, val)
			}); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBadgerDBPutValue512B(b *testing.B) {
	b.ReportAllocs()
	badgerDB, err := badger.Open(badgerOpts)
	if err != nil {
		panic(err)
	}
	defer deferClose(badgerDB)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		key := testdata.GetKey(n)
		val := testdata.GetValue512B()

		if err = badgerDB.Update(
			func(txn *badger.Txn) error {
				return txn.Set(key, val)
			}); err != nil {
			b.Fatal(err)
		}
	}
}
func BenchmarkBadgerDBGet(b *testing.B) {
	b.ReportAllocs()
	badgerDB, err := badger.Open(badgerOpts)
	if err != nil {
		panic(err)
	}
	defer deferClose(badgerDB)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		err = badgerDB.View(func(txn *badger.Txn) error {
			key := []byte("key_" + fmt.Sprintf("%07d", 99))
			_, err := txn.Get(key)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			b.Fatal(err)
		}
	}
}
