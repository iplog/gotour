package main

import (
	"fmt"
	"math"
)

func main() {
	// Functions are values
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))

	// Go functions may be closures. A closure is a function value that
	// references variables from outside its body. The function may access and
	// assign to the referenced variables; in this sense the function is "bound"
	// to the variables.
	// For example, the adder function returns a closure. Each closure is bound
	// to its own sum variable.
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func adder() func(int) int {
	sum := 0 // the returned fucntion will be bound to `sum`
	// sum WILL NOT BE shared by the different returned functions
	return func(x int) int {
		sum += x
		return sum
	}
}
