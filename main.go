package main

import (
	"errors"
	"fmt"
)

var (
	errorEmptyString = errors.New("Unwilling to print an emty string")
)

func printer(msg string) error {
	if msg == "" {
		return errorEmptyString
	}
	_, err := fmt.Printf("%s\n", msg)
	return err
}

func main() {
	dayMonths := make(map[string]int)
	dayMonths["Jan"] = 31
	dayMonths["Feb"] = 28
	dayMonths["Mar"] = 31
	dayMonths["Apr"] = 30
	dayMonths["May"] = 31
	dayMonths["Jun"] = 30
	dayMonths["Jul"] = 31
	dayMonths["Aug"] = 31
	dayMonths["Sep"] = 30
	dayMonths["Oct"] = 31
	dayMonths["Nov"] = 30
	dayMonths["Dec"] = 31

	for month, day := range dayMonths {
		//fmt.Printf("%d\n", days)
		if month == "Feb" {
			fmt.Printf("%d days in February\n", day)
		}
	}

	// if err := printer(""); err != nil {
	// 	if err == errorEmptyString {
	// 		fmt.Printf("You tried to print an empty string!\n")
	// 	} else {
	// 		fmt.Printf("printer failed: %s\n", err)
	// 	}
	// 	os.Exit(1)
	// }
}
