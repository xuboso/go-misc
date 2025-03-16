package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	readFileWithScanner()
	fmt.Println("====")
	readFileWithBuf()
}

func readFileWithScanner() {
	file, err := os.Open("golangbot/test.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Loop through the file and read each lilne
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	// Check for errors during the scan
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}
}

func readFileWithBuf() {
	file, err := os.Open("golangbot/test.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	// Create a new reader
	reader := bufio.NewReader(file)

	for {
		// Read until we encounter a newline character
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				// 如果最后一行没有换行符，line 中仍然包含最后一行的内容
				if len(line) > 0 {
					fmt.Print(line)
				}
				break
			}
			log.Fatalf("error reading file: %s", err)
		}

		fmt.Print(line)
	}
}
