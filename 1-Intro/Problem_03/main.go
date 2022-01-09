package main

import (
	"fmt"
	"time"
)

func main() {
	var name string = "Linus Torvalds."
	var age int = 28

	fmt.Println("My name is ", name)
	fmt.Println("My age is ", age)
	time.Sleep(5 * time.Second)
}
