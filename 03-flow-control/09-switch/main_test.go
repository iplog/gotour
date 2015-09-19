package main

import "testing"

func TestSwitchFallthrough(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{4, "45678default"},
		{1, "default"},
		{8, "8default"},
	}

	for _, c := range cases {
		got := swithcFallthrough(c.in)
		if got != c.want {
			t.Errorf("swithcFallthrough(%d) == %s, want %s", c.in, got, c.want)
		}
	}
}
