package main

import "fmt"

func main() {

	// We'll iterate over 2 values in the `queue` channel.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// This `range` iterates over each element as its
	// received from `queue`. Because we `close`d the
	// channel above, the iteration terminates after
	// receiving the 2 elements. If we didnt `close` it
	// we had block on a 3rd receive in the loop.
	for elem := range queue {
		fmt.Println(elem)
	}
}
