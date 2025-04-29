package intermediate

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var src string
	var dst string

	flag.StringVar(&src, "src", "", "Source file path")
	flag.StringVar(&dst, "dst", "", "Destination dir")

	flag.Parse()
	if src == "" || dst == "" {
		panic("Source and destination file must be specified")
	}

	srcFile, err := os.Open(src)
	if err != nil {
		panic(err)
	}

	defer func() {
		err = srcFile.Close()
		if err != nil {
			fmt.Println("Error closing source file")
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(srcFile)
	builder := strings.Builder{}
	for scanner.Scan() {
		builder.WriteString(scanner.Text() + "\n")
	}

	err = os.Mkdir(dst, 0755)
	if err != nil {
		panic(err)
	}
	dstFile, err := os.Create(dst + "backup.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		err = dstFile.Close()
		if err != nil {
			panic(err)
		}
	}()

	_, err = dstFile.Write([]byte(builder.String()))
	if err != nil {
		panic(err)
	}
}
