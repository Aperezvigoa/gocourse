package intermediate

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// Example of paths:
	relativePath := "./data/file.txt"
	absolutePath := "/home/user/docs/file.txt"

	// Join paths using filepath.Join
	joinedPath := filepath.Join("home", "Documents", "downloads", "file.zip")
	fmt.Println("Joined Path:", joinedPath)

	// Cleaning & correcting path
	normalizedPath := filepath.Join("./data/../data/file.txt")
	fmt.Println("Normalized Path:", normalizedPath)

	// Spliting a filepath to get the dir and file
	dir, file := filepath.Split("/home/Documents/Golang/file.txt")
	fmt.Println("Dir:", dir, "File:", file)
	fmt.Println(filepath.Base("/home/Documents/Golang/file.txt"))

	// Check if a filepath is absolute or relative path
	fmt.Println("Is absolutePath variable filepath absolute?", filepath.IsAbs(absolutePath))
	fmt.Println("Is relativePath variable filepath absolute?", filepath.IsAbs(relativePath))

	// Getting the file extension from a filepath
	fmt.Println("File extension:", filepath.Ext(absolutePath))

	fmt.Println(strings.TrimSuffix(file, filepath.Ext(file)))

	// Getting relative filepath
	relPath, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Relative Path:", relPath)

	relPath, err = filepath.Rel("a/c", "a/b/t/file")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Relative Path:", relPath)
}
