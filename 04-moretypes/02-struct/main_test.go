package main

import "testing"

func TestVertex(t *testing.T) {
	var v Vertex = Vertex{1, 2}
	if v.X != 1 {
		t.Errorf("v.X == %d, want %d, got %d", v.X, 1, v.X)
	}
	if v.Y != 2 {
		t.Errorf("v.Y == %d, want %d, got %d", v.Y, 2, v.Y)
	}
}
