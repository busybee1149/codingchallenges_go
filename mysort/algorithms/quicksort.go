package algorithms

import (
	"fmt"

	"golang.org/x/exp/constraints"
)


func partition[E constraints.Ordered](words []E, start, end int) int {
	if start > end {
		fmt.Println("Invalid input provided")
		return -1
	}
	pivot := words[end]
	j := start

	for i := start; i < end; i++ {
		if words[i] <= pivot {
			words[i], words[j] = words[j], words[i]
			j++
		}
	}
	words[j], words[end] = words[end], words[j]
	return j
}

func quicksort[E constraints.Ordered](words []E, start, end int){
	if start < end {
		pivotIndex := partition(words, start, end)
		quicksort(words, start, pivotIndex - 1)
		quicksort(words, pivotIndex + 1, end)
	}
}

func QuickSort[E constraints.Ordered](words []E) []E{
	quicksort(words, 0, len(words) - 1)
	return words
}