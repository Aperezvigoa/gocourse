package basics

import "fmt"

func main() {

	sum := add(1, 2)
	fmt.Println("Sum:", sum)

	// Anonymous function
	greet := func() {
		fmt.Println("Hello Anonymous Function")
	}

	greet()

	operation := add
	result := operation(5, 6)

	fmt.Println("Result:", result)

	// Passing a function as an argument
	result1 := applyOperation(3, 2, add)
	fmt.Println("3 + 2:", result1)

	// Returning and using a function
	multiplyrBy2 := createMultiplier(2)
	fmt.Println("2 * 5:", multiplyrBy2(5))
}

func add(a, b int) int {
	return a + b
}

// function that takes a function as an argument
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// function that returns a function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
