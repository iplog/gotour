// Zero Values: Easy thing to remember:
// Variables declared without an explicit initial value are given their zero value.
// What about the `nil` value then?

package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// This would be totally wrong
	// var j int = nil
	// var g float64 = nil
	// var c bool = nil
	// var t string = nil
	// fmt.Printf("%v %v %v %q\n", j, g, c, t)

	// nil is a predeclared identifier representing the zero value for a pointer,
	// channel, func, interface, map, or slice type.
	// we cannot do test like `i == nil` as we would do in other language to check
	// if the value has been initialized.
	// --> Be careful about the test we do.
}
