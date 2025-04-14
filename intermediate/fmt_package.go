package intermediate

import "fmt"

func main() {

	// --- Printing Functions
	fmt.Print("I love ")
	fmt.Print("Golang ")
	fmt.Print(3.14159265359)

	fmt.Println()

	fmt.Println("I love ")
	fmt.Println("Golang ")
	fmt.Println(3.14159265359)

	fmt.Println()

	name := "Alex"
	age := 23
	fmt.Printf("My name is %s and my age is %d\n", name, age)
	fmt.Printf("Binary: %b, Hex: %X\n", age, age)

	// --- Formatting Functions
	s := fmt.Sprint("Hello ", "world!", 123, 456)
	fmt.Print(s)
	s = fmt.Sprintln("Hello", "world!", 123, 456)
	fmt.Print(s)

	name = "Cheddar"
	age = 5

	myCat := fmt.Sprintf("Name: %s, Age: %d\n", name, age)
	fmt.Print(myCat)

	// --- Scanning Functions
	var userName string
	var userAge int

	fmt.Println("Enter your username and your age:")
	//fmt.Scan(&userName, &userAge)
	// fmt.Scanln(&userName, &userAge)
	fmt.Scanf("%s %d", &userName, &userAge)
	fmt.Printf("Username: %s User Age: %d\n", userName, userAge)

	// --- Error Formatting Functions
	err := checkAge(15)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func checkAge(age int) error {
	if age < 18 {
		fmt.Errorf("Age %d is too young to drive.", age)
	}
	return nil
}
