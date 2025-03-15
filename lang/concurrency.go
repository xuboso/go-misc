package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu    sync.Mutex
	cond  = sync.NewCond(&mu)
	ready = false
)

func waiter(id int) {
	mu.Lock()
	for !ready {
		cond.Wait()
	}
	fmt.Printf("Goroutine %d proceeding\n", id)
	mu.Unlock()
}

func main() {
	for i := 1; i <= 3; i++ {
		go waiter(i)
	}

	mu.Lock()
	ready = true
	cond.Broadcast()
	mu.Unlock()
	time.Sleep(1 * time.Second)
}
