package main

import (
	"fmt"
	"math"
)

// initialized the delta const
const delta = 1e-6

func Sqrt(x float64) float64 {
	var z float64 = x
	var l float64 // last computed value (initialized at 0.0 by default)

	for math.Abs(l-z) > delta { // This is a while loop
		l = z
		z = z - (z*z-x)/(2*z)

		// the 2 lines above could be written this way
		// l, z = z, z-(z*z-x)/(2*z)
		// But it seems a bit harder to read
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(9))
}
