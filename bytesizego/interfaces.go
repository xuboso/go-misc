package main

import "fmt"

type Greeter interface {
	Greet() string
}

type EnglishGreeter struct{}

func (e EnglishGreeter) Greet() string {
	return "Hello!"
}

type SpanishGreeter struct{}

func (s SpanishGreeter) Greet() string {
	return "Hola!"
}

func main() {
	var greeter Greeter

	greeter = EnglishGreeter{}
	fmt.Println(greeter.Greet())

	greeter = SpanishGreeter{}
	fmt.Println(greeter.Greet())
}
