package randomUtils

import (
	"math/rand"
	"unsafe"
)

var alphabetm = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// var alphabet = []byte("45678930abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ45789")

func Generate(size int) string {
	b := make([]byte, size)
	rand.Read(b)
	for i := 0; i < size; i++ {
		b[i] = alphabetm[b[i]/5]
	}
	return *(*string)(unsafe.Pointer(&b))
}
