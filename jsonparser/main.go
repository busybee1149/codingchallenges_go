package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Arguments received", os.Args[1:])

	if len(os.Args) != 2 {
		fmt.Println("Expected only one argument")
		os.Exit(1)
	}

	argString := os.Args[1]
	filedata, err := os.ReadFile(argString)

	if err != nil {
		fmt.Println("Invalid input")
		os.Exit(1)
	}
	fileContent := string(filedata)
	validJson := json.Valid([]byte(fileContent))

	if !validJson {
		fmt.Println("Invalid JSON")
		os.Exit(1)
	}
	fmt.Println("Valid JSON")
	os.Exit(0)
}