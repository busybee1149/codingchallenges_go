package tree

type HuffmanNode interface {
	IsLeaf() bool
	Weight() int
}

type HuffmanLeafNode struct {
	element rune
	weight int
}

func (h HuffmanLeafNode) IsLeaf() bool {
	return true
}

func (h HuffmanLeafNode) Weight() int {
	return h.weight
}
func (h HuffmanLeafNode) Character() rune {
	return h.element
}

type HuffmanInternalNode struct {
	left, right HuffmanNode
	weight int
}

func (h HuffmanInternalNode) IsLeaf() bool {
	return false
}

func (h HuffmanInternalNode) Weight() int {
	return h.weight
}

func (h HuffmanInternalNode) Left() HuffmanNode {
	return h.left
}

func (h HuffmanInternalNode) Right() HuffmanNode {
	return h.right
}

type HuffmanTree struct {
	root HuffmanNode
}

func (h HuffmanTree) Root() HuffmanNode {
	return h.root
}
func (h HuffmanTree) Weight() int {
	return h.root.Weight()
}


// func BuildTree() HuffmanTree {
// }
