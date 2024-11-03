package tree

type PriorityQueue []*HuffmanTree

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Weight() > pq[j].Weight()
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*HuffmanTree)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	*pq = old[0 : n-1]
	return item
}

func BuildHuffmanTree(characterCounts map[rune]int) []*HuffmanTree {
	huffmanTree := []*HuffmanTree{}
	for character,count := range characterCounts {
		huffmanTree = append(huffmanTree, &HuffmanTree{
			root: HuffmanLeafNode{
				element: character,
				weight: count,
			},
		})
	}
	return huffmanTree
}