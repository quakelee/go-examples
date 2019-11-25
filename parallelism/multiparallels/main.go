package main

import (
	"fmt"
	"time"
)

// This example show how to use a multiple channels to control parallel program
//
func main() {
	limit := 8
	aList := []int{6, 8, 10, 12, 18, 12, 10, 8}
	bList := []int{5, 7, 9, 11, 13, 11, 9, 7}
	cList := []int{8, 10, 12, 14, 16, 14, 12, 10}
	// 3 channels for 3 groups of processes
	aChan := make(chan bool, limit)
	bChan := make(chan bool, limit)
	cChan := make(chan bool, limit)

	aCount := 0
	bCount := 0
	cCount := 0

	// Start multiple processes
	for i := 0; i < len(aList); i++ {
		go func(sec int) {
			fmt.Println("This Group A Goroutine will last", sec, "seconds")
			time.Sleep(time.Duration(sec) * time.Second)
			aChan <- true
		}(aList[i])
		go func(sec int) {
			fmt.Println("This Group B Goroutine will last", sec, "seconds")
			time.Sleep(time.Duration(sec) * time.Second)
			bChan <- true
		}(bList[i])
		go func(sec int) {
			fmt.Println("This Group C Goroutine will last", sec, "seconds")
			time.Sleep(time.Duration(sec) * time.Second)
			cChan <- true
		}(cList[i])
	}

	// Close Channel C after 16 seconds
	go func() {
		time.Sleep(time.Second * time.Duration(16))
		close(cChan)
	}()
	// Pop up multiple processes results
	for c := 0; c < 24; c++ {
		select {
		case <-aChan:
			aCount = aCount + 1
			fmt.Printf("This is A group No. %d goroutine\n", aCount)
		case <-bChan:
			bCount = bCount + 1
			fmt.Printf("This is B group No. %d goroutine\n", bCount)
		// Check channel is active
		case _, ok := <-cChan:
			if ok {
				cCount = cCount + 1
				fmt.Printf("This is C group No. %d goroutine\n", cCount)
			} else {
				fmt.Println("Group C is over")
			}
		}
	}
}
