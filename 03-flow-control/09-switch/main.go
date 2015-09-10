package main

import (
	"fmt"
	"runtime"
	"time"
)

// The switch cases are checkecked in order top to bottom. By default  a case
// body breaks automatically, unless it ends with a fallthrough statement.
// By specifying `fallthrough` all the cases will be executed.
func swithcFallthrough(i int) (str string) {
	var j int = 8
	switch i {
	case 4, 5: // We can use a coma to separate the cases
		str += "45"
		fallthrough
	case 6:
		str += "6"
		fallthrough
	case 7:
		str += "7"
		fallthrough
	case j: // We can use a variable directly in a case (or a function return)
		str += "8"
		fallthrough
	default:
		str += "default"
	}
	return
}

// Switch without a condition is the same as switch true.
// This construct can be a clean way to write long if-then-else chains.
func switchNoCondition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// // Go is checking for the type of the case. It has to be the same that the value
// // we do the `switch` on
// func switchType() {
//   var i int = 6
//   switch i {
//   case i < 7:
//     fmt.Println("hello")
//   case i == 7:
//     fmt.Println("world")
//   case i > 7:
//     fmt.Println("hello world")
//   }
// }

func main() {
	fmt.Print("Go runs on ")

	// Notice that you can set the `os` variable directly in the `switch` as you
	// can do on the `if`.
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}
