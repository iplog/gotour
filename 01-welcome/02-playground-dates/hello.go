// This script should be to test the playgroud limitations.
// cf [Inside the Go Playground](http://blog.golang.org/playground)
// Let's play a bit with the `time` package.
// The way to define date format and layout to format/parse a date is really
// smart. The reference date you have to use is:
// `Mon Jan 2 15:04:05 MST 2006`
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground!")

	currentTime := time.Now()
	fmt.Println("The time is", currentTime)

	// Local timestamp
	fmt.Println("Local time is", currentTime.Local())

	// Format a date
	format := "01/02/2006 15:04:05"
	fmt.Println("Formatted date", currentTime.Format(format))

	// Parse a date
	layout := "2006-02-01 15:04:05"
	dateStr := "2015-09-10 14:19:23"

	date, _ := time.Parse(layout, dateStr)
	fmt.Println("Parsed date", date)
}
