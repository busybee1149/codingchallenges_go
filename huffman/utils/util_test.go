package utils

import (
	"maps"
	"testing"
)

func TestGetCharacterFrequencyHappyCase(t *testing.T) {
	testString := "12abab12"
	result := GetCharacterFrequency(testString)
	expected := map[rune]int{'1':2, '2': 2,'a':2, 'b':2 }
	if !maps.Equal(result, expected) { 
		t.Fatalf("result=%v does not match expected=%v", result, expected)
	}
}

func TestGetCharacterFrequencyEmptyString(t *testing.T) {
	testString := ""
	result := GetCharacterFrequency(testString)
	expected := map[rune]int{}
	if !maps.Equal(result, expected) { 
		t.Fatalf("result=%v does not match expected=%v", result, expected)
	}
}