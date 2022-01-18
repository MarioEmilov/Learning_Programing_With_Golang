package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Given 3 4")
	a, b := 3.0, 4.0
	c := math.Sqrt(a*a + b*b)
	fmt.Println(c)
	fmt.Println("Given 6 7")
	a2, b2 := 6.0, 7.0
	c2 := math.Sqrt(a2*a2 + b2*b2)
	fmt.Println(c2)
}
