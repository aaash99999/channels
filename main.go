package main

import (
	"fmt"
	"time"
)

func main() {
	// We're running a virtual pizza shop to learn about goroutines and channels.  e.g. each worker will do
	// some task related to making a pizza, then communicate their update to the next worker.

	fmt.Printf("This is the main goroutine.\n")

	// ch will be how we communicate updates between goroutines
	ch := make(chan string)
	defer close(ch)

	go listen(ch) // listen will be our central fuction that listens for updates on the channel and reports them

	for i := 1; i < 5; i++ {
		// prepare 5 pizza bases
		go prepareBase(i, ch)
	}

	// sleep while stuff is getting done so the program doesn't exit too early
	fmt.Println("We'll sleep here for a few seconds to let other stuff happen concurrently, and to avoid our main goroutine exiting before the pizzas are ready.")
	time.Sleep(5 * time.Second)

	// exit the program
	fmt.Printf("This is the main goroutine exitting.\n")
}

// prepare the base of a pizza with pizza sauce
func prepareBase(n int, ch chan (string)) {
	tmp := fmt.Sprintf("Preparing pizza base #%d\n", n)
	ch <- tmp // send our status update to the channel
}

func listen(ch chan (string)) {
	fmt.Printf("Listening for updates in the listen function...\n")
	<-ch // TODO: I think this is blocking and screwing things up...
	fmt.Printf("Got an update!\n")
	fmt.Printf("Update: %s\n", <-ch)
}
