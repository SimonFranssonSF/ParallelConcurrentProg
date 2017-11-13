// Stefan Nilsson 2013-03-13

// This is a testbed to help you understand channels better.
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

/*** DESCRIPTION OF WHAT HAPPENS ***/
//
// TEST 1: What happens if you switch the order of the statements wgp.Wait() and close(ch)
// in the end of the main function?
// HYPOTHESIS TEST 1: I think that main will close the channel before all data is sent
// on the channel, when the channel is closed the consumers breaks their for loop and
// notifies the main goroutine that it's "done", exiting the program to early. (FALSE HYPO)
// CHANGE TEST 1: "panic: send on closed channel"
// CONCLUTION TEST 1: The channel is closed immediately, then producers are trying to
// send on a closed channel causing a panic.
//
// TEST 2: What happens if you move the close(ch) from the main function and instead
// close the channel in the end of the function Produce?
// HYPOTHESIS TEST 2: I think that the producer who ever sends all his data first will
// close the channel ch and causing a panic as the other producers still are trying
// to send data on the channel
// CHANGE TEST 2:  "panic: send on closed channel"
// CONCLUSION TEST 2: Panic caused when the producer who sent all his 8 strings
// first finishes. As channel then gets closed while the other producers still trying
// to send data on that channel.
//
// TEST 3: What happens if you remove the statement close(ch) completely?
// HYPOTHESIS TEST 3: The consumers range loop will never exit causing a deadlock
// (TURNED OUT TO BE WRONG).
// CHANGE TEST 3: The main sometimes exits before all messages are printed  more often.
// CONCLUSION TEST 3: The consumer functions will never finish as no close(ch) tells
// the for range channel is closed. And close(ch) seems to make the the main exit a
// little bit later which gives the consumer function a greater chance of atleast
// printing all strings before main finishes.
//
// TEST 4: What happens if you increase the number of consumers from 2 to 4?
// HYPOTHESIS TEST 4: Program might become faster, 4 strings might not be printed
// before main exits now.
// CHANGE TEST 4: Program became faster.
// CONCLUSION TEST 4: Program ran faster. But if you put a delay before the
// "printing line" in function consumer you will see that 4 messages don't get
// printed before main exits.
//
// TEST 5: Can you be sure that all strings are printed before the program stops?
// HYPOTHESIS TEST 5: No.
// CHANGE TEST 5: Put a delay before consumer prints the strings received in their function
// CONCLUSION TEST 5: The last strings doesn't get printed before main exits.
// Program waits until all data is sent but it does not wait until all data is
// 100% handled in consumer function, main exits before if unlucky. Chance of 2
// strings (if 2 consumers) not being printed before main exits.

/*** Own notes for interpreting the program
 Program waits until all data is sent but it does not wait until all data is
 100% handled in consumer function, main exits before if unlucky. Chance of 2
 strings (if 2 consumers) not being printed before main exits.

 1. The producers divide the work equally and sends 8 strings each on channel ch
 2. The consumers doesn't receive the 32 strings equally they just work until
    told otherwise, eg. main exits or  channels are closed.
 3. The order of how they receive the strings are different
 4. It is an unbuffered channel so a receiver needs to be ready before the sender
    can continue (it blocks itself til string is sent)
 5. The for range in both consumr functions does not break until close(ch) is called
***/
func main() {
	// Use different random numbers each time this program is executed.
	rand.Seed(time.Now().Unix())

	const strings = 32
	const producers = 4
	const consumers = 4

	before := time.Now()
	ch := make(chan string)
	wgp := new(sync.WaitGroup)
	wgp.Add(producers)
	wgc := new(sync.WaitGroup) // Wait group making sure all consumers are done
	wgc.Add(consumers)         // Add number of consumers to wg
	for i := 0; i < producers; i++ {
		go Produce("p"+strconv.Itoa(i), strings/producers, ch, wgp)
	}
	for i := 0; i < consumers; i++ {
		go Consume("c"+strconv.Itoa(i), ch, wgc)
	}

	wgp.Wait() // Wait for all producers to finish.
	close(ch)  // Change for time.Sleep(time.Second and try print something at the end of consumer it will never be reached)
	wgc.Wait() // Wait for all consumers to finish.
	fmt.Println("time:", time.Now().Sub(before))
}

// Produce sends n different strings on the channel and notifies wg when done.
func Produce(id string, n int, ch chan<- string, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		RandomSleep(100) // Simulate time to produce data.
		ch <- id + ":" + strconv.Itoa(i)
	}
	wg.Done()
}

// Consume prints strings received from the channel until the channel is closed.
func Consume(id string, ch <-chan string, wg *sync.WaitGroup) {
	for s := range ch {
		// time.Sleep(time.Second) // To see the effect of not having/having a wg for consumer to finish
		fmt.Println(id, "received", s)
		RandomSleep(100) // Simulate time to consume data.
	}
	//fmt.Println("kommer jag hit om jag tar bort close(ch)?")
	wg.Done()
}

// RandomSleep waits for x ms, where x is a random number, 0 â‰¤ x < n,
// and then returns.
func RandomSleep(n int) {
	time.Sleep(time.Duration(rand.Intn(n)) * time.Millisecond)
}
