package intermediate

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
	PhoneHomeCell
}

type Address struct {
	city    string
	country string
}

type PhoneHomeCell struct {
	home  string
	phone string
}

func main() {

	me := Person{
		firstName: "Alex",
		lastName:  "Perez-Vigo",
		age:       23,
		address: Address{
			city:    "Madrid",
			country: "Spain",
		},
		PhoneHomeCell: PhoneHomeCell{
			phone: "Fic number",
			home:  "Fic home",
		},
	}

	randomPerson := Person{
		firstName: "Michael",
		lastName:  "Jackson",
		age:       35,
	}
	randomPerson.address.city = "New York"
	randomPerson.address.country = "United States of America"

	fmt.Println(me)
	fmt.Println(randomPerson)
	fmt.Println("My first name is:", me.firstName)
	fmt.Println(me.fullName())

	fmt.Println("My actual age:", me.age)
	me.incrementAge(10)
	fmt.Println("In 10 years:", me.age)
	fmt.Println(me.home)

	fmt.Println("Are we equal?", me == randomPerson)

	// --- Anonymous Structs
	user := struct {
		username string
		email    string
	}{
		username: "user123",
		email:    "user123@example.io",
	}
	_ = user
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAge(years int) {
	p.age += years
}
