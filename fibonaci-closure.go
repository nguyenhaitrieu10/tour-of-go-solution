// Exercise: Fibonacci closure
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prev := -1
	cur := 1
	return func() int {
		prev, cur = cur, prev+cur
		return cur
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
