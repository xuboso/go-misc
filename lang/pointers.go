package main

import "fmt"

type BigList struct {
	values [1000000]int
}

func update_by_value(d BigList) {
	d.values[0] = 50
}

func update_by_pointer(d *BigList) {
	d.values[0] = 50
}

func main() {
	d := BigList{}
	update_by_pointer(&d) // fast, no copy
	fmt.Println(d.values[0])
	update_by_value(d) // slow, copies everything
	fmt.Println(d.values[0])
}
