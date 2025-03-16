package main

import (
	"fmt"
	"reflect"
)

func main() {
	var z reflect.Value
	fmt.Println(z)
	v := reflect.ValueOf((*int)(nil)).Elem()
	fmt.Println(v)
	fmt.Println(v == z)
	var i = reflect.ValueOf([]interface{}{nil}).Index(0)
	fmt.Println(i)
	fmt.Println(i.Elem() == z)
	fmt.Println(i.Elem())
}
