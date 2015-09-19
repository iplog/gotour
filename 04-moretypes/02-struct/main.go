// A struct is a collection of fields.
// (And a type declaration does what you'd expect.)
// That's the closest we'll go from Java Classes
package main

import "fmt"

type Vertex struct {
	X int
	Y int
	// This could be simply written X, Y int
}

func main() {
	v := Vertex{1, 2} // the delcaration order is respected X <- 1 and Y <- 2
	fmt.Println(v)

	// Struct fields are accessed using a dot.
	fmt.Println("v.X ==", v.X) // --> 1
	v.X = 4
	fmt.Println("v.X ==", v.X) // --> 4

	// Pointers to structs
	// Struct fields can be accessed through a struct pointer.
	// The indirection through the pointer is transparent.
	p := &v
	p.X = 42
	fmt.Println("v.X ==", v.X) // --> 42

	// But we can get a pointer to a struct field
	var q *int = &v.X
	fmt.Println(q)  // --> some memory address
	fmt.Println(*q) // --> 42

	// Other ways to declare a Vertex
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit --> Fields default to their "O value"
		v3 = Vertex{}      // X:0 and Y:0
		r  = &Vertex{1, 2} // has type *Vertex
	)

	fmt.Println(v1, r, v2, v3)
}
