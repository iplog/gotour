package main

import "fmt"
import "testing"

func TestStringer(t *testing.T) {
	cases := []struct {
		in   IPAddr
		want string
	}{
		{IPAddr{192, 168, 0, 1}, "192.168.0.1"},
		{IPAddr{127, 0, 0, 1}, "127.0.0.1"},
	}

	for _, c := range cases {
		got := fmt.Sprint(c.in)
		if got != c.want {
			t.Errorf("fmt.Sprint(%v) == %s, got %s, want %s", c.in, got, got, c.want)
		}
	}
}
