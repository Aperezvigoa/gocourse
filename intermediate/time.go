package intermediate

import (
	"fmt"
	"time"
)

func main() {
	// Current Local Time
	now := time.Now()
	fmt.Println("Time Now:", now)

	// Specific time
	specificTime := time.Date(2025, time.April, 20, 18, 30, 0, 0, time.Local)
	fmt.Println("Specific Time:", specificTime)

	// Parse time // Time Template - Mon Jan 2 15:04:05 MST 2006
	parsedTime, _ := time.Parse("2006-01-02", "2020-05-01")
	parsedTime1, _ := time.Parse("06-01-02", "20-05-01")
	fmt.Println(parsedTime)
	fmt.Println(parsedTime1)

	// Formatting Time
	t := time.Now()
	fmt.Println("Formatted Time:", t.Format("02-01-06 15-04"))

	oneDayLater := t.Add(time.Hour * 24)
	fmt.Println("One Day Later:", oneDayLater.Format("02-01-06"))

	// A 30 seconds for
	/*
		duration := time.Now().Add(time.Second * 30).Format("15-04-05")

		for {
			if duration == time.Now().Format("15-04-05") {
				fmt.Println("Han pasado 30 segundos")
				return
			}
		}
	*/

	loc, _ := time.LoadLocation("America/New_York")
	t = time.Date(2024, time.July, 8, 16, 20, 0, 0, time.UTC)

	// Convert loc to specific time zone
	tLocal := t.In(loc)

	// Perform rounding
	roundedTime := t.Round(time.Hour)
	roundedTimeLocal := roundedTime.In(loc)

	fmt.Println("Original time UTC:", t)
	fmt.Println("Local time EDT:", tLocal)
	fmt.Println("Rounded time UTC:", roundedTime)
	fmt.Println("Local Rounded time EDT:", roundedTimeLocal)

	fmt.Println("Truncated Time:", time.Now().Truncate(time.Hour)) // Round Down

	newYorkLocation, _ := time.LoadLocation("America/New_York")

	// Convert Time to Location

	timeInNY := time.Now().In(newYorkLocation)
	fmt.Println("Time in New York:", timeInNY)

}
