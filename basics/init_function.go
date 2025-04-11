package basics

import "fmt"

func init() {
	fmt.Println("Initializing package...")
}

func init() {
	fmt.Println("Second initialization...")
}

func init() {
	fmt.Println("Third initialization...")
}

func main() {

	fmt.Println("Inside the main function.")

}

// Output
/*
	Initializing package...
	Second initialization...
	Third initialization...
	Inside the main function.
*/
