// Exercise: Equivalent Binary Trees
package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for a := range ch1 {
		b, ok := <-ch2
		if !ok || a != b {
			return false
		}
	}

	if _, ok := <-ch2; ok == true {
		return false
	}
	return true
}

func main() {
	// Test Walk
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for x := range ch {
		fmt.Print(x)
	}

	// Test Same
	fmt.Printf("\n%v", Same(tree.New(1), tree.New(1)))
	fmt.Printf("\n%v", Same(tree.New(1), tree.New(2)))
	fmt.Printf("\n%v", Same(nil, tree.New(2)))
	fmt.Printf("\n%v", Same(tree.New(1), nil))
}
