package benchmarks

import (
	"fmt"
	"testing"

	"github.com/dgraph-io/badger/v3"
)

const badgerDir = baseDir + "/badgerDB"

var badgerOpts = badger.DefaultOptions(badgerDir).WithLoggingLevel(badger.ERROR)

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
		key := GetKey(n)
		val := GetValue64B()

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
		key := GetKey(n)
		val := GetValue128B()

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
		key := GetKey(n)
		val := GetValue256B()

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
		key := GetKey(n)
		val := GetValue512B()

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
