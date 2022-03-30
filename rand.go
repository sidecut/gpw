package main

import (
	"crypto/rand"
	"encoding/binary"
	"math"
)

const maxIntFloat = float64(math.MaxInt64)

func randFloat64() (r float64, err error) {
	var b [8]byte
	_, err = rand.Read(b[:])
	if err != nil {
		return
	}
	r = Float64frombytes(b[:])
	return

}

func Float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := float64(bits) / maxIntFloat
	return float
}
