package main

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

type Borrower interface {
	CanBorrow() bool
	GetUserType() string
}

type Book struct {
	ID          string
	Title       string
	Author      string
	IsAvailable bool
}

type BookError struct {
	message string
}

func (be BookError) Error() string {
	return be.message
}

type User struct {
	UserID   string
	Name     string
	Borrowed int
}

type Student struct {
	User
}

type Professor struct {
	User
}

func (s *Student) CanBorrow() bool {
	if s.Borrowed < 3 {
		return true
	} else {
		userAlert.Execute(os.Stdout, s)
		return false
	}
}

func (p *Professor) CanBorrow() bool {
	if p.Borrowed < 5 {
		return true
	} else {
		userAlert.Execute(os.Stdout, p)
		return false
	}
}

func (s Student) GetUserType() string {
	return fmt.Sprintf("User level: %T", s)
}

func (p Professor) GetUserType() string {
	return fmt.Sprintf("User level: %T", p)
}

var userAlert, _ = template.New("alert").Parse("\nUser: {{.Name}} ({{.Type}})\nBorrowed books: {{.Borrowed}}\n\n")
var bookResume, _ = template.New("book").Parse("\nBook ID: {{.ID}}\nTitle: {{.Title}}\nAuthor: {{.Author}}\nIs Available: {{.IsAvailable}}\n\n")
var userReport, _ = template.New("report").Parse("\nUserID: {{.UserID}}\nName: {{.Name}}\nTotal Borrowed: {{.Borrowed}}\n\n")

func (b Book) ShowBookInfo() error {
	err := bookResume.Execute(os.Stdout, b)

	if err != nil {
		return BookError{
			"Book not found",
		}
	}
	return nil
}

func main() {

	var library []Book
	user := Professor{
		User: User{
			UserID:   NewUserID(),
			Name:     "Alex",
			Borrowed: 0,
		},
	}

	for {
		fmt.Println("===LIBRARY MANAGEMENT SYSTEM===")
		fmt.Println("1. Add Book")
		fmt.Println("2. Search Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. User Report")
		fmt.Println("5. Exit")

		fmt.Print("Write your choice: ")
		userChoice, _ := reader.ReadString('\n')
		userChoice = strings.TrimSpace(userChoice)

		switch userChoice {
		case "1":
			registeredBook := AddBook()
			library = append(library, registeredBook)
		case "2":
			SearchBookInfo(library)
		case "3":
			choiceId := requestBookId()
			requestedBook, err := getBookById(choiceId, library)
			canBorrow := user.CanBorrow()
			if err != nil && canBorrow {
				fmt.Println(err)
			} else {
				if requestedBook.IsAvailable {
					user.Borrowed++
				}
				BorrowBook(requestedBook)
			}
		case "4":
			fmt.Println(user.GetUserType())
			err := userReport.Execute(os.Stdout, user)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func generateIDGen() func() string {
	id := 0
	return func() string {
		id++
		return fmt.Sprintf("BK-%04d", id)
	}
}

func generateUserID() func() string {
	id := 0
	return func() string {
		return fmt.Sprintf("USR-%04d", id)
	}
}

var NewBookID = generateIDGen()
var NewUserID = generateUserID()

func AddBook() Book {
	id, title, author, isAvailable := requestBookData()
	return Book{
		ID:          id,
		Title:       title,
		Author:      author,
		IsAvailable: isAvailable,
	}
}

func SearchBookInfo(l []Book) {
	choiceId := requestBookId()
	for _, b := range l {
		if choiceId == b.ID {
			b.ShowBookInfo()
		}
	}
}

func getBookById(id string, library []Book) (*Book, error) {
	for i, b := range library {
		if id == b.ID {
			return &library[i], nil
		}
	}
	return &Book{}, BookError{"Book not found"}
}

func requestBookId() string {
	fmt.Print("Book ID: ")
	choiceId, _ := reader.ReadString('\n')
	choiceId = strings.TrimSpace(choiceId)

	return choiceId
}

func BorrowBook(book *Book) {
	book.IsAvailable = false
}

func requestBookData() (string, string, string, bool) {
	fmt.Print("Name: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Author: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	return NewBookID(), title, author, true
}
