package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	result := math.Pow(2, 11)
	fmt.Println(result)
	time.Sleep(5 * time.Second)
}
