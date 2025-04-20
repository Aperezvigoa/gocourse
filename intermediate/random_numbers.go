package intermediate

import (
	"fmt"
	"math/rand"
)

func main() {

	// fmt.Println(rand.Intn(6) + 5) // min 5 max 10

	/*
		val := rand.New(rand.NewSource(time.Now().Unix()))
		fmt.Println(val.Intn(101))
	*/

	// fmt.Println(rand.Float64()) // 0.0 - 1.0

	for {
		// Show the menu
		fmt.Println("\nWelcome to the Dice Roll Game")
		fmt.Println("1. Roll the dice")
		fmt.Println("2. Exit")
		fmt.Print("Enter your choice (1 or 2): ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid choice, please enter 1 or 2")
			continue
		}

		switch choice {
		case 1:
			dice1 := rand.Intn(6) + 1
			dice2 := rand.Intn(6) + 1

			// Show results:
			fmt.Printf("You rolled a %d and a %d\n", dice1, dice2)
			fmt.Println("Result of roll:", dice1+dice2)
		case 2:
			fmt.Println("Thanks for playing.")
			return
		default:
			fmt.Println("Invalid choice, please enter 1 or 2")
		}
	}
}
