package intermediate

import "fmt"

func main() {

	fmt.Println(factorial(5))
	fmt.Println(factorial(10))

	fmt.Println(sumOfDigits(19))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func sumOfDigits(x int) int {
	if x < 10 {
		return x
	}
	return x%10 + sumOfDigits(x/10)
}
