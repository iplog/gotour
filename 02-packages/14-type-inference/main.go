package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	v := 42
	fmt.Printf("v is of type %T\n", v) // --> int

	w := 'a'
	fmt.Printf("w is of type %T\n", w) // --> int32. yep that's what chars are

	x := time.Now()
	fmt.Printf("x is of type %T\n", x) // --> time.Time

	// Another way to get the type is to use the reflect package
	var t reflect.Type
	t = reflect.TypeOf(v)
	fmt.Printf("v is of type %s\n", t) // --> int

	t = reflect.TypeOf(w)
	fmt.Printf("w is of type %s\n", t) // --> int32. yep that's what chars are

	t = reflect.TypeOf(x)
	fmt.Printf("x is of type %s\n", t) // --> time.Time
}
