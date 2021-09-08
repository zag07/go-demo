package main

import (
	"fmt"

	"github.com/bits-and-blooms/bloom/v3"
)

// this is from bloom_test.go
func main() {
	f := bloom.NewWithEstimates(1000, 0.001)
	n1 := "Love"
	n2 := "is"
	n3 := "in"
	n4 := "bloom"
	f.AddString(n1)
	n3a := f.TestAndAddString(n3)
	n1b := f.TestString(n1)
	n2b := f.TestString(n2)
	n3b := f.TestString(n3)
	f.TestString(n4)
	if !n1b {
		fmt.Printf("%v should be in.\n", n1)
	}
	if !n2b {
		fmt.Printf("%v should not be in.\n", n2)
	}
	if !n3a {
		fmt.Printf("%v should not be in the first time we look.\n", n3)
	}
	if !n3b {
		fmt.Printf("%v should be in the second time we look.\n", n3)
	}
}
