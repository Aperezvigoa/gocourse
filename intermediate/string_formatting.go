package intermediate

import "fmt"

func main() {

	num := 42
	fmt.Printf("Num: %05d\n", num) // Num: 00042

	message := "Hello"
	fmt.Printf("|%10s|\n", message)  // |     Hello|
	fmt.Printf("|%-10s|\n", message) // |Hello     |

	message1 := `
	Hello
	World
	This is a literal
	`
	fmt.Println(message1)

	sqlQuery := `SELECT * FROM useres WHERE age > 30`
}
