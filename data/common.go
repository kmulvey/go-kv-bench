package data

import "fmt"

func GetKey(n int) []byte {
	return []byte("key_" + fmt.Sprintf("%07d", n))
}

func GetValue64B() []byte {
	return []byte("valvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalv")
}

func GetValue128B() []byte {
	return []byte("valvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvval" +
		"valvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalv")
}

func GetValue256B() []byte {
	return []byte("valvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvvalvalvalvalvalvalvalvalvalvalval" +
		"valvalvalvalvalvalvalvalvalvalvvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvvalvalvalval" +
		"valvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalv")
}

func GetValue512B() []byte {
	return []byte("valvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvvalvalvalvalvalvalvalvalvalvalval" +
		"valvalvalvalvalvalvalvalvalvalvvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvvalvalvalval" +
		"valvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalv" + "valvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvvalvalvalvalvalvalvalvalvalvalval" +
		"valvalvalvalvalvalvalvalvalvalvvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvvalvalvalval" +
		"valvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalvalv")
}
