package basics

import "fmt"

var inventory = map[string]int{}
var actions []string

func main() {
	for {
		printOptions()
		var option int = selectOption()
		if option == 5 {
			break
		} else {
			run(option)
		}
	}
}

func addProduct(inventory map[string]int) {
	var product = requestProductName()
	var quantity = requestQuantity()
	inventory[product] = quantity
}

func rmvProduct(inventory map[string]int) {
	product := requestProductName()
	delete(inventory, product)
}

func printInventory(inventory map[string]int) {
	fmt.Println("Inventory:")
	for k, v := range inventory {
		fmt.Printf("Product: %v ----> Stock: %d\n", k, v)
	}
}

func requestProductName() string {
	fmt.Println("Escribe el nombre del producto:")
	var productName string
	fmt.Scanln(&productName)
	return productName
}

func requestQuantity() int {
	fmt.Println("Escribe la cantidad:")
	var productQuantity int
	fmt.Scanln(&productQuantity)
	return productQuantity
}

func selectOption() int {
	fmt.Println("Select an option:")
	var option int
	fmt.Scanln(&option)
	return option
}

func printLog() {
	for i, v := range actions {
		fmt.Println("----------------------------")
		fmt.Printf("%d - %v\n", i+1, v)
		fmt.Println("----------------------------")
	}
}

func printOptions() {
	fmt.Println("\n1. Add Product")
	fmt.Println("2. Quit Product")
	fmt.Println("3. See Inventory")
	fmt.Println("4. Log")
	fmt.Println("5. Exit\n")
}

func run(option int) {
	switch option {
	case 1:
		addProduct(inventory)
		actions = append(actions, "Added product.")
	case 2:
		rmvProduct(inventory)
		actions = append(actions, "Removed product.")
	case 3:
		printInventory(inventory)
		actions = append(actions, "Printed product.")
	case 4:
		actions = append(actions, "Printed product.")
		printLog()
	default:
		fmt.Println("Invalid option...")
	}
}
