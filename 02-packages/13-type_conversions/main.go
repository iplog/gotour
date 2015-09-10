// Type conversion:
//
// Remember that casting is mandatory

package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)
	fmt.Println(x, y, z)

	// Integer parsing
	var i int
	var j int64
	i, _ = strconv.Atoi("123")
	fmt.Println(i)

	// This is another way to Parse an Int64  `i, _ = strconv.ParseInt("1123", 10, 0)`
	j, _ = strconv.ParseInt("123", 10, 0)
	fmt.Println(int(j))

	// String to date: we alreday did that in the 01-welcome/02-playground-dates
	var layout string = "2006-02-01 15:04:05"
	var dateStr string = "2015-09-12 12:05:03"

	var date time.Time
	date, _ = time.Parse(layout, dateStr)

	fmt.Println("date", date)

	// Cast integer to String
	i = 12
	var str string = strconv.Itoa(i)

	// Print the string
	fmt.Println("integer cast to", str)
}
