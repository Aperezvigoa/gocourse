package basics

import (
	"fmt"
	"os"
)

func main() {

	defer fmt.Println("Program closed.")

	fmt.Println("Starting main function.")
	os.Exit(1)
	fmt.Println("End of main function.")

	// Output
	/*
		Starting main function.
		exit status 1
	*/
}
