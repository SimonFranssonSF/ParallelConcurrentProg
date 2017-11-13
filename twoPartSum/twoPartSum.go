package main

import "fmt"

// Add adds the numbers in a and sends the result on res.
func Add(a []int, res chan<- int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	res <- sum // Send sum to res
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	n := len(a)
	ch := make(chan int)
	go Add(a[:n/2], ch)
	go Add(a[n/2:], ch)

	a1, a2 := <-ch, <-ch // Receive from res
	fmt.Println("The sum of the values in all elements in the array:", a, "is:", a1+a2)
}
