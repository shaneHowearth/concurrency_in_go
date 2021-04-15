// Package main -
package main

import (
	"fmt"

	"github.com/shanehowearth/concurrency_in_go/pipeline"
)

func main() {
	done := make(chan interface{})
	defer close(done)

	intStream := pipeline.Generator(done, 1, 2, 3, 4)
	p := pipeline.Multiply(done, pipeline.Add(done, pipeline.Multiply(done, intStream, 2), 1), 2)

	for v := range p {
		fmt.Println(v)
	}
}
