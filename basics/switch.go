package basics

import (
	"fmt"
)

func main() {

	numberDay := 4

	switch numberDay {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
		// fallthrough
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("That's not a week day")
	}

	day := "Monday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday, go to job!")
	case "Saturday", "Sunday":
		fmt.Println("It's weekend, enjoy it!")
	default:
		fmt.Println("Uリꖌリ𝙹∴リ ↸ᔑ||")
	}

	// Switch with expressions

	number := 15

	switch {
	case number < 10:
		fmt.Println("Number is less than 10")
	case number >= 10 && number < 20:
		fmt.Println("Number is between 10 and 29")
	default:
		fmt.Println("The number is 20 or larger.")
	}

	secondNumber := 2

	switch {
	case secondNumber > 1:
		fmt.Println("secondNumber is greater than 1")
		fallthrough
	case secondNumber == 2:
		fmt.Println("secondNumber is 2.")
	default:
		fmt.Println("T∴𝙹 !¡ꖎ⚍ᓭ ᓭᒷ⍊ᒷリ ╎ᓭ ꖎᒷᓭᓭ ℸ ̣ ⍑ᔑリ ℸ ̣ ᒷリ.")
	}

	checkType(2.54)
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It's an integer")
	case float64:
		fmt.Println("It's a float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("Unknown type")
	}
}
