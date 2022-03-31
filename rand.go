package main

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"math"
	"math/rand"
)

func init() {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

func randFloat64() (r float64, err error) {
	r = rand.Float64()
	return
}
