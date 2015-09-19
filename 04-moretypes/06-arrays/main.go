// The type [n]T is an array of n values of type T.
// The expression
// var a [10]int
// declares a variable a as an array of ten integers.
// An array's length is part of its type, so arrays cannot be resized.
// This seems limiting, but don't worry; Go provides a convenient way of
// working with arrays.package main
package main

import "fmt"

func main() {
	var a [2]string // An array's length is part of its type, so arrays cannot be resized.
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	var b [2]int = [2]int{1, 2}
	fmt.Println(b[0], b[1])

	c := [3]int{11, 12, 13}
	fmt.Println(c[0], c[1])
}
