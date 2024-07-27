package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("./file.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	var frequencies []ByteFrequency = GetByteFrequency(reader)

	lastNode := CreateTreeForFrequencies(frequencies)

	lookupTable := generateTreeLookup(lastNode, NewBitSet())

	var originalTotalSizeInBits uint64 = 0
	var newTotalSizeInBits uint64 = 0

	newFile, err := os.Open("./file.txt")
	if err != nil {
		panic(err)
	}

	newReader := bufio.NewReader(newFile)

	for {
		b, err := newReader.ReadByte()
		if err != nil && !errors.Is(err, io.EOF) {
			panic(err)
		}
		if err != nil {
			break
		}
		originalTotalSizeInBits += 8

		foundBitSet := getBitSetForLookupItem(lookupTable, b)

		if foundBitSet == nil {
			panic("Error")
		}

		size := uint64(foundBitSet.Len())
		newTotalSizeInBits += size
	}

	changeInSize := originalTotalSizeInBits - newTotalSizeInBits
	newSizePercentage := (float64(changeInSize) / float64(originalTotalSizeInBits)) * 100
	fmt.Printf("Original size: %v bits\n", originalTotalSizeInBits)
	fmt.Printf("Deflated size: %v bits (deflated %.2f%%)\n", newTotalSizeInBits, newSizePercentage)
}
