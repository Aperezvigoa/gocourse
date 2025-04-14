package intermediate

import "fmt"

type RegisteredUser struct {
	name     string
	lastName string
	email    string
}

type Shape struct {
	Rectangle
}

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.length
}

// Method with value reciever
func (ru RegisteredUser) GenerateUsername() string {
	return ru.name + ru.lastName
}

// Method with pointer reciever

func (ru *RegisteredUser) ModifyEmail(newMail string) {
	ru.email = newMail
}

func main() {

	firstUser := RegisteredUser{
		name:     "John",
		lastName: "Doe",
		email:    "johndoe@github,com",
	}

	fmt.Printf("The username of %s is %s\n", firstUser.name, firstUser.GenerateUsername())

	fmt.Printf("The actual mail of %s is %s\n", firstUser.name, firstUser.email)
	firstUser.ModifyEmail("anothermail@golang.dev")
	fmt.Printf("The new mail of %s is %s\n", firstUser.name, firstUser.email)

	num := MyInt(-5)
	fmt.Println("Is num positive?", num.isPositive())

	s := Shape{
		Rectangle{
			length: 10,
			width:  20,
		},
	}
	fmt.Println("Area:", s.Area())
	fmt.Println("Area:", s.Rectangle.Area())
}

type MyInt int

// Method on a user-defined type
func (m MyInt) isPositive() bool {
	return m > 0
}
