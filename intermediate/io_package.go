package intermediate

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		log.Fatalln("Error reading from reader.", err)
	}

	fmt.Println(string(buf[:n]))
}

func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error writing to writer.", err)
	}
}

func closeResource(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatalln("Error closing.", err)
	}
}

func bufferExample() {
	var buf bytes.Buffer
	buf.WriteString("Hello Buffer!")
	fmt.Println(buf.String())
}

func multiReaderExample() {
	r1 := strings.NewReader("Hello ")
	r2 := strings.NewReader("world")
	mr := io.MultiReader(r1, r2)
	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(mr)
	if err != nil {
		log.Fatalln("Error reading...", err)
	}
	fmt.Println(buffer.String())
}

func pipeExample() {
	pr, pw := io.Pipe()
	go func() {
		pw.Write([]byte("Hello Pipe"))
		pw.Close()
	}()

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(pr)
	fmt.Println(buffer.String())
}

func writeToFile(filepath string, data string) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalln("Error opening-creating file:", err)
	}
	defer closeResource(file)

	/*
		_, err = file.Write([]byte(data))
		if err != nil {
			log.Fatalln("Error writing data...")
		}
	*/
	writer := io.Writer(file)
	_, err = writer.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error writing data..")
	}
}

func main() {
	fmt.Println("=== Read From Reader ===")
	readFromReader(strings.NewReader("Hello Golang"))

	fmt.Println("=== Write To Writer ===")
	var writer bytes.Buffer
	writeToWriter(&writer, "I love Go")
	fmt.Println(writer.String())

	fmt.Println("=== Buffer Example ===")
	bufferExample()

	fmt.Println("=== Multi Reader Example ===")
	multiReaderExample()

	fmt.Println("=== Pipe Example ===")
	pipeExample()

	fmt.Println("=== Write To File ===")
	filepath := "io.txt"
	writeToFile(filepath, "In love with Go & Linux")

	resource := MyResource{"Gopher"}
	closeResource(resource) // This works because we've implemented the io.Close method
	// With same signature :P
}

type MyResource struct {
	name string
}

func (mr MyResource) Close() error {
	fmt.Println("Closing resource...", mr.name)
	return nil
}
