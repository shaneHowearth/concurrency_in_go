// Package main -
package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

func main() {
	fmt.Println("vim-go")
	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b) // Locks a, then 2 seconds later attempts to lock b
	go printSum(&b, &a) // Locks b, then 2 seconds later attempts to lock a
	// Both the above goroutines are trying to acquire locks on values that the
	// other goroutine has already locked, and they stay locked in that death
	// spiral forever. Note that the goroutines are using pointers to a and b,
	// meaning that both goroutines are using the same 'values'.
	wg.Wait()

}
