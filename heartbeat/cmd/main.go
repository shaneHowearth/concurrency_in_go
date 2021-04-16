// Package main -
package main

import (
	"fmt"

	"github.com/shanehowearth/concurrency_in_go/heartbeat"
)

func main() {
	done := make(chan interface{})
	heartbeat, results := heartbeat.DoWork(done)
	for {
		select {
		case _, ok := <-heartbeat:
			if ok {
				fmt.Println("pulse")
			} else {
				return
			}
		case r, ok := <-results:
			if ok {
				fmt.Printf("results %v\n", r)
			} else {
				return
			}
		}
	}
}
