package main

import (
	"fmt"
	"os"
	"strings"
)


func sanitizeInput(input string) string {
	var sanitizedInput string 
	sanitizedInput = strings.TrimPrefix(input, "'")
	sanitizedInput = strings.TrimSuffix(sanitizedInput, "'")
	return strings.Trim(sanitizedInput, " ")
}

func main() {
	arguments := os.Args[1:]

	if len(arguments) != 1 {
		fmt.Println("Incorrect usage")
		os.Exit(-1)
	}

	sanitizedInput := sanitizeInput(arguments[0])
	fmt.Println(sanitizedInput)

}