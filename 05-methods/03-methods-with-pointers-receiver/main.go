// Methods can be associated with a named type or a pointer to a named type.
// We just saw two Abs methods. One on the *Vertex pointer type and the other
// on the MyFloat value type.
// There are two reasons to use a pointer receiver. First, to avoid copying the
// value on each method call (more efficient if the value type is a large
// struct). Second, so that the method can modify the value that its receiver
// points to.
// Try changing the declarations of the Abs and Scale methods to use Vertex
// as the receiver, instead of *Vertex.
// The Scale method has no effect when v is a Vertex. Scale mutates v. When v
// is a value (non-pointer) type, the method sees a copy of the Vertex and
// cannot mutate the original value.
// Abs works either way. It only reads v. It doesn't matter whether it is
// reading the original value (through a pointer) or a copy of that value.
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) ScaleNoEffect(f float64) { // won't work (we do a copy and that's the copy that we modify)
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) AbsBis() float64 { // less efficient as we do a copy but will work
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs()) // has changed

	w := Vertex{3, 4}
	fmt.Printf("Before scaling: %v, Abs: %v\n", w, w.Abs())
	w.ScaleNoEffect(5)
	fmt.Printf("After scaling: %v, Abs: %v\n", w, w.Abs()) // -> same than before
}
