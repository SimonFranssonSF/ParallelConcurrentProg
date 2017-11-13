package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prevVal := 0
	currentVal := 1
	return func() int {
		fib := prevVal
		prevVal, currentVal = currentVal, prevVal+currentVal
		return fib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
