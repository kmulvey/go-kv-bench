package benchmarks

import "fmt"

func GetKey(n int) []byte {
	return []byte("key_" + fmt.Sprintf("%07d", n))
}

func GetValue64B() []byte {
	return make([]byte, 64)
}

func GetValue128B() []byte {
	return make([]byte, 128)
}

func GetValue256B() []byte {
	return make([]byte, 256)
}

func GetValue512B() []byte {
	return make([]byte, 512)
}
