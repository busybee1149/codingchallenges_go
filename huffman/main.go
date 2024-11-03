package main

import (
	"fmt"
	"os"
	"huffman/utils"
	"huffman/tree"
)

	
func check(e error) {
	if e != nil {
	   os.Exit(-1)
	}
 }

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Expected only one argument for filename")
		os.Exit(-1)
	}

	filename := arguments[1]
	file, err := os.ReadFile(filename)
	check(err)
	
	characterCounts := utils.GetCharacterFrequency(string(file))
	huffmanTree := tree.BuildHuffmanTree(characterCounts)
	fmt.Println(huffmanTree)
}