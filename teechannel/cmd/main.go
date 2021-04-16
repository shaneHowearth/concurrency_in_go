// Package main -
package main

import (
	"fmt"

	"github.com/shanehowearth/concurrency_in_go/pipeline"
	"github.com/shanehowearth/concurrency_in_go/teechannel"
)

func main() {

	done := make(chan interface{})
	defer close(done)
	out1, out2 := teechannel.Tee(done, pipeline.Take(done, pipeline.Repeat(done, 1, 2), 4))
	for val1 := range out1 {
		fmt.Printf("out1: %v, out2: %v\n", val1, <-out2)
	}
}
