// Another way to declare variable that is shorter
// Note: if you read the [spec](https://golang.org/ref/spec#Short_variable_declarations)
// you can notice that
// > Unlike regular variable declarations, a short variable declaration may
// > redeclare variables provided they were originally declared earlier in the
// > same block (or the parameter lists if the block is the function body) with
// > the same type, and at least one of the non-blank variables is new.
// > As a consequence, redeclaration can only appear in a multi-variable short
// > declaration. Redeclaration does not introduce a new variable; it just
// > assigns a new value to the original.
// and
// > Short variable declarations may appear only inside functions. In some
// > contexts such as the initializers for "if", "for", or "switch" statements,
// > they can be used to declare local temporary variables.
// Consequently: try to not use them everywhere
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
