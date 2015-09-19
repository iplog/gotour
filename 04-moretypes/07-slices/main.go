// A slice points to an array of values and also includes a length.
// []T is a slice with elements of type T.

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}

	// another way to loop other a slice. Would also work with an array
	for j, v := range s {
		fmt.Printf("s[%d] == %d\n", j, v)
	}

	// Slices can be re-sliced, creating a new slice value that points to the same array.
	// The expression
	// ```
	// s[lo:hi]
	// ```
	// evaluates to a slice of the elements from lo through hi-1, inclusive. Thus
	// ```
	// s[lo:lo]
	// ```
	// is empty and
	// ```
	// s[lo:lo+1]
	// ```
	// has one element.

	fmt.Println("s ==", s)
	fmt.Println("s[1:4] ==", s[1:4]) // Looks quite pythonic :)

	// missing low index implies 0
	fmt.Println("s[:3] ==", s[:3])

	// missing high index implies len(s)
	fmt.Println("s[4:] ==", s[4:])

	// overflow --> will raise `panic: runtime error: slice bounds out of range`
	// fmt.Println("s[4:] ==", s[4:12])

	// Slices are created with the make function. It works by allocating a zeroed
	// array and returning a slice that refers to that array:
	// ```
	// a := make([]int, 5)  // len(a)=5
	// ```
	// To specify a capacity, pass a third argument to make:
	// ```
	// b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	// b = b[:cap(b)] // len(b)=5, cap(b)=5
	// b = b[1:]      // len(b)=4, cap(b)=4
	// ```
	a := make([]int, 5)
	printSlice("a", a) // --> a len=5 cap=5 [0 0 0 0 0]
	b := make([]int, 0, 5)
	printSlice("b", b) // --> b len=0 cap=5 []
	c := b[:2]
	printSlice("c", c) // --> c len=2 cap=5 [0 0]
	d := c[2:5]
	printSlice("d", d) // --> d len=3 cap=3 [0 0 0]

	// The zero value of a slice is nil.
	// A nil slice has a length and capacity of 0.
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}

	// Adding elements to a slice
	// `append` is the built-in function  you just need
	a = []int{}
	printSlice("a", a) //a len=0 cap=0 []

	// append works on nil slices.
	a = append(a, 0)
	printSlice("a", a) // a len=1 cap=1 [0]

	// the slice grows as needed.
	a = append(a, 1)
	printSlice("a", a) // a len=2 cap=2 [0 1]

	// we can add more than one element at a time.
	a = append(a, 2, 3, 4)
	printSlice("a", a) // a len=5 cap=6 [0 1 2 3 4]

	// Adding more element
	a = append(a, 10, 12, 13)
	printSlice("a", a) // a len=8 cap=12 [0 1 2 3 4 10 12 13]

	// The `cap` auto increases
}

func printSlice(s string, x []int) {
	// `cap` is a really important built-in function
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
