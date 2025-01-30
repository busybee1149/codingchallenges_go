package parser

import (
	"strconv"
	"strings"
	"unicode"
)

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(t T) {
	s.items = append(s.items, t)
}

func (s *Stack[T]) Length() int {
	return len(s.items)
}

func (s *Stack[T]) Pop() T {
	stackLength := s.Length()
	var item T
	if stackLength > 0 {
		item = s.items[stackLength-1]
	}
	s.items = s.items[:stackLength-1]
	return item
}

func (s *Stack[T]) Top() T {
	stackLength := s.Length()
	var item T
	if stackLength > 0 {
		item = s.items[stackLength-1]
	}
	return item
}

func newStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

var operatorPriority = map[rune]int{
	'*': 1,
	'/': 1,
	'%': 1,
	'+': 2,
	'-': 2,
}

func nextNumber(expression string, startIndex int) (int, int) {
	for ; startIndex < len(expression) && !(expression[startIndex] >= '0' && expression[startIndex] <= '9'); startIndex++ {}

	x := 0
	for ; startIndex < len(expression) && (expression[startIndex] >= '0' && expression[startIndex] <= '9'); startIndex++ {
		x = x * 10 + int(expression[startIndex] - '0')
	}
	return x, startIndex
}

func InfixToPostfix(tokens string) string {
	var output strings.Builder
	stack := newStack[rune]()

	tokensLength := len(tokens)

	for index := 0; index < tokensLength;  {
		charAtIndex := tokens[index]
		if unicode.IsDigit(rune(charAtIndex)) {
			number, nextIndex := nextNumber(tokens, index)
			output.WriteString(strconv.Itoa(number))
			output.WriteRune(',')
			index = nextIndex
		} else {
			 if charAtIndex == '(' {
				stack.Push(rune(charAtIndex))
			 } else if charAtIndex == ')' {
				topOfStack := stack.Pop()
				for ok := true; ok; ok = (topOfStack != '(') {
					output.WriteRune(topOfStack)
					output.WriteRune(',')
					topOfStack = stack.Pop()
				}
			} else {
				characterPriority, ok := operatorPriority[rune(charAtIndex)]
				if ok {
					topOfStack := stack.Top()
					topOfStackPriority, ok := operatorPriority[topOfStack]
					if ok && topOfStackPriority <= characterPriority {
						output.WriteRune(topOfStack)
						output.WriteRune(',')
						stack.Pop()
					} else {
						stack.Push(rune(charAtIndex))
					}
				}
			}
			index++
		}
	}

	for stack.Length() > 0 {
		lastCharacter := stack.Pop()
		output.WriteRune(lastCharacter)
		output.WriteRune(',')
	}
	return output.String()
}
