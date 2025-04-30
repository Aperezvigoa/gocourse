package intermediate

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	XMLName xml.Name `xml:"person"` // XML name field is a special field used by the encoding XML package
	Name    string   `xml:"name"`   // And it's used to determine the name of the XML when marshalling
	Age     int      `xml:"age"`
	Email   string   `xml:"email,omitempty"`
	Address Address  `xml:"address"`
}

type Address struct {
	City   string `xml:"city"`
	Street string `xml:"street,omitempty"`
}

func main() {
	person1 := Person{
		Name:  "John",
		Age:   25,
		Email: "johndoe@git.io",
		Address: Address{
			Street: "123 St",
			City:   "London",
		},
	}
	// Marshal XML
	xmlData1, err := xml.Marshal(&person1)
	if err != nil {
		log.Fatalln("Failed to marshal xml:", err)
		return
	}

	fmt.Println("XML:", string(xmlData1))

	// Marshal XML
	xmlData1, err = xml.MarshalIndent(&person1, "", "\t")
	if err != nil {
		log.Fatalln("Failed to marshal xml:", err)
		return
	}
	fmt.Println("XML Indent:\n", string(xmlData1))

	// Unmarshal XML
	xmlRaw1 := "<person><name>Alex</name><age>23</age><email>aperezvigoa@github.com</email><address><city>Barcelona</city><street>C/ Bernabeu</street></address></person>"
	var person2 Person
	err = xml.Unmarshal([]byte(xmlRaw1), &person2)
	if err != nil {
		log.Fatalln("Failed to unmarshal xml:", err)
	}
	fmt.Println("Unmarshalled XML:", person2)
	fmt.Println("Local:", person2.XMLName.Local)
	fmt.Println("Space:", person2.XMLName.Space)

	book := Book{
		ISBN:   "B4-59-0F-00-1A",
		Author: "Cervantes",
		Title:  "Don Quijote De La Mancha",
		Pseudo: Pseudo{
			Pseudo:     "Pseudo",
			PseudoAttr: "PseudoAttr",
		},
	}

	xmlBook, err := xml.MarshalIndent(&book, " * ", "  ")
	if err != nil {
		log.Fatalln("Failed to marshal xml:", err)
	}
	fmt.Println(string(xmlBook))
}

type Book struct {
	XMLName xml.Name `xml:"book"`
	ISBN    string   `xml:"isbn,attr"` // attr Make ISBN an ATTRIBUTE NOT A CHILD
	Title   string   `xml:"title,attr"`
	Author  string   `xml:"author,attr"`
	Pseudo  Pseudo   `xml:"pseudo"`
}

type Pseudo struct {
	Pseudo     string `xml:"pseudo"`
	PseudoAttr string `xml:"pseudoattr,attr"`
}

// <book isbn="B4-59-0F-00-1A" title="Don Quijote De La Mancha" author="Cervantes"></book>
