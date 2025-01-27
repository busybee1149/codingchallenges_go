package parser

import "unicode"

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

func InfixToPostfix(tokens string) string {
	var output []rune
	stack := newStack[rune]()

	tokensLength := len(tokens)

	for index := 0; index < tokensLength; index++ {
		charAtIndex := tokens[index]
		if unicode.IsDigit(rune(charAtIndex)) {
			output = append(output, rune(charAtIndex))
		} else if charAtIndex == '(' {
			stack.Push(rune(charAtIndex))
		} else if charAtIndex == ')' {
			topOfStack := stack.Pop()
			for ok := true; ok; ok = (topOfStack != '(') {
				output = append(output, topOfStack)
				topOfStack = stack.Pop()
			}
		} else {
			characterPriority, ok := operatorPriority[rune(charAtIndex)]
			if ok {
				topOfStack := stack.Top()
				topOfStackPriority, ok := operatorPriority[topOfStack]
				if ok && topOfStackPriority <= characterPriority {
					output = append(output, topOfStack)
					stack.Pop()
				} else {
					stack.Push(rune(charAtIndex))
				}
			}
		}
	}

	for stack.Length() > 0 {
		lastCharacter := stack.Pop()
		output = append(output, lastCharacter)
	}
	return string(output)
}
