package main

import "testing"

//Tests a case to see if the add functions works with goroutines
func addTest(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7} // sum equals 28
	expected := 28
	n := len(a)
	ch := make(chan int)
	go Add(a[:n/2], ch)
	go Add(a[n/2:], ch)

	a1, a2 := <-ch, <-ch // Receive from res

	answer := a1 + a2
	if answer != expected {
		t.Error("Something is wrong, answer aquired:", answer, "answer expected,", expected)
	}

}
