package basics

import "fmt"

const MAXPLAYERS = 3

var ranks = [3]string{"Gold", "Silver", "Bronze"}
var registeredPlayers = [MAXPLAYERS]string{}
var playerStatistics = make(map[string][2]int)

func main() {

	defer fmt.Println("Thank you for use the player statistics system")

	for i := 0; i < MAXPLAYERS; i++ {
		registerNewPlayer(i, requestName())
	}
	printPlayerStatistics()
}

func requestName() string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var name string
	fmt.Print("Name: ")
	fmt.Scanln(&name)

	if checkRegisteredPlayers(name) {
		panic("The player is registered...")
	}

	return name
}

func giveRank(puntuation int) {
	switch {
	case puntuation >= 500:
		fmt.Printf("\nThe player is registered with the rank: %v\n\n", ranks[0])
	case puntuation >= 300 && puntuation < 500:
		fmt.Printf("\nThe player is registered with the rank: %v\n\n", ranks[1])
	default:
		fmt.Printf("\nThe player is registered with the rank: %v\n\n", ranks[2])
	}
}

func printPlayerStatistics() {
	fmt.Println("Final Statistics:")
	fmt.Println("=================")

	fmt.Println("Average score:", getAveragescore())
	fmt.Println("Max score:", getMaxscore())
}

func getAveragescore() int {
	var average int
	for _, v := range playerStatistics {
		average += v[1]
	}
	return average / MAXPLAYERS
}

func getMaxscore() int {
	var max int
	for _, v := range playerStatistics {
		if v[1] > max {
			max = v[1]
		}
	}
	return max
}

func registerNewPlayer(i int, name string) {
	levelAndPuntutation := [2]int{}
	registeredPlayers[i] = name
	fmt.Print("Level: ")
	fmt.Scanln(&levelAndPuntutation[0])
	fmt.Print("score: ")
	fmt.Scanln(&levelAndPuntutation[1])

	playerStatistics[name] = levelAndPuntutation
	giveRank(levelAndPuntutation[1])
}

func checkRegisteredPlayers(name string) bool {

	for _, v := range registeredPlayers {
		if v == name {
			return true
		}
	}
	return false
}
