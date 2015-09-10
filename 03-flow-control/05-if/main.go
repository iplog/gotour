// We start implementing some small function so it's time to check if everythig
// is working.
// Go does not have assertions so we'll directly add unit tests \o/.
package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	// here we should maybe use `strconv.Itoa`
	return fmt.Sprint(math.Sqrt(x))
}

// Simple fizz buzz implementation to check the `else` syntax.
func fizzBuzz(n int) (str string) {
	if n < 0 {
		return "buzzfizz"
	} else {
		if n%3 == 0 {
			str += "fizz"
		}
		if n%5 == 0 {
			str += "buzz"
		}
	}
	return
}

// if with a short statement
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(fizzBuzz(15))
	fmt.Println(fizzBuzz(3))
	fmt.Println(fizzBuzz(5))
	fmt.Println(fizzBuzz(-5))
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}
