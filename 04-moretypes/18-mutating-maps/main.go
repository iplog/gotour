// Insert or update an element in map m:
// ```
// m[key] = elem
// ```
// Retrieve an element:
// ```
// elem = m[key]
// ```
// Delete an element:
// ```
// delete(m, key)
// ```
// Test that a key is present with a two-value assignment:
// ```
// elem, ok = m[key]
// ```
// If key is in m, ok is true. If not, ok is false and elem is the zero value for the map's element type.

// Similarly, when reading from a map if the key is not present the result is the zero value for the map's element type.

// Note: if elem or ok have not yet been declared you could use a short declaration form:
// ```
// elem, ok := m[key]
// ```

package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"]) // --> 42

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"]) // --> 48

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"]) // --> 0

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok) // --> 0, false
}
