// Package main -
package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("vim-go")
	cadence := sync.NewCond(&sync.Mutex{})
	go func() {
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast()
		}
	}()

	takeStep := func() {
		cadence.L.Lock()
		cadence.Wait()
		cadence.L.Unlock()
	}

	// tryDir allows a person to try and move in a direction - returning success
	// or failure. Each direction is represented as a count of the number of
	// people trying to move in that direction.
	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
		fmt.Fprintf(out, " %v", dirName)
		// Declare our intention to move in direction
		atomic.AddInt32(dir, 1)
		// simulate constant cadence amongst all parties
		takeStep()
		if atomic.LoadInt32(dir) == 1 {
			fmt.Fprintf(out, ". Success!")
			return true
		}
		takeStep()
		// Step in direction was a failure, so decrement direction
		atomic.AddInt32(dir, -1)
		return false
	}

	var left, right int32
	tryLeft := func(out *bytes.Buffer) bool { return tryDir("left", &left, out) }
	tryRight := func(out *bytes.Buffer) bool { return tryDir("right", &right, out) }

	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer func() {
			fmt.Println(out.String())
		}()
		defer walking.Done()
		fmt.Fprintf(&out, "%v is trying to scoot:", name)
		for i := 1; i < 5; i++ { // artificial limit so the program isn't stuck forever
			if tryLeft(&out) || tryRight(&out) { // Attempt left, then right, and if they both fail "kill"
				fmt.Println("Killing!")
				return
			}
		}
		fmt.Fprintf(&out, "\n%v tosses her hands up in exasperation!", name)
	}

	var peopleInHallway sync.WaitGroup
	peopleInHallway.Add(2)
	go walk(&peopleInHallway, "Alice")
	go walk(&peopleInHallway, "Barbara")
	peopleInHallway.Wait()

}
