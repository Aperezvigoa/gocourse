package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Substraction:", result)

	result = a * b
	fmt.Println("Multiplication:", result)

	result = a / b
	fmt.Println("Division:", result)

	result = a % b
	fmt.Println("Remainder:", result)

	const PI float64 = 22.0 / 7
	fmt.Println(PI)

	// Overflow signed integers

	maxInt := 9223372036854775807 // int64 max value
	fmt.Println("Max Int:", maxInt)
	maxInt += 16
	fmt.Println("Overflow:", maxInt)

	// Overflow unsigned integers

	var uMaxInt uint64 = 18446744073709551615 // uint64 max value
	fmt.Println("uMaxInt:", uMaxInt)
	uMaxInt += 16
	fmt.Println("U-Overflow:", uMaxInt)

	// Underflow small float

	var smallFloat float64 = 1.0e-323
	fmt.Println("Small float:", smallFloat)

	smallFloat /= math.MaxFloat64
	fmt.Println("Underflow:", smallFloat)
}
