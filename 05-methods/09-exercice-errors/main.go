// Copy your Sqrt function from the earlier exercise and modify it to return an
// error value.
// Sqrt should return a non-nil error value when given a negative number, as it
// doesn't support complex numbers.
// Create a new type
// ```
// type ErrNegativeSqrt float64
// ```
// and make it an error by giving it a
// ```
// func (e ErrNegativeSqrt) Error() string
// ```
// method such that ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative
// number: -2".
// Note: a call to fmt.Sprint(e) inside the Error method will send the program
// into an infinite loop. You can avoid this by converting e first:
// fmt.Sprint(float64(e)). Why?
// Change your Sqrt function to return an ErrNegativeSqrt value when given a
// negative number.
package main

import (
	"fmt"
	"math"
)

const delta = 1e-6

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Sqrt: negative number %g", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	var z float64 = x
	var l float64 // last computed value (initialized at 0.0 by default)

	for math.Abs(l-z) > delta { // This is a while loop
		l = z
		z = z - (z*z-x)/(2*z)

		// the 2 lines above could be written this way
		// l, z = z, z-(z*z-x)/(2*z)
		// But it seems a bit harder to read
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
