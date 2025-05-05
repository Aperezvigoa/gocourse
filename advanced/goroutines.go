package advanced

import (
	"fmt"
	"log"
	"time"
)

// Goroutines are just functions that leave the main thread and run in the background
// and come back to join the mean thread once the functions are finished/ready to return
// any values.
// Goroutines do not stop the program flow and are non blocking.

// A goroutine starts when it's created and runs concurrently with other goroutines.
// A goroutine contains a function and when the function is finished, it exit

func main() {

	var err error

	fmt.Println("Beginning of program.")
	go sayHello()
	fmt.Println("After sayHello function.")

	go func() {
		err = doWork()
	}()
	// err = go doWork() is not accepted

	go printNumbers()
	go printLetters()

	time.Sleep(2 * time.Second)

	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Work completed successfully")
	}
}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from Goroutine")
}

func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println("Number:", i, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "abcde" {
		fmt.Println(string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	// Simulate work
	time.Sleep(1 * time.Second)
	return fmt.Errorf("an error ocurred in doWork")
}
