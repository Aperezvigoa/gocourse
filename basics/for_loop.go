package basics

import "fmt"

func main() {
	// Simple iteration over a range
	for i := 1; i <= 5; i++ {
		fmt.Println("Hi", i)
	}

	// Iteration over collection
	numbers := []int{1, 2, 3, 4, 5, 6}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("Value is %d\n", i)
	}

	rows := 5
	// Outer loop
	for i := 1; i <= rows; i++ {
		// Inner loop
		for x := 1; x <= rows-i; x++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// New features
	for i := range 10 {
		fmt.Println(10 - i)
	}
}
