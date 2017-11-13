package main

import (
	"math"
	"testing"
)

func TestWordCount(t *testing.T) {
	// Answer to Sqrt(2)
	answer := math.Sqrt(2)
	q, _, _ := SqrtDelta(2, 0)

	// If they aren't equal, that is delta == 0 then error is displayed since function didn't do as intended.
	if q != answer {
		t.Error("not enough precision, expected", answer, "got ", q)
	}
}
