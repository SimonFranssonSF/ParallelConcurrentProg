package main

import "testing"

// Creates a 3 by 3 picture with gray/blue values that are to be compared with the intended result.
// Displays error if the pic doesn't contain the intended data.
func PicTest(t *testing.T) {
	answer := [][]uint8{{0, 0, 1}, {0, 1, 1}, {1, 1, 2}}
	q := Pic(3, 3)
	if len(answer) == len(q) {
		for i, v := range q {
			for i2, v2 := range v {
				if v2 != answer[i][i2] {
					t.Error("Expected", answer[i][i2], "got ", v2)
				}
			}
		}
	} else {
		t.Error("Expected len", len(answer), "got ", len(q))
	}
}
