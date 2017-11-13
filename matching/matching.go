// http://www.nada.kth.se/~snilsson/concurrency/
package main

import (
	"fmt"
	"sync"
)

/*** Questions ***/
/*
What happens if you remove the go-command from the Seek call in the main function?
No concurrency, hence, the order of who sent and received on channel match is sequential and becomes static,
the same result every time. As match is a buffered channel the program will not cause deadlock.

What happens if you switch the declaration wg := new(sync.WaitGroup) to var wg sync.WaitGroup and the parameter wg *sync.WaitGroup to wg sync.WaitGroup?
Deadlock, as there is no pointer to the waitgroup wg in main, the seek function handles copies of the wg,
hence the seek function never tells when goroutines in wg are done, they only tell copies which isn't blocking main.

What happens if you remove the buffer on the channel match?
Deadlock occurs (when no receiver), the last goroutine will wait for a receiver instead of putting data in buffer.
There will not be a receiver as the number of goroutines is uneven so  wg.done() will never be called
and the program gets stuck at wg.Wait() in main. If the number of people would've been even it would've still worked though.

What happens if you remove the default-case from the case-statement in the main function?
When everyone in people receives or sent a message a deadlock will be caused as main will wait to receive on channel
without any default case but there will not be sent anything else on channel match. with a default case select statement is not blocking
and without a default case the select statement is blocking when handeling channels.
*/

// This programs demonstrates how a channel can be used for sending and
// receiving by any number of goroutines. It also shows how  the select
// statement can be used to choose one out of several communications.
func main() {
	people := []string{"Anna", "Bob", "Cody", "Dave", "Eva"}
	match := make(chan string, 1) // Make room for one unmatched send.
	wg := new(sync.WaitGroup)
	wg.Add(len(people))
	for _, name := range people {
		go Seek(name, match, wg)
	}
	wg.Wait()
	select {
	case name := <-match:
		fmt.Printf("No one received %sâ€™s message.\n", name)
	default:
		// There was no pending send operation.
	}
}

// Seek either sends or receives, whichever possible, a name on the match
// channel and notifies the wait group when done.
func Seek(name string, match chan string, wg *sync.WaitGroup) {
	//time.Sleep(time.Duration(20+rand.Intn(10)) * time.Millisecond) // Used to demonstrate the function of the program.
	select {
	case peer := <-match:
		fmt.Printf("%s sent a message to %s.\n", peer, name)
	case match <- name:
		// Wait for someone to receive my message.
	}
	wg.Done()
}

/*
func init() {
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}
*/
