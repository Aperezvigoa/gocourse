package intermediate

import "fmt"

func main() {

	var firstPointer *int // Pointer declaration
	var simpleInteger int = 10
	firstPointer = &simpleInteger // Pointer assignation

	fmt.Println("SimpleInteger Value:", simpleInteger)
	fmt.Println("SimpleInteger Address:", &simpleInteger)
	fmt.Println("FirstPointer:", firstPointer)
	fmt.Println("FirstPointer dereference:", *firstPointer)

	var nilPointer *int
	fmt.Println("NilPointer:", nilPointer) // NilPointer: <nil>

	modifyValue(nilPointer)
	modifyValue(firstPointer)
}

func modifyValue(ptr *int) {
	defer fmt.Println("ModifyValue function was executed.")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if ptr != nil {
		*ptr *= 2
		fmt.Println("Modification was successfull.")
	} else {
		panic("The pointer was <nil>!!")
	}
}
