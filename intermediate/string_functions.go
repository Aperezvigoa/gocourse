package intermediate

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {

	str1 := "A nice blue gopher"

	fmt.Println("Length:", len(str1))

	str2 := "Go"
	str3 := "lang"
	str4 := str2 + str3

	fmt.Println(str4)
	fmt.Println("str4[:2]:", str4[:2])

	// ---String Conversion
	num := 18
	numToStr := strconv.Itoa(num)
	fmt.Printf("Now %s is a %T\n", numToStr, numToStr)

	// ---String Splitting
	fruits := "Apple-Orange-Banana"
	fmt.Println("Fruits before split:", fruits)
	fruitParts := strings.Split(fruits, "-")
	fmt.Println("Fruits after split:", fruitParts)

	countries := []string{"Spain", "Germany", "France", "Italy"}
	fmt.Println("Countries before join:", countries)
	joinedCountries := strings.Join(countries, "-")
	fmt.Println("Countries after join:", joinedCountries)

	fmt.Println("joinedCountries contains Spain?", strings.Contains(joinedCountries, "Spain"))
	fmt.Println("Lets replace France.")
	joinedCountries = strings.Replace(joinedCountries, "France", "Norway", 1)
	fmt.Println(joinedCountries)

	// ---String Trimming
	strwspace := " Hello World"
	fmt.Println("strwspace before trim:", strwspace)
	strwspace = strings.TrimSpace(strwspace)
	fmt.Println("strwspace after trim:", strwspace)

	fmt.Println("HELLO to lowercase:", strings.ToLower("HELLO"))
	fmt.Println("hello to uppercase:", strings.ToUpper("hello"))

	whatILove := "I love my cats!"
	fmt.Println(strings.Repeat(whatILove, 20))

	banana := "Banana"
	fmt.Println("Count of A in Banana:", strings.Count(banana, "a"))

	fmt.Println("Banana has prefix Ba:", strings.HasPrefix(banana, "Ba"))
	fmt.Println("Banana has sufix na:", strings.HasSuffix(banana, "na"))

	str5 := "Hello, 123 Go! 11"
	re := regexp.MustCompile(`\d+`)
	fmt.Println(re.FindAllString(str5, -1))

	str6 := "Hello, 2010 Go! 2025"
	fmt.Println(utf8.RuneCountInString(str6))

	// String builder

	var builder strings.Builder

	// Writing some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("world!")

	// Convert builder to a string
	result := builder.String()
	fmt.Println(result)

	builder.Reset()
	// Using Writerune to add a character
	builder.WriteRune('a')

	builder.WriteString("How are you?")

	result = builder.String()
	fmt.Println(result)
}
