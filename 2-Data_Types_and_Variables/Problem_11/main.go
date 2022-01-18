package main

import (
	"fmt"
)

func main() {
	fmt.Println("Given 1234")
	var a int = 1234
	var b int = a % 10
	a = a / 10
	var c int = a % 10
	a = a / 10
	var d int = a % 10
	a = a / 10
	fmt.Println(b + c + d + a)

}
