package intermediate

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("line_filters_sample_file.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	defer func() {
		err = file.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println("Closing file:", file.Name())
	}()

	scanner := bufio.NewScanner(file)

	lineNumber := 1

	// Keyword to filter lines
	keyword := "important"

	// Read and filter lines
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			updatedLine := strings.ReplaceAll(line, keyword, "necessary")
			fmt.Printf("%d. Filtered line: %s\n", lineNumber, line)
			fmt.Printf("%d. Updated line: %s\n", lineNumber, updatedLine)
			lineNumber++
		}
	}
	if scanner.Err() != nil {
		fmt.Println("Error scanning file:", scanner.Err())
	}
}
