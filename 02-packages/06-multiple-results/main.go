// Basic example of a function returning multiple values.
// In python you would use a tuple and in JS an object or eventually a set in
// ES6.
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
