package main

import "fmt"

func main() {
	// try with no buffer and see the diff. HINT: deadlock!
	messages := make(chan string, 1)
	messages <- "ping"

	go func() {
		v := <-messages
		fmt.Println(v)
	}()

	messages <- "ping2"
	fmt.Println(<-messages)
}
