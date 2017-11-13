package main

import "testing"

func TestWordCount(t *testing.T) {
	// Str with test data and expected result
	str := "A man a plan a canal panama."
	a := map[string]int{"A": 1, "man": 1, "a": 2, "plan": 1, "canal": 1, "panama.": 1}
	q := WordCount(str)

	// Testing the all key/value pairs from the string str,
	// function throws and error if it doesn't return the expected result
	for k, _ := range q {
		if q[k] != a[k] {
			t.Error("Expected", a[k], "got ", q[k], k)
		}
	}
}
