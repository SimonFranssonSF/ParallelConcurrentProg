package fibonacci

import "testing"

func TestFibonacci(t *testing.T) {
	// First 10 fibonacci numbers that I want to test
	a := [10]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}

	// Testing the fibonacci function throws and error if it doesn't return the expected result
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fib := f()
		if int(fib) != a[i] {
			t.Error("Expected", a[i], "got ", fib)
		}
	}
}
