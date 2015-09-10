// Variables
// The var statement declares a list of variables; as in function argument lists, the type is last.
// A var statement can be at package or function level. We see both in this example.

// Important: if not directly initialized a variable requires to be typed
// Good:
// var j int
// var j = 12 -> the varianble will take the type of the initial value (int in this case)
// Not working
// var j
// j = 12

package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	var j = 12
	fmt.Println(i, c, python, java)
	fmt.Println(j)
}
