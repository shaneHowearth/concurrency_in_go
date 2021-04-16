// Package main -
package main

import (
	"fmt"

	"github.com/shanehowearth/concurrency_in_go/bridge"
)

func main() {
	genVals := func() <-chan <-chan interface{} {
		chanStream := make(chan (<-chan interface{}))
		go func() {
			defer close(chanStream)
			for i := 0; i < 10; i++ {
				stream := make(chan interface{}, 1)
				stream <- i
				close(stream)
				chanStream <- stream
			}
		}()
		return chanStream
	}

	for v := range bridge.Bridge(nil, genVals()) {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}
