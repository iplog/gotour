package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func WalkNode(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	WalkNode(t.Left, ch)
	ch <- t.Value
	WalkNode(t.Right, ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkNode(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	chT1, chT2 := make(chan (int)), make(chan (int))
	go Walk(t1, chT1)
	go Walk(t2, chT2)

	for {
		v1, ok1 := <-chT1
		v2, ok2 := <-chT2

		if v1 != v2 || ok1 != ok2 {
			return false
		}
		if !ok1 && !ok2 {
			return true
		}
	}
}

func main() {
	fmt.Print("Checks if tree.New(1) == tree.New(1) ")
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}

	fmt.Print("Checks if tree.New(1) == tree.New(2) ")
	if Same(tree.New(1), tree.New(2)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}
}
