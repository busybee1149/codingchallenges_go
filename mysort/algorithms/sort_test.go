package algorithms

import (
	"slices"
	"testing"
)


func TestQuickSortWithValidInput(t *testing.T) {
	words := []string { "w1", "gold", "for", "held"}
	expected := []string { "w1", "gold", "for", "held"}
	slices.Sort(expected)
	sorted := QuickSort(words) 
	if !slices.Equal(sorted, expected) {
		t.Fatalf("sorting failed for valid input %v", words)
	}
}