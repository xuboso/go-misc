/**
 * https://www.bytesizego.com/blog/single-flight
 */
package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

var (
	group singleflight.Group
	cache = make(map[string]string)
	mu    sync.Mutex
)

func fetchData(key string) (string, error) {
	// Simulate a slow operation
	time.Sleep(2 * time.Second)
	return "data for " + key, nil
}

func getData(key string) (string, error) {
	mu.Lock()
	if data, found := cache[key]; found {
		mu.Unlock()
		return data, nil
	}
	mu.Unlock()

	v, err, _ := group.Do(key, func() (interface{}, error) {
		data, err := fetchData(key)
		if err != nil {
			return nil, err
		}
		mu.Lock()
		cache[key] = data
		mu.Unlock()
		return data, nil
	})
	if err != nil {
		return "", err
	}
	return v.(string), nil
}

func main() {
	keys := []string{"key1", "key1", "key2", "key1", "key2"}
	var wg sync.WaitGroup
	for _, key := range keys {
		wg.Add(1)
		go func(k string) {
			defer wg.Done()
			data, err := getData(k)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Fetched:", data)
		}(key)
	}
	wg.Wait()
}
