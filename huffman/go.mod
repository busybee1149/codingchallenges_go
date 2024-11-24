module huffman

go 1.23.1

replace huffman/utils => ./utils

require (
	huffman/tree v0.0.0-00010101000000-000000000000
	huffman/utils v0.0.0-00010101000000-000000000000
)

require (
	github.com/Workiva/go-datastructures v1.1.5 // indirect
	github.com/golang-collections/go-datastructures v0.0.0-20150211160725-59788d5eb259 // indirect
)

replace huffman/tree => ./tree
