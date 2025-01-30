package parser

import (
	_ "slices"
	"testing"
)


func TestInfixToPostfixSimpleInput(t *testing.T) {
	tokens := "(10 * 23) - (41 / 3)"
	expected := "10,23,*,41,3,/,-,"
	result := InfixToPostfix(tokens)
	if result != expected {
		t.Fatalf("expected %v, actual %v", expected, result)
	}
}