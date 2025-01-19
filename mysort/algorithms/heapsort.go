package algorithms

import "container/heap"


type StringHeap []string 

func (h StringHeap) Len() int { return len(h) }
func (h StringHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h StringHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *StringHeap) Push(x any) {
	(*h) = append((*h), x.(string))
}

func (h *StringHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func HeapSort(words []string) []string {
	linesHeap := &StringHeap {}
	heap.Init(linesHeap)
	for _, word := range words {
		heap.Push(linesHeap, word)
	}
	var sortedWords []string
	for linesHeap.Len() > 0 {
		sorted := heap.Pop(linesHeap)
		sortedWords = append(sortedWords, sorted.(string))
	}
	return sortedWords
}