package basics

import (
	"fmt"
	"maps"
)

func main() {
	goMap := make(map[string]int)
	fmt.Println(goMap)

	goMap["Monday"] = 29
	goMap["Tuesday"] = 27
	goMap["Wednesday"] = 28
	goMap["Thursday"] = 23
	goMap["Friday"] = 26
	fmt.Println(goMap)

	fmt.Println("Friday temperature:", goMap["Friday"], "degrees")

	goMap["Saturday"] = 22
	fmt.Println("Added Saturday:", goMap)

	// Deleting a key
	delete(goMap, "Saturday")
	fmt.Println("Removed Saturday:", goMap)

	// Clearing map
	clear(goMap)
	fmt.Println("Cleared map:", goMap)

	value, ok := goMap["Friday"]
	fmt.Println("Value:", value)
	fmt.Println("Is a value associated with key 'Friday'?", ok)

	goMapTwo := map[int]string{
		1: "Cheddar",
		2: "Whiskers",
		3: "Lily",
	}
	fmt.Println("GoMapTwo:", goMapTwo)

	copyMap := map[int]string{}

	maps.Copy(copyMap, goMapTwo)
	fmt.Println("copyMap:", copyMap)

	if maps.Equal(copyMap, goMapTwo) {
		fmt.Println("CopyMaps is equal to GoMapTwo")
	}

	fmt.Println("Copy map:")
	for k, v := range copyMap {
		fmt.Printf("Key: %d - Value: %v\n", k, v)
	}

	var goMapThree map[string]string
	if goMapThree == nil {
		fmt.Println("GoMapThree is initialized to nil value")
	} else {
		fmt.Println("GoMapThree is not initialized to nil value")
	}

	nilVar := goMapThree["a"]
	fmt.Println(nilVar)

	goMapFour := make(map[int]string)

	goMapFour[1] = "Aperezvigoa"

	fmt.Println("GoMapFour length:", len(goMapFour))

	// Multidimensional Map

	goMapFive := make(map[string]map[int]string)

	goMapFive["MapOne"] = goMapFour
	goMapFive["MapTwo"] = goMapTwo

	fmt.Println(goMapFive)
}
