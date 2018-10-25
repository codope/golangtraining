package main

import (
	"fmt"
	"sync"
	"time"
)

var Wait sync.WaitGroup = sync.WaitGroup{}
var Counter int = 0
var Lock sync.Mutex = sync.Mutex{}

// go run -race racecondition.go
func main() {

	for routine := 1; routine <= 2; routine++ {

		Wait.Add(1)
		go Routine(routine)
	}

	Wait.Wait()
	fmt.Printf("Final Counter: %d\n", Counter)
}

func Routine(id int) {

	for count := 0; count < 2; count++ {

		//Lock.Lock()

		value := Counter
		time.Sleep(1 * time.Nanosecond)
		value++
		fmt.Println("Incremented value.")
		Counter = value

		//Lock.Unlock()
	}

	Wait.Done()
}
