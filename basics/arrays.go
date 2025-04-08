package basics

import "fmt"

func main() {

	var numbers [5]int
	fmt.Println(numbers) // [0 0 0 0 0]

	numbers[0] = 20
	fmt.Println(numbers) // [20 0 0 0 0]

	// numbers[8] = 16  Out of bounds

	catNames := [3]string{"Cheddar", "Whiskers", "Lily"}
	fmt.Println(catNames)

	integerArray := [3]int{1, 2, 3}
	copyIntegerArray := integerArray

	copyIntegerArray[0] = 4
	fmt.Println("Integer Array:", integerArray)
	fmt.Println("Copy Integer Array:", copyIntegerArray)

	for i := 0; i < len(catNames); i++ {
		fmt.Println("I love my cat", catNames[i])
	}

	// short way to iterate an array
	for index, value := range catNames {
		fmt.Printf("index: %d - value: %v\n", index, value)
	}

	// Short way without index (using _ underscore || blank identifier)
	for _, v := range catNames {
		fmt.Printf("Short way :) -> %v\n", v)
	}

	arrayOne := [3]int{1, 2, 3}
	arrayTwo := [3]int{1, 2, 3}

	fmt.Println("Is arrayOne equals to arrayTwo:", arrayOne == arrayTwo)

	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix)

	originalArray := [3]int{1, 2, 3}
	var copyOgArray *[3]int

	copyOgArray = &originalArray
	fmt.Println("OriginalArray", originalArray)
	fmt.Println("CopyOGArray", copyOgArray)

	copyOgArray[0] = 5

	fmt.Println("OriginalArray", originalArray)
	fmt.Println("CopyOGArray", *copyOgArray)
}
