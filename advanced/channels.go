package advanced

import (
	"fmt"
)

// We use channels to enable safe and efficient communication between concurrent go routines.
// We need to recieve values into a channel inside a go routine
// Syntaxis: variable := make(chan type)
func main() {
	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		greeting <- greetString // Blocking because it is continuosly trying to receive values, it is ready to recieve continuous flow of data.
		greeting <- "World"

		for _, e := range "abcde" {
			greeting <- "Alphabet:" + string(e)
		}
	}()
	// go func() {
	// 	receiver := <-greeting
	// 	fmt.Println(receiver)
	//	receiver = <-greeting
	//	fmt.Println(receiver)
	// }()

	receiver := <-greeting
	fmt.Println(receiver)
	// receiver = <-greeting
	// fmt.Println(receiver)
	fmt.Println(<-greeting)

	for range 5 {
		rcvr := <-greeting
		fmt.Println(rcvr)
	}

	// time.Sleep(1 * time.Second)
	fmt.Println("End of program")
}
