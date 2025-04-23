package intermediate

import (
	"fmt"
	"os"
)

func main() {

	// --- Creating a file
	file, err := os.Create("writing_files_text.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// --- Writing data to the file
	data := []byte("Hello Go!\n")
	_, err = file.Write(data)

	if err != nil {
		fmt.Println("Error writing the file:", err)
		return
	}

	fmt.Println("Data has been written to the file successfully")

	file, err = os.Create("writeString.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
	}
	defer file.Close()

	text := "I love programming in Golang, I'm thinking in purchase a Gopher puff"

	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Error writing the file:", err)
	}

	fmt.Println("String has been written to the file successfully")
}
