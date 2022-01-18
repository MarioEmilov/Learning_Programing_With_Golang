package main

import (
	"fmt"
)

func main() {
	fmt.Println("Before swapping")
	var a = 14
	fmt.Println(a)
	var b = 32
	fmt.Println(b)
	fmt.Println()

	fmt.Println("After swapping")
	var c = a
	a = b
	b = c
	fmt.Println(a)
	fmt.Println(b)

}
