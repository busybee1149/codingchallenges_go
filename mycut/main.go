package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type cutParameters struct {
	Delimiter string
	Columns []int
	FileName string
}

func checkErrorAndExit(err error, errorMessage string) {
	if err != nil {
		fmt.Println(errorMessage)
		os.Exit(-1)
	}
}

func extractColumns(arg string) []int {
	columns := []int{} 
	columnString := arg[2:]
	var columnStrings []string
	if strings.Contains(columnString, " ") {
		columnStrings = strings.Split(columnString, " ")
	} else if strings.Contains(columnString, ",") {
		columnStrings = strings.Split(columnString, ",")
	} else {
		number, err := strconv.Atoi(columnString)
		checkErrorAndExit(err, "Invalid Column value")
		columns = append(columns, number)
	}
	for _, value := range columnStrings {
		number, err := strconv.Atoi(value)
		checkErrorAndExit(err, "Invalid Column value")
		columns = append(columns, number)
	}
	return columns
}

func extractCutParameters(args []string) cutParameters {
	cutParameters := cutParameters{
		Delimiter: "\t",
	}
	for _, arg := range args {
		switch {
		case strings.HasPrefix(arg, "-f"):
			cutParameters.Columns = extractColumns(arg)
		case strings.HasPrefix(arg, "-d"):
			cutParameters.Delimiter = arg[2:]
		default:
			cutParameters.FileName = arg
		}
	}
	return cutParameters
}

func main() {
	argsAfterProgramName := os.Args[1:]
	cutParameters := extractCutParameters(argsAfterProgramName)

	var inputReader *bufio.Reader
	if len(cutParameters.FileName) != 0 && cutParameters.FileName != "-" {
		givenFile, err := os.Open(cutParameters.FileName)
		checkErrorAndExit(err, "Invalid file")
		inputReader = bufio.NewReader(givenFile)
	} else {
		inputReader = bufio.NewReader(os.Stdin)
	}

	for {
		line, err := inputReader.ReadString('\n')
		lineColumns := strings.Split(line, cutParameters.Delimiter)
		columnsSelected := make([]string, len(cutParameters.Columns))
		for index, column := range cutParameters.Columns {
			columnsSelected[index] = lineColumns[column]
		}
		fmt.Println(strings.Join(columnsSelected, cutParameters.Delimiter))
		if err != nil {
			if err == io.EOF {
				break
			}
			os.Exit(-1)
		}
	}

}