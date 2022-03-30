package main

import (
	"flag"
	"fmt"
	"os"
)

var pwl *int = flag.Int("pwl", 12, "Password length")
var n *int = flag.Int("n", 10, "Number of passwords")

func main() {
	flag.Parse()

	for i := 0; i < *n; i++ {
		password, err := pronounceable(*pwl)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			fmt.Printf("%v\n", password)
		}
	}
}
