package main

import (
	"context"
	"fmt"
	"time"
)

func f2(ctx context.Context, c chan int) {
	i := 0
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			i++
			fmt.Println("Generated", i)
			c <- i
		case <-ctx.Done():
			ticker.Stop()
			fmt.Println("Ticker stopped. Closing channel.")
		}
	}
}

func main() {
	c := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	go f2(ctx, c)

	for v := range c {
		fmt.Println("Received ", v)
		if v == 5 {
			fmt.Println("We got 5, and we're done.")

			cancel()
			break

		}
	}
}

