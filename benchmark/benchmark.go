package benchmark

import (
	"fmt"
	"time"
)

// Timer holds the start time for benchmarking
type Timer struct {
	start time.Time
}

// Start initializes and returns a Timer instance
func Start() *Timer {
	return &Timer{start: time.Now()}
}

// PrintElapsed prints the elapsed time since the Timer was started
func (t *Timer) PrintElapsed() {
	elapsed := time.Since(t.start)
	fmt.Printf("Elapsed time: %v\n", elapsed)
}
