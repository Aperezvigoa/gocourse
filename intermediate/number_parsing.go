package intermediate

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "12345"

	// --- strconv.Atoi

	parsedNumber, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
		return
	}
	fmt.Println("ParsedNumber:", parsedNumber)

	// --- strconv.ParseInt

	parsedNumber1, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
	}
	fmt.Println("ParsedNumber1:", parsedNumber1)

	// --- Parsing Float

	floatStr := "3.14"
	parsedFloat, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
	}
	fmt.Println("ParsedFloat:", parsedFloat)

	// --- Binary String

	binaryStr := "10101010"
	parsedBinToInt, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
	}
	fmt.Println("parsedBinToInt:", parsedBinToInt)

	// --- Hex String

	hexStr := "32FF"
	hexParsed, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
	}
	fmt.Println("hexParsed:", hexParsed)
}
