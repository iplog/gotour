// An interface type is defined by a set of methods.
// A value of interface type can hold any value that implements those methods.
package main

import (
	"fmt"
	"math"
	"os"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// A type implements an interface by implementing the methods. There is no
// explicit declaration of intent; no "implements" keyword.
// Implicit interfaces decouple implementation packages from the packages
// that define the interfaces: neither depends on the other.
// It also encourages the definition of precise interfaces, because you don't
// have to find every implementation and tag it with the new interface name.
type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// // In the following line, v is a Vertex (not *Vertex)
	// // and does NOT implement Abser.
	// a = v

	fmt.Println(a.Abs())

	var w Writer
	// os.Stdout implements Writer
	w = os.Stdout

	fmt.Fprintf(w, "hello, writer\n")
}
