package protocol

import (
	_ "errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	SIMPLE_STRING = '+'
	BULK_STRING   = '$'
	INTEGER       = ':'
	ARRAY         = '*'
	ERROR_STRING  = '-'
)

const SEPARATOR string = "\r\n"

var (
	NULL_BULK_STRING = BulkString{length: 0, content: "-1", prefix: BULK_STRING}
)

type Serializable interface {
	Serialize() string
}
type Integer struct {
	value int
	prefix rune
}
func NewInteger(number int) Integer {
	return Integer{prefix: INTEGER, value: number}
}
type String struct {
	value string
	prefix rune
}

func NewString(message string) String {
	return String{ value: message, prefix: SIMPLE_STRING }
}
type Array struct {
	Elements []interface{}
	length int
	prefix rune 
}

func NewArray(arguments ...interface{}) Array{
	var contents []interface{}
	contents = append(contents, arguments...)
	return Array{
		Elements: contents,
		length: len(contents),
		prefix: ARRAY,
	}
}

type BulkString struct {
	length int
	content string
	prefix rune
}

func NewBulkString(message string) BulkString {
	return BulkString{ length: len(message), content: message, prefix: BULK_STRING}
}

func (bulkString *BulkString) ContentString() string {
	return bulkString.content
}

type ErrorString struct {
	message string
	prefix rune
}

func NewError(message string) ErrorString {
	return ErrorString{ prefix: ERROR_STRING, message: message}
}

func (str String) Serialize() string {
	return fmt.Sprintf("%c%s%s", str.prefix, str.value, SEPARATOR)
}

func (err ErrorString) Serialize() string {
	return fmt.Sprintf("%c%s%s", err.prefix, err.message, SEPARATOR)
}
func (integer Integer) Serialize() string {
	return fmt.Sprintf("%c%s%s", integer.prefix, strconv.Itoa(integer.value), SEPARATOR)
}

func (bulkstring BulkString) Serialize() string {
	if bulkstring.length > 0 { 
		return fmt.Sprintf("%c%s%s%s%s", bulkstring.prefix, strconv.Itoa(bulkstring.length), SEPARATOR, bulkstring.content, SEPARATOR)
	} else {
		return fmt.Sprintf("%c%s%s", bulkstring.prefix, bulkstring.content, SEPARATOR)
	}
}

func (array Array) Serialize() string {
	var stringbuilder strings.Builder
	stringbuilder.WriteRune(array.prefix)
	stringbuilder.WriteString(strconv.Itoa(array.length))
	stringbuilder.WriteString(SEPARATOR)
	for i := 0; i < len(array.Elements); i++ {
		element := array.Elements[i]
		switch element.(type) {
		case String:
			element := element.(String)
			stringbuilder.WriteString(element.Serialize())
		case Integer:
			element := element.(Integer)
			stringbuilder.WriteString(element.Serialize())
		case ErrorString:
			element := element.(ErrorString)
			stringbuilder.WriteString(element.Serialize())
		case BulkString:
			element := element.(BulkString)
			stringbuilder.WriteString(element.Serialize())
		case Array:
			element := element.(Array)
			stringbuilder.WriteString(element.Serialize())
		}
	}
	return stringbuilder.String()
}




func DeserializeInteger(str string) (Integer, error) {
	number, err := strconv.Atoi(str[1:])
	if str[0] != INTEGER || err != nil {
		return Integer{}, fmt.Errorf("Deserialization failed for %s", str)
	} else {
		return NewInteger(number), nil
	}
}

func DeserializeString(str string) (String, error) {
	if str[0] != SIMPLE_STRING {
		return NewString(""), fmt.Errorf("Deserialization failed for %s", str)
	} else {
		return NewString(str[1:]), nil
	}
}

func DeserializeBulkString(str string) (BulkString, error) {
	if str[0] != BULK_STRING {
		return NewBulkString(""), fmt.Errorf("Deserialization failed for %s", str)
	} else {

		return NewBulkString(str[1:]), nil
	}
}

func DeserializeError(str string) (ErrorString, error) {
	if str[0] != ERROR_STRING {
		return NewError(""), fmt.Errorf("Deserialization failed for %s", str)
	} else {
		return NewError(str[1:]), nil
	}
}

func deserializeBasicTypes(message string) (interface{}, error) {
	var deserialized interface{}
	var err error

	switch message[0] {
	case SIMPLE_STRING:
		deserialized, err = DeserializeString(message)
	case ERROR_STRING:
		deserialized, err = DeserializeError(message)
	case INTEGER:
		deserialized, err = DeserializeInteger(message)
	default:
		deserialized, err = nil, fmt.Errorf("Deserialization failed")
	}
	return deserialized, err
}

func deserializeArray(str string) (Array, error) {
	if str[0] != ARRAY {
		return NewArray(), fmt.Errorf("Deserialization failed for %s", str)
	} else {
		first, second := extractTwoParts(str)
		arrayLength, numberConvErr := strconv.Atoi(first[1:])

		if numberConvErr != nil {
			return NewArray(), fmt.Errorf("Deserialization failed for %s", str)
		}

		var elements []interface{}
		var deserialized interface{}
		var err error

		for i := 0; i < arrayLength; i++ {
			switch second[0] {
			case SIMPLE_STRING, ERROR_STRING, INTEGER:
				part1, part2 := extractTwoParts(second)
				deserialized, err = deserializeBasicTypes(part1)
				second = part2
			case BULK_STRING:
				parts := strings.SplitAfterN(second[1:], SEPARATOR, 2)
				bslengthstring, _ := extractTwoParts(parts[0])
				var bslength int 
				bslength, err = strconv.Atoi(bslengthstring)
				bscontent, rest := extractTwoParts(parts[1])
				if bslength != len(bscontent) {
					err = fmt.Errorf("Deserialization failed for %s", str)
				} else {
					deserialized = NewBulkString(bscontent)
				}
				second = rest

			default:
				return NewArray(), fmt.Errorf("Deserialization failed for %s", str)
			}
			if err != nil {
				return NewArray(), fmt.Errorf("Deserialization failed for %s", str)
			} else {
				//fmt.Println("Deserialized element", deserialized, i)
				elements = append(elements, deserialized)
			}
		}
		return  NewArray(elements...), nil
	} 
}

func Deserialize(message string) (Array, error) {
	//fmt.Println("trying to deserialize", message)
	return deserializeArray(message)
}


func extractTwoParts(message string) (string, string) {
	firstPart, rest, _ := strings.Cut(message, SEPARATOR)
	return firstPart, rest
}