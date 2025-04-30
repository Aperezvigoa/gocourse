package intermediate

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name" db:"first_name" xml:"first_name"` // Multiple tags
	LastName  string `json:"last_name"`
	Age       int    `json:"age,omitempty"`
	Omitted   string `json:"-"`
}

func main() {
	person1 := Person{FirstName: "John", LastName: "Doe", Age: 42}

	jsonData1, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println(string(jsonData1))

}
