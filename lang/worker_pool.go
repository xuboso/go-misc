package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * https://medium.com/@ahamrouni/mastering-concurrency-in-go-from-goroutines-to-semaphores-123fdd150213
 */
func worker(id int, wg *sync.WaitGroup, semaphore chan struct{}) {
	defer wg.Done()

	// Acquire the semaphore
	semaphore <- struct{}{}

	fmt.Printf("Worker %d is processing\n", id)
	time.Sleep(1 * time.Second) // Simulate work
	fmt.Printf("Worker %d has finished\n", id)

	// Release the semaphore
	<-semaphore
}

func main() {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 3) // limit to 3 concurrent workers

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go worker(i, &wg, semaphore)
	}

	wg.Wait()
	fmt.Println("All workers have completed")
}
