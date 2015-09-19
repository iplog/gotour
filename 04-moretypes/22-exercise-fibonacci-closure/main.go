// Implement a fibonacci function that returns a function (a closure) that
// returns successive fibonacci numbers.
// Fibonacci
// f(0) -> 1
// f(1) -> 1
// f(2) -> f(0) + f(1) = 2
// f(3) -> f(1) + f(3) = 3
// A smart way to

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	cur, next := 0, 1
	return func() int {
		cur, next = next, cur+next
		return cur
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
