// Nothing really complicated here. we jsut have to noticie that the types is
// coming after the parameter name.
// This is a big change compared to Java and C
// Read [Go's Declaration Syntax](http://blog.golang.org/gos-declaration-syntax)
// for more information
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// Let's add multiply function
func multiply(x int, y int) int {
	return x * y
}

// A shorter notation would be
// func multiply(x, y int) int {
//   return x * y
// }
// In this case we still mean that x and y are `int`. It seems a bit harder to
// read.

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(multiply(42, 13))
}
