package basics

import "fmt"

func main() {

	// ... Ellipsis

	statement, total := sum("The result is:", 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(statement, total)

	statementTwo, totalTwo := sum("The result is:")
	fmt.Println(statementTwo, totalTwo)

	sliceOne := []int{1, 2, 3, 4, 5, 6}
	statementThree, sequenceThree := sum("The result of slice sum is:", sliceOne...)
	fmt.Println(statementThree, sequenceThree)
}

func sum(stringToReturn string, nums ...int) (string, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return stringToReturn, total
}
