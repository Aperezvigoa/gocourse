package intermediate

import (
	"bufio"
	"fmt"
	"html/template"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Ingredients struct {
	Name  string
	Type  string
	Power int
	Rune  rune
}

type PotionReport struct {
	MixtureNum int
	Date       string
	PotionA    string
	PotionB    string
	Result     string
	Potency    int
}

var potions = map[string]Ingredients{}

var mixture = func() func() int {
	numberOfMixture := 0
	return func() int {
		numberOfMixture++
		return numberOfMixture
	}
}

type customError struct {
	text string
}

func (e customError) Error() string {
	return e.text
}

var getMixture = mixture()

var reader = bufio.NewReader(os.Stdin)

func init() {
	potions["Potion of Invisibility"] = Ingredients{
		Name:  "Potion of Invisibility",
		Type:  "Stealth",
		Power: rand.Intn(100) + 1,
		Rune:  '🧿',
	}
	potions["Elixir of Eternal Youth"] = Ingredients{
		Name:  "Elixir of Eternal Youth",
		Type:  "Restoration",
		Power: rand.Intn(100) + 1,
		Rune:  '🧬',
	}
	potions["Draught of Strength"] = Ingredients{
		Name:  "Draught of Strength",
		Type:  "Enhancement",
		Power: rand.Intn(100) + 1,
		Rune:  '💪',
	}
	potions["Essence of Night Vision"] = Ingredients{
		Name:  "Essence of Night Vision",
		Type:  "Perception",
		Power: rand.Intn(100) + 1,
		Rune:  '🔮',
	}
	potions["Potion of Fire Resistance"] = Ingredients{
		Name:  "Potion of Fire Resistance",
		Type:  "Protection",
		Power: rand.Intn(100) + 1,
		Rune:  '🫧',
	}
	potions["Brew of Healing"] = Ingredients{
		Name:  "Brew of Healing",
		Type:  "Healing",
		Power: rand.Intn(100) + 1,
		Rune:  '💉',
	}
	potions["Tonic of Speed"] = Ingredients{
		Name:  "Tonic of Speed",
		Type:  "Agility",
		Power: rand.Intn(100) + 1,
		Rune:  '🎿',
	}
}

func main() {
	// --- Creating the result file
	results, err := os.Create("potions.txt")
	CheckError(err)

	defer func() {
		err = results.Close()
		CheckError(err)
		fmt.Println("Done writing potions.txt")
	}()

	defer fmt.Println("Thank you for visit the wizard of Goland")
	for {
		PrintMenu()
		choice, err := reader.ReadString('\n')
		CheckError(err)
		choice = strings.TrimSpace(choice)
		switch choice {
		case "1":
			SeePotions()
		case "2":
			a, b, err := SelectPotions()
			if err != nil {
				fmt.Println(err)
				continue
			}
			pr, err := MixturePotions(a, b, getMixture())
			if err != nil {
				fmt.Println(err)
			}
			ReportMixture(pr, results)

		case "3":
			return
		}
	}
}

func PrintMenu() {
	fmt.Println("\n1. See Potions")
	fmt.Println("2. Mix potions")
	fmt.Println("3. Exit")
}

func SeePotions() {
	index := 1
	for name, ingredients := range potions {
		fmt.Println("=====================")
		fmt.Printf("%d - %s %c:\nType: %s\nPower: %d\n", index, name, ingredients.Rune, ingredients.Type, ingredients.Power)
		index++
	}
}

func SelectPotions() (Ingredients, Ingredients, error) {
	fmt.Printf("Select a Potion, entering the name: ")
	choiceA, _ := reader.ReadString('\n')
	fmt.Printf("Select another Potion, entering the name: ")
	choiceB, _ := reader.ReadString('\n')
	choiceA = strings.TrimSpace(choiceA)
	choiceB = strings.TrimSpace(choiceB)

	valueA, existsA := potions[choiceA]
	valueB, existsB := potions[choiceB]

	if !existsA || !existsB {
		return valueA, valueB, customError{"❗🧪 One or more potions were not found... Check your ingredients! 🧪❗\n"}
	}
	return valueA, valueB, nil
}

func MixturePotions(a Ingredients, b Ingredients, mixtureNum int) (PotionReport, error) {
	defer fmt.Println("🧹✨ Cleaning the cauldron... Ready for the next brew! ✨🧹")
	potionReport := PotionReport{
		MixtureNum: mixtureNum,
		Date:       time.Now().Format("2006-01-02 15:04:05"),
		PotionA:    a.Name,
		PotionB:    b.Name,
	}
	temperature := rand.Intn(10) + 25

	fmt.Printf("Mixing... Temperature: %dºC\n", temperature)
	for i := 0; i < 7; i++ {
		time.Sleep(2 * time.Second)
		temperature += rand.Intn(10) + 5
		fmt.Printf("Turning up the heat at %dºC\n", temperature)
	}

	switch {
	case temperature >= 100:
		fmt.Println("🧨 EXPLOSION!!!!!!! 🎇\nThe mixture failed...")
		potionReport.Result = "FAILURE"
		potionReport.Potency = 0
		return potionReport, customError{"💥🧪💀 BOOM! The cauldron has exploded! Magical disaster unleashed! 💀🧪💥\n"}
	case temperature >= 80:
		fmt.Println("🎆 CONGRATULATIONS! 🎆\n✨The mixture succeeded...✨")
		potionReport.Result = "SUCCESS"
		potionReport.Potency = temperature
		return potionReport, nil
	default:
		fmt.Println("🔴 Can't mixture the potions... We need more temperature 🔴")
		return potionReport, customError{"⚠️🧪 The potions couldn't be mixed... More heat or magic is needed! 🧪⚠️\n"}
	}
}

func ReportMixture(potionReport PotionReport, file *os.File) {
	report := template.Must(template.New("potionreport").Parse(
		`=== 🧙‍♂️ Poción número {{.MixtureNum}} 🧙‍♂️ ===
Fecha: {{.Date}}
Ingredientes: {{.PotionA}} & {{.PotionB}}
Resultado: {{.Result}}
Potencia: {{.Potency}}/100
`))

	err := report.Execute(file, potionReport)
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
