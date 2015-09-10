// There are no proper assertions in Go, just an error method
// we need to list the cases and loop other to stay at least a bit dry
// --> Using the _Testify_ lib is gonna gfix that but let's keep it simple
// stupid for now

// To run the test do
// ```shell
// cd 03-flow-control/05-if
// go test
// ```
// Reports are quite ugly...
package main

import (
	"testing"
)

func TestSqrt(t *testing.T) {
	cases := []struct {
		in   float64
		want string
	}{
		{4, "2"},
		{-4, "2i"},
		{9, "3"},
	}

	for _, c := range cases {
		got := sqrt(c.in)
		if got != c.want {
			t.Errorf("sqrt(%f) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{-1, "buzzfizz"},
		{0, "fizzbuzz"},
		{6, "fizz"},
		{10, "buzz"},
		{15, "fizzbuzz"},
		{17, ""},
	}

	for _, c := range cases {
		got := fizzBuzz(c.in)
		if got != c.want {
			t.Errorf("sqrt(%f) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestPow(t *testing.T) {
	cases := []struct {
		in   []float64
		want float64
	}{
		{[]float64{2, 2, 10}, 4},
		{[]float64{2, 2, 2}, 2},
	}

	for _, c := range cases {
		got := pow(c.in[0], c.in[1], c.in[2])
		if got != c.want {
			t.Errorf("sqrt(%f, %f, %f) == %f, want %f", c.in[0], c.in[1], c.in[2], got, c.want)
		}
	}
}
