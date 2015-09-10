// My first packeage
// Nothing much here. `go fmt` and `go import` are really friendly.

// One thing to notice about random number though:
// Go's math/rand package provides [pseudorandom number generation]
// (https://en.wikipedia.org/wiki/Pseudorandom_number_generator).
// But you can use a Seed to change the sequence of generated numbers.
// This behaviour can be handy when unit testing for example.
// You can use `time.Now().UTC().UnixNano()` as a seed to make it change.

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favortie number is ", rand.Intn(10))
	fmt.Println("Another random int", rand.Intn(10))

	// Try to Seed with 1
	rand.Seed(1)
	fmt.Println("First sequence")
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(10))
	}

	// Seed again with 1: Numbers are the same
	rand.Seed(1)
	fmt.Println("Second sequence")
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(10))
	}

	// Seed with 2
	rand.Seed(2)
	fmt.Println("Third sequence")
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(10))
	}
}
