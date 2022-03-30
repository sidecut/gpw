package main

import (
	"crypto/rand"
	"fmt"
	"testing"
)

func TestRandFloat64(t *testing.T) {
	result1, err := randFloat64()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("result1 = %v", result1)

	result2, err := randFloat64()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("result2 = %v", result2)

	if result1 == result2 {
		t.Fatalf("%v should not be equal to %v", result1, result2)
	}
}

func TestRand(t *testing.T) {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	CheckError(err)

	fmt.Println(b) // [193 174 178 8 109 249 203 255 176 153]
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
