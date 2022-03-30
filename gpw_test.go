package main

import (
	"testing"
	"testing/quick"
)

func TestGpw(t *testing.T) {
	pwl := 12
	quick.Check(func() bool {
		pwd, err := pronounceable(pwl)
		if err != nil {
			t.Error(err)
			return false
		}

		if len(pwd) != pwl {
			t.Errorf("Password %v is the wrong length %v", pwd, len(pwd))
			return false
		}

		return true
	}, nil)
}
