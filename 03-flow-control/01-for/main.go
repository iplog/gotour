package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// Removes pre and post statement and get a while \o/
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// Creates an infinite loop. Use ctrl+c to quit
	for {
	}
}
