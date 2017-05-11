package main

import "fmt"

func main() {
	dayMonths := map[string]int{
		"Jan": 31,
		"Feb": 28,
		"Mar": 31,
		"Apr": 30,
		"May": 31,
		"Jun": 30,
		"Jul": 31,
		"Aug": 31,
		"Sep": 30,
		"Oct": 31,
		"Nov": 30,
		"Dec": 31,
	}

	// dayMonths := make(map[string]int)
	// dayMonths["Jan"] = 31
	// dayMonths["Feb"] = 28
	// dayMonths["Mar"] = 31
	// dayMonths["Apr"] = 30
	// dayMonths["May"] = 31
	// dayMonths["Jun"] = 30
	// dayMonths["Jul"] = 31
	// dayMonths["Aug"] = 31
	// dayMonths["Sep"] = 30
	// dayMonths["Oct"] = 31
	// dayMonths["Nov"] = 30
	// dayMonths["Dec"] = 31

	// days, ok := dayMonths["Jun"]
	// if !ok {
	// 	fmt.Printf("Can't get days for Jun")
	// } else {
	// 	fmt.Printf("Days in June %d\n", days)
	// }

	has28 := 1

	delete(dayMonths, "Feb")
	delete(dayMonths, "February")

	for _, days := range dayMonths {
		if days == 28 {
			has28++
		}
	}
	fmt.Printf("%d months 28 days\n", has28)
}
