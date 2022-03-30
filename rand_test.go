package main

import (
	"fmt"
	"testing"
	"testing/quick"
)

func TestRandFloat64(t *testing.T) {
	quick.Check(func() bool {
		result1, err := randFloat64()
		if err != nil {
			t.Error(err)
			return false
		}
		t.Logf("result1 = %v", result1)

		result2, err := randFloat64()
		if err != nil {
			t.Error(err)
			return false
		}
		t.Logf("result2 = %v", result2)

		if result1 == result2 {
			t.Logf("%v should not be equal to %v", result1, result2)
			return false
		}
		return true
	}, nil)
}

func TestRandRange(t *testing.T) {
	var low, high float64
	low = 1
	high = 0

	quick.Check(func() bool {
		result, err := randFloat64()
		if err != nil {
			t.Error(err)
			return false
		}

		if result < 0 {
			t.Errorf("Value %v < 0", result)
			return false
		}

		if result >= 1 {
			t.Errorf("Value %v >=1", result)
			return false
		}

		if result < low {
			low = result
		}
		if result > high {
			high = result
		}

		return true
	}, nil)

	if low > .01 {
		t.Fatalf("Low value %v is too high", low)
	}
	if high < 0.99 {
		t.Fatalf("High value %v is too low", high)
	}
}

// func TestRand(t *testing.T) {
// 	c := 10
// 	b := make([]byte, c)
// 	_, err := rand.Read(b)
// 	CheckError(err)

// 	fmt.Println(b) // [193 174 178 8 109 249 203 255 176 153]
// }

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
