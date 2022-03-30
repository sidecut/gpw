package main

import (
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
	var low, high, sum float64
	var count int
	low = 1
	high = 0

	quick.Check(func() bool {
		count++
		result, err := randFloat64()
		if err != nil {
			t.Error(err)
			return false
		}
		sum += result
		t.Log(result)

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

	if low > .1 {
		t.Fatalf("Low value %v is too high", low)
	}
	if high < 0.9 {
		t.Fatalf("High value %v is too low", high)
	}

	avg := sum / float64(count)
	t.Logf("Average is %v", avg)
	if avg < .45 || avg > .55 {
		t.Fatalf("Average %v is out of range", avg)
	}
}
