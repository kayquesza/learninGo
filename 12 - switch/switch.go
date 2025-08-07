package main

import "fmt"

func dayOfWeek(day int) string {
	switch day {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	case 7:
		return "Sunday"
	default:
		return "Invalid day"
	}
}

func dayOfWeekTwo(day int) string {

	var dayOfWeek string

	switch {
	case day == 1:
		dayOfWeek = "Monday"
		fallthrough // fallthrough joga para o próximo case
	case day == 2:
		dayOfWeek = "Tuesday"
		dayOfWeek = "Wednesday"
	case day == 3:
	case day == 4:
		dayOfWeek = "Thursday"
	case day == 5:
		dayOfWeek = "Friday"
	case day == 6:
		dayOfWeek = "Saturday"
	case day == 7:
		dayOfWeek = "Sunday"
	default:
		dayOfWeek = "Invalid day"
	}
	return dayOfWeek
}

func main() {
	fmt.Println("Implementing a Switch.")

	fmt.Println("Day 1:", dayOfWeek(3))
	fmt.Println("Day 2:", dayOfWeekTwo(5))
	fmt.Println("Day 3:", dayOfWeekTwo(1))
	fmt.Println("Day 4:", dayOfWeek(8))

}
