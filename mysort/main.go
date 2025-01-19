package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"algorithms"
)

type SortParameters struct {
	enforceUniqueness bool
	fileName string
	algorithm string
}

func SortFileContent(sortParameters SortParameters) ([]string, error) {
	file, err := os.Open(sortParameters.fileName)
	if err != nil {
		fmt.Println("Error reading the file")
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var givenWords []string
	valuesMap := make(map[string]bool)
	for scanner.Scan() {
		word := scanner.Text()
		_, ok := valuesMap[word]
		if !ok {
			valuesMap[word] = true
			givenWords = append(givenWords, word)
		} else if !sortParameters.enforceUniqueness {
			givenWords = append(givenWords, word)
		}
	}
	var sortedWords []string
	switch sortParameters.algorithm {
	case "heapsort":
		sortedWords = algorithms.HeapSort(givenWords)
	case "quicksort":
		sortedWords = algorithms.QuickSort(givenWords)
	default:
		sortedWords = algorithms.QuickSort(givenWords)
	}
	return sortedWords, nil
}


func ReadFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error reading the file")
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func extractCutParameters(args []string) SortParameters {
	sortParameters := SortParameters{
		enforceUniqueness: false,
		algorithm: "quicksort",
	}
	for _, arg := range args {
		switch {
		case strings.HasPrefix(arg, "-u"):
			sortParameters.enforceUniqueness = true
		default:
			sortParameters.fileName = arg
		}
	}
	return sortParameters
}

func main() {
	cliArguments := os.Args[1:]
	if len(cliArguments) < 1 {
		fmt.Println("Insufficient arguments")
		os.Exit(-1)
	}
	sortParameters := extractCutParameters(cliArguments)

	fileContentByLines, err := SortFileContent(sortParameters)
	if err != nil {
		os.Exit(-1)
	}

	output := bufio.NewWriter(os.Stdout)
	for _, line := range fileContentByLines {
	 	output.WriteString(fmt.Sprintln(line))
	}
	output.Flush()
}