package main

import (
	"fmt"
	"sync"
)

var (
	mu    sync.Mutex // guards bal
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	// If you uncomment the below statement, you will see the
	// update happening as expected and in certain order.
	// fmt.Println("Current bal inside: ", bal)
	mu.Unlock()
}

// NOTE: incorrect!
// Mutex locks are not re-entrant.
// It's not possible to lock an already locked mutex.
// In the code shown, the Deposit function attempts to acquire the lock a second time.
// This leads to a deadlock and nothing can proceed.
func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()

	// Deposit2() also tries to acquire a lock
	Deposit(-amount)
	if Balance() < 0 {
		Deposit(amount)
		return false // insufficient funds
	}
	return true
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}

func main() {
	// Deposit2 [1..1000] concurrently.

	max := 10
	var n sync.WaitGroup
	for i := 1; i <= max; i++ {
		n.Add(1)
		go func(amount int) {
			Deposit(amount)
			// If you uncomment the below statement, you will see the
			//  that the actual reading of the data is uncertain.
			// fmt.Println("Current bal outside: ", Balance2())
			n.Done()
		}(i)
	}
	n.Wait()

	got, want := Balance(), (max+1)*max/2
	if got != want {
		fmt.Errorf("Balance2 = %d, want %d", got, want)
	} else {
		fmt.Printf("All Ok! Balance2 = %d, want %d\n", got, want)
	}
}
