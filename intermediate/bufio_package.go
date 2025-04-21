package intermediate

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(strings.NewReader("Hello, bufio package! It's cool\n"))

	// Reading byte slice
	data := make([]byte, 20)
	n, err := reader.Read(data) // Transfer the 'data' readed to data

	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	fmt.Printf("Data readed: %v bytes: %s\n", n, data[:n])

	readedData, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading string:", err)
		return
	}

	fmt.Println("Read string:", readedData)

	fmt.Println("---------------------")
	fmt.Println("B U F I O W R I T E R")
	fmt.Println("---------------------\n")

	writer := bufio.NewWriter(os.Stdout)

	// Writing byte slice
	wData := []byte("Hello, bufio package!\n")
	wN, err := writer.Write(wData) // Writed in buffer, not in "os.Stdout"
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", wN)

	// Flush the buffer to ensure all data is written to os.Stdout
	err = writer.Flush()

	if err != nil {
		fmt.Println("Error flushing data:", err)
		return
	}

	// Writing string
	str := "Simple string\n"

	b, err := writer.WriteString(str)
	if err != nil {
		fmt.Println("Error writing string:", err)
		return
	}

	fmt.Printf("Writed %d bytes in buffer\n", b)

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing string:", err)
		return
	}
}
