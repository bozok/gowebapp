package main

import (
	"errors"
	"fmt"
	"os"
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
	if err := printer(""); err != nil {
		if err == errorEmptyString {
			fmt.Printf("You tried to print an empty string!\n")
		} else {
			fmt.Printf("printer failed: %s\n", err)
		}
		os.Exit(1)
	}
}
