package main

import (
	"fmt"
	"runtime"
)

func main() {
	num := runtime.NumCPU()
	fmt.Println("Number of CPUs:", num)
}
