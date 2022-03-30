package main

import "testing"

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
