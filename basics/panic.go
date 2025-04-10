package basics

import (
	"fmt"
)

func main() {

	// panic (interface{})

	// Valid input
	process(10)
	// Invalid input
	// process(-5)

	anotherProcess(-15)
}

func process(input int) {
	if input < 0 {
		panic("Invalid input.")
	} else {
		fmt.Println(input, "* 5:", input*5)
	}
}

func anotherProcess(i int) {
	defer fmt.Println("Deferred Message 1")
	defer fmt.Println("Deferred Message 2")

	if i < 0 {
		fmt.Println("Before panic")
		panic("Im in panic.")
		fmt.Println("After panic")
		defer fmt.Println("Deferred Message 3")
	}
	defer fmt.Println("Deferred Message 4")
}
