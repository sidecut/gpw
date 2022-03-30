package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var pwl *int = flag.Int("pwl", 12, "Password length")
var n *int = flag.Int("n", 10, "Number of passwords")
var seed *int64 = flag.Int64("seed", time.Now().UnixNano(), "Random number seed")

func main() {
	flag.Parse()
	rand.Seed(*seed)

	for i := 0; i < *n; i++ {
		password := pronounceable(*pwl)
		fmt.Printf("%v\n", password)
	}
}
