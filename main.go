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
	/* SLICE USAGE */
	// newSlice := make([]float32, 5, 10) // 5 elemanli ve max 10 kapasiteli slice tanimi
	// newSlice[0] = 1.32
	// newSlice = append(newSlice, math.Pi, 2.277) // Slice'in sonuna eleman ekledi
	// fmt.Println(newSlice)

	// newSlice1 := []int{3, 5, 1, 7, 8} // 5 elemanli ve kapasiteli bir slice
	// newSlice2 := make([]int, 3)       // 3 elemanli ve kapasiteli bir slice
	// copy(newSlice2, newSlice1)
	// fmt.Println(newSlice1)
	// fmt.Println(newSlice2)

	/* MAP USAGE */
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
	//
	// for month, day := range dayMonths {
	// 	//fmt.Printf("%d\n", days)
	// 	if month == "Feb" {
	// 		fmt.Printf("%d days in February\n", day)
	// 	}
	// }

	/* CUSTOM ERROR USAGE */
	// if err := printer(""); err != nil {
	// 	if err == errorEmptyString {
	// 		fmt.Printf("You tried to print an empty string!\n")
	// 	} else {
	// 		fmt.Printf("printer failed: %s\n", err)
	// 	}
	// 	os.Exit(1)
	// }
}
