package intermediate

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// --- Opening File
	file, err := os.Open("reading_file.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer func() {
		fmt.Println("Closing file:", file.Name())
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}()

	fmt.Println("File opened successfully.")

	/* Reading all the file into data and then printing out into console
	data := make([]byte, 1024) // Buffer to read data into
	bytesReaded, err := file.Read(data)

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Bytes readed:", bytesReaded)
	fmt.Printf("File content: %s\n", string(data))
	*/

	// --- Reading Line by Line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line:", line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

}
