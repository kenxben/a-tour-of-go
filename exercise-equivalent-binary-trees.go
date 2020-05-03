package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	//fmt.Println(*t)

	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	} 
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	
	for i := 0; i < cap(ch1); i++  {
		if (<-ch1) != (<-ch2) {
			 return false
		}
	}
	return true
}

func main() {
	ch := make(chan int, 10)
	go Walk(tree.New(1), ch)
	//fmt.Println(*tree.New(1))
	for i := 0; i < cap(ch); i++  {
		//_ = i
		c := <-ch
		fmt.Println(c)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

