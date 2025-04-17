package intermediate

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Names struct {
}

func main() {

	// First way
	/*
		// tmpl := template.New("example")

		tmpl, err := template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n")

		if err != nil {
			panic(err)
		}

		// Defining data for welcome message template
		data := map[string]interface{}{
			"name": "John",
		}
		data1 := map[string]any{
			"name": "Alex",
		}

		data2 := map[string]interface{}{
			"name": "Michael",
		}

		fullData := map[string]map[string]any{
			"data":  data,
			"data1": data1,
			"data2": data2,
		}

		for k := range fullData {
			err = tmpl.Execute(os.Stdout, fullData[k])
			if err != nil {
				panic(err)
			}
		}
	*/

	// Second way

	/*
		tmpl := template.Must(template.New("Example").Parse("Hello {{.name}}! We are testing templates.\n"))

		data := map[string]interface{}{
			"name": "John",
		}

		err := tmpl.Execute(os.Stdout, data)
		if err != nil {
			panic(err)
		}
	*/

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// Define named template for different types of
	templates := map[string]string{
		"welcome":      "Welcome, {{.name}}!, We' re glad you joined.\n",
		"notification": "{{.name}}, you have a new notification: {{.notification}}\n",
		"error":        "Oops! An error occured: {{.errorMessage}}\n",
	}

	// Parsing and storing templates
	parsedTemplates := make(map[string]*template.Template)

	for name, tmpl := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	for {
		// Show menu
		fmt.Println("--- MENU ---")
		fmt.Println("1. Join")
		fmt.Println("2. Get Notification")
		fmt.Println("3. Get Error")
		fmt.Println("4. Exit")

		fmt.Println("Choice an option:")
		userChoice, _ := reader.ReadString('\n')
		userChoice = strings.TrimSpace(userChoice)

		var data map[string]interface{}
		var tmpl *template.Template

		switch userChoice {
		case "1":
			tmpl = parsedTemplates["welcome"]
			data = map[string]interface{}{"name": name}
		case "2":
			fmt.Println("Enter your notification:")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)

			tmpl = parsedTemplates["notification"]
			data = map[string]interface{}{"name": name, "notification": notification}
		case "3":
			fmt.Println("Enter your error:")
			userError, _ := reader.ReadString('\n')
			userError = strings.TrimSpace(userError)

			tmpl = parsedTemplates["error"]
			data = map[string]interface{}{"errorMessage": userError}
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice.")
			continue
		}

		// Render and printing the template
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()
		err := tmpl.Execute(os.Stdout, data)
		if err != nil {
			panic(err)
		}
	}
}
