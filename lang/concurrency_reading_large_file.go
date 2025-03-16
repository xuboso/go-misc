package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

func main() {

	const mb = 1024 * 1024
	const gb = 1024 * mb

	// A waitGroup to wait for all go-routines to finish.
	wg := sync.WaitGroup{}

	// This channel is used to send very read word in various go-routines.
	channel := make(chan string)

	// A dictionary which stores the count of unique words.
	dict := make(map[string]int64)

	// Done is a channel to signal the main thread that all the words
	// has been entered in the dictionary.
	done := make(chan bool, 1)

	// Read all incoming words from the channel and add them to the dictionary.
	go func() {
		for s := range channel {
			dict[s]++
		}

		// Signal the main thread that all the words have entered the dictionary.
		done <- true
	}()

	// Current signifies the counter for bytes of the file.
	var current int64

	// Limit signifies the chunk size of the to be processed.
	var limit int64 = 500 * mb

	for i := 0; i < 2; i++ {
		wg.Add(2)

		go func() {
			read(current, limit, "gameofthrones.txt", channel)
			fmt.Printf("%d thread has been completed \n", i)
			wg.Done()
		}()

		// Increment the current by 1+(last byte of the chunk).
		current += limit + 1
	}

	// Wait for all goroutines to complete.
	wg.Wait()
	close(channel)

	// Wait for dictionary to process all the words.
	<-done
	close(done)
}

func read(offset int64, limit int64, fileName string, channel chan string) {
	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	// Move the pointer of the file to the start of designated chunk.
	file.Seek(offset, 0)
	reader := bufio.NewReader(file)

	// This block of code ensures that the start of chunk is a new word.
	// If a character is encountered at the given position it moves a few bytes
	// till the end of the word.
	if offset != 0 {
		_, err = reader.ReadBytes(' ')
		if err == io.EOF {
			fmt.Println("EOF")
			return
		}

		if err != nil {
			panic(err)
		}
	}

	var cummulativeSize int64
	for {
		// Break if read size has exceed the chunk size.
		if cummulativeSize > limit {
			break
		}

		b, err := reader.ReadBytes(' ')

		// Break if end of file is encountered.
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		cummulativeSize += int64(len(b))
		s := strings.TrimSpace(string(b))
		if s != "" {
			// Send the read word in the channel to enter into dictionary.
			channel <- s
		}
	}
}
