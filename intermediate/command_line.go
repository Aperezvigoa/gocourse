package intermediate

import (
	"flag"
	"fmt"
)

// The flag package allows defining flags with various types
// like int, string, bool etc... And automatically parses
// command line arguments into these flags.

func main() {
	// Define flags
	var name string
	var age int
	var male bool

	// Creating Flags
	flag.StringVar(&name, "name", "Linux", "Name of the user")
	flag.IntVar(&age, "age", 18, "Age of the user")
	flag.BoolVar(&male, "male", false, "Is the user male?")

	// Parsing Flags
	flag.Parse()

	// How to use? go run command_line.go -name Alex -age 23 -male true
	// We can use go run command_line.go --help to see flags

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Male:", male)
}
