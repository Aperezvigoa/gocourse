package intermediate

import (
	"fmt"
)

func swap[T any](a, b T) (T, T) {
	return b, a
}

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) pop() (T, bool) {

	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func (s Stack[T]) printAll() {
	if len(s.elements) == 0 {
		fmt.Println("Empty")
	} else {
		fmt.Println("Stack elements:")
		for _, v := range s.elements {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}

func main() {
	x := 187
	y := 32

	fmt.Println("X and Y before swap:", x, y)
	x, y = swap(x, y)
	fmt.Println("X and Y after swap:", x, y)

	maleName := "John"
	femaleName := "Hannah"

	fmt.Printf("Male name: %s, Female name: %s\n", maleName, femaleName)
	maleName, femaleName = swap(maleName, femaleName)
	fmt.Printf("Male name: %s, Female name: %s\n", maleName, femaleName)

	intStack := Stack[int]{}

	intStack.push(1)
	intStack.push(10)
	intStack.push(100)

	intStack.printAll()
	intStack.pop()
	intStack.printAll()
	intStack.pop()
	intStack.pop()
	fmt.Println("Stack is empty:", intStack.isEmpty())
	intStack.printAll()

	strStack := Stack[string]{}
	strStack.push("Cheddar")
	strStack.push("Whiskers")
	strStack.push("Lily")

	strStack.printAll()
	strStack.pop()
	strStack.printAll()
	strStack.pop()
	strStack.pop()
	fmt.Println("Stack is empty:", strStack.isEmpty())
	strStack.printAll()
}
