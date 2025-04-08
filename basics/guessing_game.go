package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Generate a random number between 1 and 100
	target := random.Intn(100) + 1

	fmt.Println("Welcome to the Simplest Guessing Game")
	fmt.Println("-------------------------------------")
	fmt.Println("I have choosen a number between 1 and 100")
	fmt.Println("Can you guess it?")

	var guess int

	for {
		fmt.Println("Enter your guess:")
		fmt.Scanln(&guess)

		if guess == target {
			println("Congrats! You guessed it")
			break
		} else if guess > target {
			fmt.Println("Too high! Try again!")
		} else {
			fmt.Println("Too low! Try again!")
		}
	}

}
