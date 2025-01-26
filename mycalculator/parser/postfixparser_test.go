package parser

import (
	_ "slices"
	"testing"
)


func TestInfixToPostfixSimpleInput(t *testing.T) {
	tokens := "(1 * 2) - (4 / 3)"
	expected := "12*43/-"
	result := InfixToPostfix(tokens)
	if result != expected {
		t.Fatalf("expected %v, actual %v", expected, result)
	}
}