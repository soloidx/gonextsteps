package main

import (
	"fmt"
)

const state = "Lima"

var name string

func main() {

	name = "Ider Delzo"
	from := `Huancayo`
	var n int

	fmt.Printf("Hello, fellow %s Gophers!\n", state)
	fmt.Printf("My name is %s and I'm from %s.\n", name, from)
	fmt.Printf("By the time %d o'clock comes around, we'll know how to code in Go!\\n", n)
	fmt.Println("Let's get started!")
}
