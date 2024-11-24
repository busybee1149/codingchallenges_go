package tree

import (
	"container/heap"
	"fmt"
	_ "strconv"
)

type PriorityQueue []*HuffmanTree

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Weight() < pq[j].Weight()
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(*HuffmanTree))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	*pq = old[0 : n-1]
	return item
}

func buildPriorityQueue(characterCounts map[rune]int) *PriorityQueue {
	huffmanTree := make(PriorityQueue, len(characterCounts))
	fmt.Println("Total Characters found: ", len(characterCounts))
	i := 0
	for character, count := range characterCounts {
		//fmt.Printf("Character = %c, weight =  %d \n", character, count)
		huffmanTree[i] = &HuffmanTree{
			root: HuffmanLeafNode{
				element: character,
				weight: count,
			},
		}
		i++
	}
	return &huffmanTree
}

func BuildHuffmanCodingTree(characterCounts map[rune]int)map[rune]string {
	priorityQueue := buildPriorityQueue(characterCounts)
	heap.Init(priorityQueue)

	for priorityQueue.Len() > 1 {
		first := heap.Pop(priorityQueue).(*HuffmanTree)
		second := heap.Pop(priorityQueue).(*HuffmanTree)
		new := &HuffmanTree{
			root: HuffmanInternalNode{
				left: first.Root(),
				right: second.Root(),
				weight: first.Weight() + second.Weight(),
			},
		}
		priorityQueue.Push(new)
	}
	root := heap.Pop(priorityQueue).(*HuffmanTree)

	characterCodes := traverseTree(root)
	return characterCodes
}

func preorderTraversal(root *HuffmanNode, codeArray *[]rune, characterCodes *map[rune]string) {
	if root == nil {
		return 
	}
	if (*root).IsLeaf() {
		leaf := (*root).(HuffmanLeafNode)
		(*characterCodes)[leaf.element] = string(*codeArray)
		//fmt.Printf("leaf node %s, %s \n", strconv.QuoteRune(leaf.element), string(*codeArray))
	} else {
		internalNode := (*root).(HuffmanInternalNode)
		leftArray := make([]rune, len(*codeArray) + 1)
		copy(leftArray, *codeArray)
		leftArray[len(*codeArray)] = '0'
		rightArray := make([]rune, len(*codeArray) + 1)
		copy(rightArray, *codeArray)
		rightArray[len(*codeArray)] = '1'
		preorderTraversal(&internalNode.left, &leftArray, characterCodes)
		preorderTraversal(&internalNode.right, &rightArray, characterCodes)
	}
}

func traverseTree(huffmanTree *HuffmanTree) map[rune]string {
	characterCodes := make(map[rune]string)
	codeArray := []rune{}
	preorderTraversal(&huffmanTree.root, &codeArray, &characterCodes)
	return characterCodes
}