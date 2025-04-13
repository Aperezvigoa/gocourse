package intermediate

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	message := "A blue gopher"
	secondMessage := "A blue \tgopher"
	thirdMessage := "A blue \rgopher"
	rawMessage := `A\nblue\ngopher`

	fmt.Println(message)       // A blue gopher
	fmt.Println(secondMessage) // A blue  gopher
	fmt.Println(thirdMessage)  // gopher
	fmt.Println(rawMessage)    // A\nblue\gopher

	fmt.Println("Length of message:", len(message))

	fmt.Println("First character in message:", message[0]) // Ascii
	greeting := "Hello "
	name := "Alice"
	greetingWithName := greeting + name

	fmt.Println(greetingWithName)

	str1 := "Apple"
	str2 := "apple"
	str3 := "App"
	str4 := "app"

	fmt.Println("Apple == apple:", str1 == str2)
	fmt.Println("Apple < apple:", str1 < str2)
	fmt.Println("Apple > App:", str1 > str3)
	fmt.Println("App > app:", str3 > str4)

	messageToIterate := "I love Golang!"
	for _, c := range messageToIterate {
		// fmt.Printf("Char at index %d: %c\n", i, c)
		fmt.Printf("%c", c)
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCount([]byte(messageToIterate)))

	animal := "Gopher"
	programmingLanguage := "Golang"
	bestPLanguage := animal + programmingLanguage
	fmt.Println(bestPLanguage)

	var ch rune = 'A'
	fmt.Println("Rune value:", ch)
	fmt.Printf("Char value: %c\n", ch)

	chString := string(ch)
	fmt.Println("Rune as string:", chString)

	fmt.Printf("Type of chString: %T\n", chString)

	const NIHONGO = "日本語" // Japanase text
	fmt.Println(NIHONGO)

	japaneseMessage := "こんにちは"

	for _, rune := range japaneseMessage {
		fmt.Printf("%c\n", rune)
	}

	r := '🤑'
	fmt.Println(r)
	fmt.Printf("%c\n", r)
}
