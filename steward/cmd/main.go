// Package main -
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/shanehowearth/concurrency_in_go/bridge"
	"github.com/shanehowearth/concurrency_in_go/pipeline"
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

	doWorkFn := func(
		done <-chan interface{},
		intList ...int,

	) (steward.StartGoroutineFn, <-chan interface{}) {
		intChanStream := make(chan (<-chan interface{}))
		intStream := bridge.Bridge(done, intChanStream)
		doWork = func(
			done <-chan interface{},
			pulseInterval time.Duration,

		) <-chan interface{} {
			intStream := make(chan interface{})
			heartbeat := make(chan interface{})
			go func() {
				defer close(intStream)
				select {
				case intChanStream <- intStream:
				case <-done:
					return

				}
				pulse := time.Tick(pulseInterval)
				for {
				valueLoop:
					for _, intVal := range intList {
						if intVal < 0 {
							log.Printf("negative value: %v\n", intVal)
							return

						}
						for {
							select {
							case <-pulse:
								select {
								case heartbeat <- struct{}{}:
								default:

								}
							case intStream <- intVal:
								continue valueLoop
							case <-done:
								return

							}

						}

					}

				}

			}()
			return heartbeat

		}
		return doWork, intStream

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

	log.SetFlags(log.Ltime | log.LUTC)
	log.SetOutput(os.Stdout)
	done = make(chan interface{})
	defer close(done)
	doWork, intStream := doWorkFn(done, 1, 2, -1, 3, 4, 5)
	doWorkWithSteward = steward.NewSteward(1*time.Millisecond, doWork)
	doWorkWithSteward(done, 1*time.Hour)
	for intVal := range pipeline.Take(done, intStream, 6) {
		fmt.Printf("Received: %v\n", intVal)

	}
}
