package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"errors"
	"fmt"
	"huffman/tree"
	"huffman/utils"
	"io"
	"os"
	_ "strings"
	"github.com/golang-collections/go-datastructures/bitarray"
)

	
func check(e error) {
	if e != nil {
	   os.Exit(-1)
	}
}

func bitStringToArray(bitString string)bitarray.BitArray {
	stringLength := len(bitString)
	binaryString := bitarray.NewBitArray(uint64(stringLength))

	for index, character := range bitString {
		bitIndex := stringLength - index - 1
		if (character == '1') {
			binaryString.SetBit(uint64(bitIndex))
		}
	}
	return binaryString
}
// func stringToBits(bitString string)[]byte {
//     lenB := len(bitString) / 8 + 1
//     bs := make([]byte,lenB)
//     count,i := 0,0
//     var now byte
//     for _,v := range bitString {
//         if count == 8 {
//             bs[i] = now
//             i++
//             now, count = 0,0
//         }
//         now = now << 1 + byte(v - '0')
//         count++
//     }
//     if count != 0 {
//         bs[i] = now << (8-byte(count))
//         i++
//     }

//     bs = bs[:i:i]
//     return bs
// }


func main() {
	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Println("Expected 2 arguments, input and output file")
		os.Exit(-1)
	}

	filename := arguments[1]
	file, err := os.ReadFile(filename)
	check(err)

	characterCounts := utils.GetCharacterFrequency(string(file))

	fmt.Println("Output file:", arguments[2])

	outputFile, err := os.Create(arguments[2])
	check(err)
	defer outputFile.Close()

	//Write Header, with character frequencies
	fileToWrite := bufio.NewWriter(outputFile)
	check(err)
	
	b := new(bytes.Buffer)
    	e := gob.NewEncoder(b)
	err = e.Encode(characterCounts)
	check(err)


	headerBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(headerBytes, uint32(b.Len()))
	fileToWrite.Write(headerBytes)
	fileToWrite.Write(b.Bytes())


	huffmanCodes := tree.BuildHuffmanCodingTree(characterCounts)
	
	for character := range file {
		characterCode := huffmanCodes[rune(character)]
		bytesNeeded := len(characterCode) / 8 + 1
		bytesToWrite := make([]byte, bytesNeeded)
		bitStringArray := bitStringToArray(characterCode)
		binary.Encode(bytesToWrite, binary.LittleEndian, bitStringArray)
		fileToWrite.Write(bytesToWrite)
	}
	fileToWrite.Flush()

	outputRead, err := os.Open(outputFile.Name())
	check(err)

	fileStats, err := outputRead.Stat()
	check(err)
	fmt.Println("File Size", fileStats.Size())

	headerLengthRBytes := make([]byte, 4)
	headerLengthR, err := outputRead.Read(headerLengthRBytes)
	check(err)
	headerLength := binary.LittleEndian.Uint32(headerLengthRBytes)
	fmt.Printf("Header Length read from output: %d, %d\n", headerLength, headerLengthR)

	offset, err := outputRead.Seek(4, io.SeekStart)
	check(err)

	headerContentBytes := make([]byte, headerLength)
	headerContentBytesRead, err := outputRead.ReadAt(headerContentBytes, offset)
	check(err)
	fmt.Printf("Header Content Bytes Read: %d\n", headerContentBytesRead)

	bufferForDecode := bytes.NewBuffer(headerContentBytes)
	decoder := gob.NewDecoder(bufferForDecode)
	decodedMap := make(map[rune]int)
	decoder.Decode(&decodedMap)

	encodedTextOffset, err := outputRead.Seek((int64(headerContentBytesRead + 4)), io.SeekStart)
	check(err)
	oneByte := make([]byte, 1)
	for {
		_, err := outputRead.ReadAt(oneByte, encodedTextOffset)
		if err != nil && !errors.Is(err, io.EOF) {
			fmt.Println(err)
			break
		}
	 
		// process the one byte b[0]
		fmt.Printf("%b\n", oneByte)
		if err != nil {
			// end of file
			break
		}
		encodedTextOffset++
	}


	
	fmt.Println(decodedMap)
}