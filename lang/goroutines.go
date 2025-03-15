package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {
	fmt.Println("Semaphore example starting!")

	var (
		ctx        = context.TODO()
		maxWorkers = 5                                        // max goroutines running at one time
		sem        = semaphore.NewWeighted(int64(maxWorkers)) // semaphore
		tasks      = make([]int, 100)
	)

	// loop through slice of tasks to be completed, by index
	for i := range tasks {
		// Acquire will block if there are already maxWorkers goroutines running,
		// until one is released.
		if err := sem.Acquire(ctx, 1); err != nil {
			fmt.Printf("Failed to acquire semaphore: %v", err)
			break
		}

		go func(i int) {
			// be sure to release the worker when the work is done!
			defer sem.Release(1)
			doLongTask(i)
		}(i)
	}

	if err := sem.Acquire(ctx, int64(maxWorkers)); err != nil {
		fmt.Printf("Failed to acquire semaphore: %v", err)
	}

	fmt.Println("Completed all tasks!")
}

func doLongTask(i int) {
	fmt.Println("doing long work")
	// wait some time, simulate IO block or something
	time.Sleep(3 * time.Second)
	fmt.Printf("completed long task %d\n", i)
}
