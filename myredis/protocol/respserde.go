package protocol

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type String string
type Array []interface{}
type BulkString struct {
	length int
	content string
}
type Error struct {
	message string
}

const (
	NULL_BULK_STRING = iota 
	NULL_ARRAY
)

const SEPARATOR string = "\r\n"



func ExtractFirstPart(message string) (string, string) {
	firstPart, rest, _ := strings.Cut(message, SEPARATOR)
	return firstPart, rest
}

func ParseInteger(message string) int {
	number, err := strconv.Atoi(message)
	if err != nil {
		fmt.Println("TODO: Invalid integer")
	}
	return number
}


func deserializeArrayValues(command string)(Array, error){
	//Clients send commands to a Redis server as an array of bulk strings.
	arrayInput, rest, found := strings.Cut(command, SEPARATOR)
	if !found || arrayInput[0] != '*' {
		return nil, errors.New("Invalid Command from the client")
	}
	inputLength, err := strconv.Atoi(arrayInput[1:])
	if err != nil {
		return nil, err
	}
	var bulkstrings []interface{}
	for i := 0; i < inputLength; i++ {
		parsedValue, remainder := Deserialize(rest)
		bulkstrings = append(bulkstrings, parsedValue)
		rest = remainder
	}
	return bulkstrings, nil 
}


func Deserialize(message string) (interface{}, string) {
	
	//check first byte to determine type
	firstByte := message[0]

	restOfMessage := message[1:]
	left, right := ExtractFirstPart(restOfMessage)
	switch firstByte {
	case '+':
		return left, right 
	case '-':
		return Error{message: left}, right 
	case ':':
		return ParseInteger(left), right 
	case '$':
		bslength := ParseInteger(left)
		if bslength != -1 {
			left, right := ExtractFirstPart(right)
			return BulkString{length: bslength, content: left}, right
		} else {
			return NULL_BULK_STRING, ""
		}			
	case '*':
		array, _ := deserializeArrayValues(restOfMessage)
	     return array, ""
	default:
		return nil, ""
	}
}