package main

import (
	"context"
	"fmt"
	"time"
)

func SimulateProcessing(ctx context.Context) error {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Finished processing for request ID:", ctx.Value("requestID"))
		return nil
	case <-ctx.Done():
		fmt.Println("Request ID:", ctx.Value("requestID"), "- Error:", ctx.Err())
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	requestID := "12345"
	ctx = context.WithValue(ctx, "requestID", requestID)

	SimulateProcessing(ctx)
}
