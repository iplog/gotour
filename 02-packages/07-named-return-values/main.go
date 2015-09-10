// A quick exampel to show how to return named values
// the codign style is not consistent here (just meessing around to test)
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func add(x, y int) (sum int) {
	sum = x + y
	return
}

func main() {
	// `split` returns multiple named values
	var i, j int = split(17)
	fmt.Println(i, j)

	// `add` returns a single named value
	s := add(3, 4)
	fmt.Println(s)
}
