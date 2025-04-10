package basics

import (
	"errors"
	"fmt"
)

func main() {
	resultado, resto := divide(10, 3)
	fmt.Println("10 / 3:")
	fmt.Println("Resultado:", resultado)
	fmt.Println("Resto:", resto)

	fmt.Println("\nComparing 3 - 2")
	firstCompare, firstError := compare(3, 2)
	if firstError != nil {
		fmt.Println("Error:", firstError)
	} else {
		fmt.Println(firstCompare)
	}

	fmt.Println("\nComparing 5 - 9")
	secondCompare, secondError := compare(5, 9)
	if secondError != nil {
		fmt.Println("Error:", secondError)
	} else {
		fmt.Println(secondCompare)
	}

	fmt.Println("\nComparing 4 - 4")
	thirdCompare, thirdError := compare(4, 4)
	if thirdError != nil {
		fmt.Println("Error:", thirdError)
	} else {
		fmt.Println(thirdCompare)
	}

}

func divide(dividendo int, divisor int) (int, int) {
	resultado := dividendo / divisor
	resto := dividendo % divisor
	return resultado, resto
}

func compare(a int, b int) (string, error) {
	if a > b {
		return "A is greater than B", nil
	} else if b > a {
		return "B is greater than A", nil
	} else {
		return "", errors.New("a is equal than b")
	}
}
