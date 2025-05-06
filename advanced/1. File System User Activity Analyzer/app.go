package exercise

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

type User struct {
	Connections  []time.Time
	Username     string
	TotalActions []string
	Files        []string
}

type CustomError struct {
	code    int
	message string
}

func (c CustomError) Error() string {
	return fmt.Sprintf("%d: %s", c.code, c.message)
}

func main() {
	var directory string
	var user string
	var summary bool
	var output string
	var format string

	flag.StringVar(&directory, "dir", "", "Directory to analyze")
	flag.StringVar(&user, "user", "", "Filter by the username")
	flag.BoolVar(&summary, "smmry", false, "Show summary only")
	flag.StringVar(&output, "output", "output.txt", "Output file name")
	flag.StringVar(&format, "format", "txt", "Output format | txt | base64")

	flag.Parse()

	if strings.TrimSpace(directory) == "" {
		log.SetPrefix("FATAL ERROR")
		log.Fatalln("The directory can't be blank!")
		return
	}

	filepaths, err := detectFiles(directory)
	if err != nil {
		fmt.Println(err)
		return
	}

	openedFiles := make([]*os.File, 0)

	for _, p := range filepaths {
		file, err := getFile(p)
		if err != nil {
			panic(err)
		}
		defer func() {
			file.Close()
			log.Println("File closed.")
		}()
		openedFiles = append(openedFiles, file)
	}

	dirtData := extractFileInfo(openedFiles...)
	usersMap := registerUserData(dirtData)

	if summary == false {
		if strings.TrimSpace(user) != "" {
			go func() {
				uFile, err := os.OpenFile("Filtered-User.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
				if err != nil {
					uFile.Close()
					panic(err)
				}
				defer uFile.Close()
				for _, u := range usersMap {
					if u.Username == user {
						filterUser(uFile, u)
					}
				}
			}()
		}

		if format == "base64" {
			go func() {
				bFile, err := os.OpenFile("encoded-file.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
				if err != nil {
					bFile.Close()
					panic(err)
				}
				defer bFile.Close()
				base64enc(bFile)
			}()
		}
	}
	go func() {
		eFile, err := os.OpenFile(output, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
		if err != nil {
			eFile.Close()
			panic(err)
		}
		defer eFile.Close()
		fullExport(eFile, usersMap)
	}()
	time.Sleep(2 * time.Second)
}

func detectFiles(directory string) ([]string, error) {

	filepaths := make([]string, 0)
	err := filepath.WalkDir(directory, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			go func() {
				filepaths = append(filepaths, path)
			}()
		}
		return nil
	})
	time.Sleep(1 * time.Second)
	if err != nil {
		log.SetPrefix("READING ERROR: ")
		log.Println("Something went wrong while reading dir")
		return nil, err
	} else if len(filepaths) == 0 {
		log.SetPrefix("EMPTY DIR ERROR: ")
		log.Println("Can't detect any file in the directory")
		return nil, CustomError{code: 1, message: "Can't detect any file in the directory"}
	} else {
		log.SetPrefix("SUCCESS: ")
		log.Println("Filepaths readed successfully")
		return filepaths, nil
	}
}

func getFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		log.SetPrefix("File Error")
		log.Println("Error opening file")
		return nil, err
	}
	return file, nil
}

func extractFileInfo(files ...*os.File) string {
	var builder strings.Builder

	for _, f := range files {
		scanner := bufio.NewScanner(f)
		go func() {
			for scanner.Scan() {
				builder.WriteString(scanner.Text() + "\n")
			}
		}()
		err := scanner.Err()
		if err != nil {
			panic(err)
		}
		time.Sleep(500 * time.Microsecond)
	}

	return builder.String()
}

func registerUserData(data string) map[string]User {
	registeredUsers := make(map[string]User, 0)

	linebyline := strings.Split(strings.TrimSpace(data), string('\n'))

	for _, i := range linebyline {
		splitData := strings.Split(i, "|")

		if _, ok := registeredUsers[strings.TrimSpace(splitData[1])]; !ok {
			connection, err := time.Parse("2006-01-02T15:04:05Z", strings.TrimSpace(splitData[0]))
			if err != nil {
				log.Fatalln("Error parsing connection time!")
			}
			splitedName := strings.Split(strings.TrimSpace(splitData[1]), ":")
			splitedAction := strings.Split(strings.TrimSpace(splitData[2]), ":")
			user := User{
				Connections:  []time.Time{connection},
				Username:     splitedName[1],
				TotalActions: []string{splitedAction[1]},
				Files:        []string{strings.TrimSpace(splitData[3])},
			}
			registeredUsers[strings.TrimSpace(splitData[1])] = user
		} else {
			go func() {
				user := registeredUsers[strings.TrimSpace(splitData[1])]
				connection, err := time.Parse("2006-01-02T15:04:05Z", strings.TrimSpace(splitData[0]))
				if err != nil {
					log.Fatalln("Error parsing connection time!")
				}
				splitedAction := strings.Split(strings.TrimSpace(splitData[2]), ":")

				user.Connections = append(user.Connections, connection)
				user.TotalActions = append(user.TotalActions, splitedAction[1])
				user.Files = append(user.Files, strings.TrimSpace(splitData[3]))

				registeredUsers[strings.TrimSpace(splitData[1])] = user
			}()

			time.Sleep(200 * time.Millisecond)
		}
	}
	return registeredUsers
}

func fullExport(f *os.File, users map[string]User) {
	fullExport := template.Must(template.New("full-export").Parse(`
User Profile: {{.Username}}

Connections:
{{range .Connections}}
  - {{.Format "02-Jan-2006 15:04:05"}}
{{else}}
  No connections found.
{{end}}

Total Actions:
{{range .TotalActions}}
  - {{.}}
{{else}}
  No actions recorded.
{{end}}

Files:
{{range .Files}}
  - {{.}}
{{else}}
  No files uploaded.
{{end}}
`))
	for _, u := range users {
		err := fullExport.Execute(f, u)
		if err != nil {
			panic(err)
		}
	}
}

func filterUser(f *os.File, user User) {
	temp := template.Must(template.New("only-user").Parse(`
==========================================
              User Profile
==========================================
Username:  {{.Username}}

------------------------------------------
Connections:
------------------------------------------
{{range .Connections}}
  - {{.Format "02-Jan-2006 15:04:05"}}
{{else}}
  No connections found.
{{end}}

------------------------------------------
Total Actions:
------------------------------------------
{{range .TotalActions}}
  - {{.}}
{{else}}
  No actions recorded.
{{end}}

------------------------------------------
Files Uploaded:
------------------------------------------
{{range .Files}}
  - {{.}}
{{else}}
  No files uploaded.
{{end}}

==========================================
`))
	err := temp.Execute(f, user)
	if err != nil {
		panic(err)
	}
}

func base64enc(f *os.File) {
	fullTemp := `
User Profile: {{.Username}}
Connections:
{{range .Connections}}
  - {{.Format "02-Jan-2006 15:04:05"}}
{{else}}
  No connections found.
{{end}}

Total Actions:
{{range .TotalActions}}
  - {{.}}
{{else}}
  No actions recorded.
{{end}}

Files:
{{range .Files}}
  - {{.}}
{{else}}
  No files uploaded.
{{end}}`

	_, err := f.Write([]byte(base64.StdEncoding.EncodeToString([]byte(fullTemp))))
	if err != nil {
		panic(err)
	}
}
