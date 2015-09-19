package main

import (
	"math"
	"testing"
)

func TestVertexAbs(t *testing.T) {
	cases := []struct {
		in   Vertex
		want float64
	}{
		{Vertex{}, math.Sqrt(0)},
		{Vertex{1, 1}, math.Sqrt(2)},
		{Vertex{1, 2}, math.Sqrt(5)},
		{Vertex{2, 2}, math.Sqrt(8)},
	}

	for _, c := range cases {
		got := c.in.Abs()
		if got != c.want {
			t.Errorf("v.Abs() == %f for %f, got %f, want %f", got, c.in, got, c.want)
		}
	}
}

func TestMyFloatAbs(t *testing.T) {
	cases := []struct {
		in, want MyFloat
	}{
		{2, 2},
		{-1, 1},
		{0, 0},
	}

	for _, c := range cases {
		got := MyFloat.Abs(c.in)
		if got != float64(c.want) {
			t.Errorf("MyFloat.Abs(%f) == %f, got %f, want %f", c.in, got, got, c.want)
		}
	}
}
