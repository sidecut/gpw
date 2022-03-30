package main

import (
	crypto_rand "crypto/rand"
	"math"
	"math/big"
	math_rand "math/rand"
)

const maxIntFloat = float64(math.MaxInt64)

func init() {
	seed, err := crypto_rand.Int(crypto_rand.Reader, big.NewInt(1048576))
	if err != nil {
		panic(err)
	}
	math_rand.Seed(seed.Int64())
}

func randFloat64() (r float64, err error) {
	r = math_rand.Float64()
	return
}
