// Package main -
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/shanehowearth/concurrency_in_go/fanin"
	"github.com/shanehowearth/concurrency_in_go/pipeline"
)

func main() {
	done := make(chan interface{})
	defer close(done)
	start := time.Now()
	rand := func() interface{} { return rand.Intn(50000000) }
	randIntStream := pipeline.ToInt(done, pipeline.RepeatFn(done, rand))
	numFinders := runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders.\n", numFinders)
	finders := make([]<-chan interface{}, numFinders)
	fmt.Println("Primes:")
	for i := 0; i < numFinders; i++ {
		finders[i] = pipeline.PrimeFinder(done, randIntStream)
	}
	for prime := range pipeline.Take(done, fanin.FanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}
	fmt.Printf("Search took: %v\n", time.Since(start))
}
