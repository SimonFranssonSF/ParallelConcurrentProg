package main

import (
	"fmt"
)

// This program should go to 11, but sometimes it only prints 1 to 10.
// Problem: Channel ch gets closed when last message is received but
// main exits before function Print is able to print last piece of data.
// FIX: Make sure all values are printed before main exits.
// Eg. set a time.delay, scan from standard input, sync package (use waitgroup) or communication with channels.
// Here, using communication with channels, exit main when done stops blocking (after 11 is printed stdout).
func main() {
	ch := make(chan int)
	//channel to be used for synchronizing main and goroutine
	done := make(chan bool)
	go Print(ch, done)
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	fmt.Println("Väntar på utskrift ska bli klar...")
	close(ch) //closes last channel when last message is sent
	<-done
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
// Added a parameter channel bool which will be closed when this function is finished
// and therefore tell the main function that it's done.
func Print(ch <-chan int, done chan<- bool) {
	for n := range ch {
		fmt.Println(n)
	}
	close(done) // Close channel done when last msg is printed to stdout
}
