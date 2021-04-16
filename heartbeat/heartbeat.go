// Package heartbeat -
package heartbeat

import (
	"math/rand"
	"time"
)

// DoWork -
func DoWork(done <-chan interface{}) (<-chan interface{}, <-chan int) {
	heartbeatStream := make(chan interface{}, 1)
	workStream := make(chan int)
	go func() {
		defer close(heartbeatStream)
		defer close(workStream)
		for i := 0; i < 10; i++ {
			select {
			case heartbeatStream <- struct{}{}:
			default:
			}
			select {
			case <-done:
				return
			case workStream <- rand.Intn(10):
			}
		}
	}()
	return heartbeatStream, workStream
}

// DoWorkDelay -
func DoWorkDelay(
	done <-chan interface{},
	nums ...int,

) (<-chan interface{}, <-chan int) {
	heartbeat := make(chan interface{}, 1)
	intStream := make(chan int)
	go func() {
		defer close(heartbeat)
		defer close(intStream)
		time.Sleep(2 * time.Second)
		for _, n := range nums {
			select {
			case heartbeat <- struct{}{}:
			default:

			}
			select {
			case <-done:
				return
			case intStream <- n:

			}

		}

	}()
	return heartbeat, intStream

}
