package main

import (
	"fmt"
	"time"
)

func f1(c chan int) {
	i := 0
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			i++
			fmt.Println("Generated", i, " ... going to generate new value soon.")
			c <- i
		}
	}
}

func main() {
	c := make(chan int)

	go f1(c)

	for v := range c {
		fmt.Println("Received ", v)
		if v == 5 {
			fmt.Println("We got 5, and we're done.")

			// if the program didn't end here, the goroutine would still be running
			time.Sleep(10 * time.Second)
			break
		}
	}
}
