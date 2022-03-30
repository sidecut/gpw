package main

import (
	"crypto/rand"
	"encoding/binary"
	"math"
)

func randFloat64() (float64, error) {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return 0, err
	}
	return Float64frombytes(b), nil
}

func Float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}
