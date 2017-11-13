package main

import (
	"testing"
	"time"
)

// Creates a 3 by 3 picture with gray/blue values that are to be compared with the intended result.
// Displays error if the pic doesn't contain the intended data.
func RemindTest(t *testing.T) {
	
	go Remind("Test 1", 2*time.Second)

	select { }

}
