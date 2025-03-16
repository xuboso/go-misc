package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	emp := Employee{
		firstName: "john",
		lastName:  "Doe",
		age:       30,
	}

	fmt.Println("Employee:", emp)
	fmt.Println("Employee Age:", emp.age)
}
