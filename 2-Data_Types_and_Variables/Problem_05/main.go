package main

import "fmt"

func main() {
	var num1 float32 = 3.14
	fmt.Print("float32 with value 3.14 -> int = ")
	var number1 int = int(num1)
	fmt.Println(number1)
	var num2 float64 = 12.3456789
	fmt.Print("float64 with value 12.3456789 -> float32 = ")
	var number2 float32 = float32(num2)
	fmt.Println(number2)
	var num3 int16 = 1234
	fmt.Print("int16 with value 1234 -> int8 = ")
	var number3 int8 = int8(num3)
	fmt.Println(number3)
	var num4 int16 = 1234
	fmt.Print("int16 with value 1234 -> uint8 = ")
	var number4 uint8 = uint8(num4)
	fmt.Println(number4)

}
