package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	if (t.Left != nil) {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if (t.Right != nil) {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
	var ch1 chan int = make(chan int, 10)
	var ch2 chan int = make(chan int, 10)

	go func(){
		Walk(t1, ch1)
		close(ch1)
	}()
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if ok1 != ok2{
			return false
		}

		if !ok1 {
			break
		}
		if (v1 != v2){
			return false
		}
	}

	return true
}

func cEquavilantBinarytrees() {
	fmt.Printf("%v\n", Same(tree.New(1), tree.New(1)));
	fmt.Printf("%v\n", Same(tree.New(1), tree.New(2)));
}
