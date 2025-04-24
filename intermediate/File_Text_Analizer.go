package gocourse

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"text/template"
)

type Analizator struct {
	OriginalText string
	Results      map[string]interface{}
	Errors       []error
}

// Analizer Functions
func (a *Analizator) CalculateHash() {
	hash256 := sha256.Sum256([]byte(a.OriginalText))
	hash512 := sha512.Sum512([]byte(a.OriginalText))

	a.Results["Hash256"] = fmt.Sprintf("%x", hash256)
	a.Results["Hash512"] = fmt.Sprintf("%x", hash512)
}

func (a *Analizator) Statistics() {
	words := strings.Fields(a.OriginalText)
	charCount, maxLenWord := countChars(words)

	a.Results["Total_words"] = len(words)
	a.Results["Character_count"] = charCount
	a.Results["Max_length_word"] = maxLenWord
}

func (a *Analizator) Encode() {
	encodedText := base64.StdEncoding.EncodeToString([]byte(a.OriginalText))
	a.Results["Encoded"] = encodedText
}

// Text template
var temp = template.Must(template.New("result").Parse(
	`--------------------
Text Analysis: {{.OriginalText}}
--------------------
Hashes:
 - SHA256: {{index .Results "Hash256"}}
 - SHA512: {{index .Results "Hash512"}}
Statistics:
 - Words count: {{index .Results "Total_words"}}
 - Char count: {{index .Results "Character_count"}}
 - Max Length Word: {{index .Results "Max_length_word"}}
Encoded Text:
 - Base64: {{index .Results "Encoded"}}`,
))

func (a *Analizator) PrintAnalysis() error {
	err := temp.Execute(os.Stdout, a)
	if err != nil {
		fmt.Println(err)
		return CustomError{ErrorMessage: "Can't execute analysis!"}
	}
	return nil
}

// Custom Errors
type CustomError struct {
	ErrorMessage string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("Error %s", e.ErrorMessage)
}

func main() {
	analizator := Analizator{Results: make(map[string]interface{})}

	file, err := os.Open("anything.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer func() {
		err = file.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println("Closing file:", file.Name())
	}()

	builder := strings.Builder{}
	scanner := bufio.NewScanner(file)
	var fileText string
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
		builder.WriteRune('\n')
	}
	fileText = builder.String()
	builder.Reset()

	analizator.OriginalText = fileText
	analizator.Statistics()
	analizator.CalculateHash()
	analizator.Encode()
	err = analizator.PrintAnalysis()
	if err != nil {
		fmt.Println(err)
	}
}

func countChars(arr []string) (int, string) {
	var count int
	maxLenWord := arr[0]
	for _, v := range arr {
		count += len(v)
		if len(maxLenWord) < len(v) {
			maxLenWord = v
		}
	}
	return count, maxLenWord
}
