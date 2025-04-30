package intermediate

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string  `json:"name"`
	Age       int     `json:"age,omitempty"` // If Age is not declared, omit it
	Email     string  `json:"email"`
	Address   Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

func main() {
	person := Person{FirstName: "John", Age: 23, Email: "johndoe@gmail.com"}

	// Marshalling
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println(string(jsonData))

	person1 := Person{
		FirstName: "Jane",
		Age:       23,
		Email:     "janesmith@git.io",
		Address: Address{
			City:  "New York",
			State: "New York",
		},
	}

	jsonData1, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println(string(jsonData1))

	// JSON Unmarshalling
	jsonData2 := `{"full_name": "Jenny Doe", "emp_id": "0009", "age": 30, "address": {"city": "San Jose", "state": "CA"}}`
	employeeFromJson := Employee{}
	err = json.Unmarshal([]byte(jsonData2), &employeeFromJson)
	if err != nil {
		fmt.Println("Error unmarshalling to JSON:", err)
		return
	}
	fmt.Println(employeeFromJson)
	fmt.Println("Jane's Age increased by 5 years", employeeFromJson.Age+5)
	fmt.Println("Janeś City", employeeFromJson.Address.City)

	// Encoding a list of Address instances into a JSON string
	listOfCityState := []Address{
		{City: "Las Vegas", State: "NV"},
		{City: "New York", State: "NY"},
		{City: "San Jose", State: "CA"},
		{City: "Modesto", State: "CA"},
		{City: "Clearwater", State: "FL"},
	}
	jsonAddresses, err := json.Marshal(listOfCityState)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println(string(jsonAddresses))

	// Handling unknown JSON structures
	jasonData3 := `{"name": "John", "age": 23, "address": {"city": "New York", "state": "NY"}}`

	var data map[string]interface{}

	err = json.Unmarshal([]byte(jasonData3), &data)
	if err != nil {
		fmt.Println("Error unmarshalling to JSON:", err)
		return
	}
	fmt.Println("Handling Unknown JSON Structure")
	fmt.Println(data)
}

type Employee struct {
	FullName string  `json:"full_name"`
	EmpID    string  `json:"emp_id"`
	Age      int     `json:"age"`
	Address  Address `json:"address"`
}
