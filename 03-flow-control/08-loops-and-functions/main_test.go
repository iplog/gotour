// Quick test to check that the result we get are in the correct delta
package main

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	cases := []float64{2, 4, 9}

	for _, c := range cases {
		got := Sqrt(c)
		want := math.Sqrt(c)
		if math.Abs(got-want) > delta {
			t.Errorf("sqrt(%f) == %f, want %f", c, got, want)
		}
	}
}
