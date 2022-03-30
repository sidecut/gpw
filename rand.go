package main

import (
	"crypto/rand"
	"encoding/binary"
	"math"
	"math/big"
)

const maxIntFloat = float64(math.MaxInt64)

func randFloat64() (r float64, err error) {
	n, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		return
	}

	r = float64(n.Int64()) / maxIntFloat
	return
}

func Float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := float64(bits) / maxIntFloat / 2
	return float
}
