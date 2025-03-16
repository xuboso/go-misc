package main

import "fmt"

func main() {
	var num int = 10
	var ptr *int = &num // Pointer stores the address of num

	fmt.Println("Before modification, num:", num) // Before modification, num: 10

	*ptr = 100                                   // Modifying the value at the memory address
	fmt.Println("After modification, num:", num) // After modification, num: 100
}
