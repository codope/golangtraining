package main

import "fmt"

//import "time"

var c = make(chan string)
var quit = make(chan bool)

func main() {
	go fn()

	for {
		select {
		case msg, ok := <-c:
			if !ok {
				c = nil
				return
			} else {
				fmt.Println("received: ", msg)
				quit <- true
			}
		}
	}
}

func fn() {
	for {
		select {
		case <-quit:
			fmt.Println("closin chan c")
			close(c)
			return
			//case <-time.After(time.Nanosecond):  // this is ok
		default: // using default causes deadlock
			c <- "Image"
		}
	}
}
