package intermediate

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Creating temporary file
	tempfile, err := os.CreateTemp("", "temporaryFile")
	checkError(err)
	fmt.Println(tempfile.Name())

	// Defer closing and removing temp file
	defer func() {
		err = tempfile.Close()
		checkError(err)
		fmt.Println("Closed file:", tempfile.Name())
		err = os.Remove(tempfile.Name())
		checkError(err)
		fmt.Println("Removed file:", tempfile.Name())
	}()

	// Creating temporary dir
	tmpdir, err := os.MkdirTemp("", "tempDir")
	checkError(err)
	defer func() {
		err = os.RemoveAll(tmpdir)
		checkError(err)
		fmt.Println("Removed dir:", tmpdir)
	}()
	fmt.Println("Created tempDir:", tmpdir)
}
