package intermediate

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
)

//go:embed example.txt
var content string

//go:embed
var basics embed.FS

func main() {
	fmt.Println("Embeded content:", content)

	readed, err := basics.ReadFile("basics/hello.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Embedded file content:", string(readed))

	err = fs.WalkDir(basics, "basics", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
