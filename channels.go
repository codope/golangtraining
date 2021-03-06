package main

import "fmt"

func main() {

	// Create a new channel with `make(chan val-type)`.
	// Channels are typed by the values they convey.
	messages := make(chan string)

	// blocking example
	//messages <- "hello"
	//go func() {
	//	msg := <-messages
	//	fmt.Println(msg)
	//}()

	// _Send_ a value into a channel using the `channel <-`
	// syntax. Here we send `"ping"`  to the `messages`
	// channel we made above, from a new goroutine.
	go func() {
		messages <- "ping"
		messages <- "ping2"
		messages <- "ping3"
	}()
	go func() { messages <- "ping4" }()
	go func() { messages <- "ping5" }()

	// The `<-channel` syntax _receives_ a value from the
	// channel. Here we'll receive the `"ping"` message
	// we sent above and print it out.
	msg := <-messages
	fmt.Println(msg)
	msg = <-messages
	fmt.Println(msg)
	msg = <-messages
	fmt.Println(msg)
	msg = <-messages
	fmt.Println(msg)
	msg = <-messages
	fmt.Println(msg)
}
