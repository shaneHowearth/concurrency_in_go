// Package main -
package main

import (
	"fmt"
	"sync"

	"github.com/shanehowearth/concurrency_in_go/replicatedrequests"
)

func main() {
	done := make(chan interface{})
	result := make(chan int)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go replicatedrequests.DoWork(done, i, &wg, result)
	}
	firstReturned := <-result
	close(done)
	wg.Wait()
	fmt.Printf("Received an answer from #%v\n", firstReturned)
}
