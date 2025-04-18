package intermediate

import (
	"fmt"
	"regexp"
)

func main() {

	fmt.Println("He said, \"I am great\"")

	// --- Compile a regex patter to match email address

	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9._-]+\.[a-zA-Z]{2,}`)

	// Test strings
	email1 := "alex@github.com"
	email2 := "invalidEmail@"

	// Match
	fmt.Println("Email1:", re.MatchString(email1))
	fmt.Println("Email2:", re.MatchString(email2))

	// Capturing Groups
	// Compile a regex pattern to capture date components

	re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

	// Test string
	date1 := "2024-07-30"

	submatches := re.FindStringSubmatch(date1)
	fmt.Println(submatches[0])
	fmt.Println(submatches[1])
	fmt.Println(submatches[2])
	fmt.Println(submatches[3])

	// Source string

	str := "Hello world"

	re = regexp.MustCompile(`[aeiou]`)

	result := re.ReplaceAllString(str, "*")
	fmt.Println(result)

	// i - case insensitive
	// m - multi line model
	// s - dot matches all

	re = regexp.MustCompile(`(?i)go`) // Match go (case insensitive) go - Go - GO - gO

	// Test string
	anotherText := "I love gOlang"
	fmt.Println(re.MatchString(anotherText))
}
