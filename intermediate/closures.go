package intermediate

import "fmt"

func main() {

	fmt.Println("adder out main:")

	sequence := adder()
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	fmt.Println("\nSecond sequence\n")
	secondSequence := adder()
	fmt.Println(secondSequence())
	fmt.Println(secondSequence())

	fmt.Println("\nsubtraction:")

	substraction := func() func(int) int {

		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	fmt.Println(substraction(1))
	fmt.Println(substraction(1))
	fmt.Println(substraction(1))
	fmt.Println(substraction(1))
}

func adder() func() int {
	i := 0
	fmt.Println("Previous value of i:", i)
	return func() int {
		i++
		fmt.Println("Actual value of i:", i)
		return i
	}
}
