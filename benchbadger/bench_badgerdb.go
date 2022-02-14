package benchbadger

import (
	"fmt"
	"os"
	"testing"

	"github.com/dgraph-io/badger/v3"
	"github.com/kmulvey/go-kv-bench/data"
)

const dir = "testdata/badgerDB"

func init() {
	os.RemoveAll(dir)
	os.Mkdir(dir, 0755)
}

func BenchmarkBadgerDBPutValue64B(b *testing.B) {
	b.ReportAllocs()
	badgerDB, err := badger.Open(badger.DefaultOptions(dir).WithLoggingLevel(badger.ERROR))
	if err != nil {
		panic(err)
	}
	defer func() {
		var err = badgerDB.Close()
		if err != nil {
			panic(err)
		}
	}()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		key := data.GetKey(n)
		val := data.GetValue64B()

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
	badgerDB, err := badger.Open(badger.DefaultOptions(dir))
	if err != nil {
		panic(err)
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		key := data.GetKey(n)
		val := data.GetValue128B()

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
	badgerDB, err := badger.Open(badger.DefaultOptions(dir))
	if err != nil {
		panic(err)
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		key := data.GetKey(n)
		val := data.GetValue256B()

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
	badgerDB, err := badger.Open(badger.DefaultOptions(dir))
	if err != nil {
		panic(err)
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		key := data.GetKey(n)
		val := data.GetValue512B()

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
	badgerDB, err := badger.Open(badger.DefaultOptions(dir))
	if err != nil {
		panic(err)
	}
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
