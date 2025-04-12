package intermediate

import "fmt"

func main() {
	var firstDigits = []int{0, 1}
	fmt.Println(getFibonacci(10, firstDigits))
}

func getFibonacci(digits int, numbers []int) []int {
	if digits == 0 {
		return numbers
	}
	numbers = append(numbers, numbers[len(numbers)-1]+numbers[len(numbers)-2])
	return getFibonacci(digits-1, numbers)
}
