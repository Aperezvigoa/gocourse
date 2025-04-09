package basics

import (
	"fmt"
	"slices"
)

func main() {
	// var numbers []int
	// var numbers1 = []int{1, 2, 3}

	// numbers2 := []int{3, 2, 1}

	// slice := make([]int, 5)

	a := [5]int{5, 6, 7, 8, 9}
	sliceTwo := a[1:4]

	fmt.Println("SliceTwo:", sliceTwo)

	sliceTwo = append(sliceTwo, 1, 2, 3)
	fmt.Println("SliceTwo:", sliceTwo)

	sliceCopy := make([]int, len(sliceTwo))
	copy(sliceCopy, sliceTwo)

	fmt.Println("SliceCopy:", sliceCopy)

	var nilSlice []int

	_ = nilSlice

	for i, v := range sliceCopy {
		fmt.Printf("At index: %d - value: %d\n", i, v)
	}

	sliceCopy[2] = 22

	for _, v := range sliceCopy {
		fmt.Println("New SliceCopy values:", v)
	}

	// Useful function

	if slices.Equal(sliceTwo, sliceCopy) {
		fmt.Println("The slices are equal.")
	} else {
		fmt.Println("The slices are not equal, let me fix that...")
		copy(sliceCopy, sliceTwo)
	}

	fmt.Println("Let's check it again.")

	if slices.Equal(sliceCopy, sliceTwo) {
		fmt.Println("Both slices are equal.")
	}

	// Multi-dimensional slice

	twoDSlice := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoDSlice[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDSlice[i][j] = i + j
		}
	}

	fmt.Println("twoDSlice:", twoDSlice)

	sliceThree := sliceCopy[2:5]
	fmt.Println("sliceThree:", sliceThree)

	fmt.Println("\nTest")

	testSlice := []int{}
	for i := range 100 {
		testSlice = append(testSlice, i+1)
	}

	fmt.Println(testSlice)
}
