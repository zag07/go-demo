package main

import (
	"fmt"

	"github.com/bits-and-blooms/bitset"
)

// See more tests in bitset_test.go
func ExampleBitSet() {
	var b bitset.BitSet
	b.Set(8)
	fmt.Println("len", b.Len())
	fmt.Println("count", b.Count())
	fmt.Println(b.Test(3))
	fmt.Println(b.Test(8))
	// Output:
	// len 9
	// count 1
	// false
	// true
}
