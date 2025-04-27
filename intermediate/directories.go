package intermediate

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// --- Creating directory
	// checkError(os.Mkdir("subdir", 0755))

	// --- Deleting directory
	// defer checkError(os.RemoveAll("subdir"))

	// --- Creating files
	// os.WriteFile("subdir/data", []byte(""), 0755)

	// --- Creating directory with subdirectories
	// checkError(os.MkdirAll("newSubDir/data/test", 0755))
	// checkError(os.MkdirAll("newSubDir/data/test1", 0755))
	// checkError(os.MkdirAll("newSubDir/data/test2", 0755))
	// checkError(os.MkdirAll("newSubDir/data/test3", 0755))
	// os.WriteFile("newSubDir/data/test/test1.txt", []byte("test1"), 0755)

	result, err := os.ReadDir("newSubDir/data/")
	checkError(err)

	for _, entry := range result {
		fmt.Println(entry.Name(), entry.Type(), entry.IsDir())
	}

	// --- Change directory
	err = os.Chdir("newSubDir/data")
	checkError(err)

	// Now we are working on "newSubDir/data" as default
	result, err = os.ReadDir(".")
	checkError(err)
	fmt.Println("====================================")
	for _, entry := range result {
		fmt.Println(entry.Name(), entry.Type(), entry.IsDir())
	}

	// Changing directory, going back 2 levels
	err = os.Chdir("../../")
	checkError(err)
	fmt.Println("====================================")
	directory, err := os.Getwd() // Get working directory
	checkError(err)
	fmt.Println(directory)

	fmt.Println("====================================")

	// filepath.Walk and filepath.WalkDir
	pathfile := "newSubDir/"

	err = filepath.WalkDir(pathfile, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		fmt.Println(path)
		return nil
	})
	checkError(err)
	err = os.RemoveAll("newSubDir")
	checkError(err)
}
