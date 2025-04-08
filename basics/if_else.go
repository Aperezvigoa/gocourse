package basics

import "fmt"

func main() {

	age := 23
	if age >= 18 {
		fmt.Println("You are an adult.")
	}

	temperature := 23.0
	if temperature >= 30 {
		fmt.Println("Its hot outside.")
	} else if temperature >= 16 {
		fmt.Println("Its a nice temperature.")
	} else {
		fmt.Println("Its cold outside.")
	}

	score := 83
	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 80 {
		fmt.Println("Grade B")
	} else if score >= 70 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade D")
	}

	testNumber := 30
	// Nested conditions
	if testNumber%2 == 0 {
		if testNumber%3 == 0 {
			fmt.Printf("The number %d is divisible by both 2 and 3\n", testNumber)
		} else {
			fmt.Printf("The number %d is only divisible by 2\n", testNumber)
		}
	} else {
		fmt.Printf("The number %d is not divisible by 2", testNumber)
	}
}
