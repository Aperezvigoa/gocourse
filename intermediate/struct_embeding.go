package intermediate

import "fmt"

type person struct {
	name string
	age  int
}

type employee struct {
	person // Embedded struct
	empId  string
	salary float64
}

func (p person) introduce() {
	fmt.Printf("Hello, I'm %s and I am %d years old.\n", p.name, p.age)
}

func (e employee) introduce() {
	fmt.Printf("Hello, I'm %s, my employeeID is %s and my salary is %2.f\n", e.name, e.empId, e.salary)
}

func main() {
	emp1 := employee{
		person: person{
			name: "Alex",
			age:  23,
		},
		empId:  "E0001",
		salary: 1350.68,
	}

	fmt.Println("Emp1 Name:", emp1.name) // Accessing the embedded struct field directly
	fmt.Println("Emp1 Age:", emp1.age)   // Same above
	fmt.Println("Emp1 ID:", emp1.empId)
	fmt.Println("Emp1 salary:", emp1.salary)

	emp1.introduce()
}
