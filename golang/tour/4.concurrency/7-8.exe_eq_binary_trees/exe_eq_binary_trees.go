package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		go Walk(t.Left, ch)
	}
	if t.Right != nil {
		go Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch := make(chan int)
	vals := map[int]bool{}

	go Walk(t1, ch)
	for i := 0; i < 10; i++ {
		v := <-ch
		vals[v] = true
	}

	go Walk(t2, ch)
	for i := 0; i < 10; i++ {
		v := <-ch
		if _, ok := vals[v]; ok == false {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
