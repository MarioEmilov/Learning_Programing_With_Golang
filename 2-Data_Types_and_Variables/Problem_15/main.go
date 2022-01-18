package main

import (
	"fmt"
)

func main() {
	fmt.Println("Given 1")
	lbs := 2.2046226218
	kg1 := 1
	fmt.Println(float64(kg1) * lbs)

	fmt.Println("Given 78.5")
	kg2 := 78.5
	fmt.Println(float64(kg2) * lbs)

}
