package main

import (
	"fmt"
	"time"
)

// This example show how to create a parallel program with semaphore(rate limit) by goroutine
//
func main() {
	timeList := []int{2, 4, 6, 8, 10, 12, 14, 16, 18}
	// Parallel goroutine limit
	limit := 3
	// Semaphore channel
	sem := make(chan bool, limit)
	for _, si := range timeList {
		// Push job sign into rate limit channel
		sem <- true

		go func(sec int) {
			// Pop up job sign when job is done
			defer func() { <-sem }()
			fmt.Println("This Goroutine will last", sec, "seconds")
			time.Sleep(time.Duration(sec) * time.Second)
		}(si)
	}
	// Block main goroutine until all of sub goroutine is done
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}

}
