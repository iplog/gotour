package main

import (
	"fmt"
	"math"
)

// You could also use this notation
// import "fmt"
// import "math"
// However it looks like `go import` use the above one by default so that's
// what we'll use.

func main() {
	fmt.Printf("Now you have %g problems.", math.Nextafter(2, 3))
}
