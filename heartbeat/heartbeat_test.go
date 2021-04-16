package heartbeat

import "testing"

func TestDoWork_GeneratesAllNumbers(t *testing.T) {
	done := make(chan interface{})
	defer close(done)
	intSlice := []int{0, 1, 2, 3, 5}

	heartbeat, results := DoWorkDelay(done, intSlice...)
	<-heartbeat
	i := 0
	for r := range results {
		if expected := intSlice[i]; r != expected {
			t.Errorf("index %v: expected %v, but received %v,", i, expected, r)
		}
		i++
	}
}
