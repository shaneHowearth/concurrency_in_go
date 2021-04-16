// Package main -
package main

import (
	"log"
	"os"
	"time"

	"github.com/shanehowearth/concurrency_in_go/steward"
)

func main() {
	log.SetOutput(os.Stdout)

	log.SetFlags(log.Ltime | log.LUTC)
	doWork := func(done <-chan interface{}, _ time.Duration) <-chan interface{} {
		log.Println("ward: Hello, I'm irresponsible!")
		go func() {
			log.Println("ward: I am halting.")
		}()
		return nil
	}
	doWorkWithSteward := steward.NewSteward(4*time.Second, doWork)
	done := make(chan interface{})
	time.AfterFunc(9*time.Second, func() {
		log.Println("main: halting steward and ward.")
		close(done)
	})
	for range doWorkWithSteward(done, 4*time.Second) {
	}
	log.Println("Done")
}
