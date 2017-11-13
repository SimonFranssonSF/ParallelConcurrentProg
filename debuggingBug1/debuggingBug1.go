package main

import (
	"fmt"
)

// "I want this program to print "Hello World", but it doesn't work."
//
// PROBLEM: "Channels send/receive operations blocks until the other side is ready,"
// as the other side will not be ready since it's the same side (same goroutine) deadlock occurs.
//
// FIX: Make the channel buffered so that the channel isn't blocked until a
// receiver is ready. As the receiver is the main goroutine itself it will
// never be ready if it blocks itself when sending on unbuffered channel but
// on a buffered channel, it will just put that data in the buffere and continue running.
//
// Another way of fixing it would be to make another goroutine which can send
// "Hello world" in unbuffered channel ch so that goroutine 1 can receive from it.
func main() {
	ch := make(chan string, 1)
	ch <- "Hello World!"
	fmt.Println(<-ch)
}
