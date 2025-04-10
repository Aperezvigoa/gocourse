package basics

import "fmt"

func main() {
	usingDefer() // First -> Fourth -> Third -> Second

	testingDefer(2)
}

func usingDefer() {
	fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("Third")
	fmt.Println("Fourth")
}

func testingDefer(i int) {
	defer fmt.Println("Deferred i:", i)
	i++
	fmt.Println("Value of i:", i)
	defer fmt.Println("Deferred incremented i:", i)
}
