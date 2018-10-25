package main

import (
	"fmt"
	"sync"
)

var (
	rwMutex sync.RWMutex // guards bal
	bal     int
)

func Deposit2(amount int) {
	rwMutex.Lock()
	bal = bal + amount
	// If you uncomment the below statement, you will see the
	// 	update happening as expected and in certain order.
	// fmt.Println("Current bal inside: ", bal)
	rwMutex.Unlock()
}

func Balance2() int {
	rwMutex.RLock() // readers lock
	defer rwMutex.RUnlock()
	return bal
}

func main() {
	// Deposit2 [1..1000] concurrently.

	max := 10
	var n sync.WaitGroup
	for i := 1; i <= max; i++ {
		n.Add(1)
		go func(amount int) {
			Deposit2(amount)
			// If you uncomment the below statement, you will see the
			//	that the actual reading of the data is uncertain.
			// fmt.Println("Current bal outside: ", Balance2())
			n.Done()
		}(i)
	}
	n.Wait()

	got, want := Balance2(), (max+1)*max/2
	if got != want {
		fmt.Errorf("Balance2 = %d, want %d", got, want)
	} else {
		fmt.Printf("All Ok! Balance2 = %d, want %d\n", got, want)
	}
}
