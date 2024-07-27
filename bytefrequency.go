package main

import (
	"bufio"
	"errors"
	"io"
	"sort"
)

type ByteFrequency struct {
	value     byte
	frequency int
	node      *huffmanNode // Legacy. Must create another structure that has the frequencies and the node
}

func GetByteFrequency(br *bufio.Reader) []ByteFrequency {
	var foundBytes map[byte]int = make(map[byte]int)

	for {
		b, err := br.ReadByte()
		if err != nil && !errors.Is(err, io.EOF) {
			panic(err)
		}
		if err != nil {
			break
		}
		existingCount := foundBytes[b]
		foundBytes[b] = existingCount + 1
	}

	var frequencies []ByteFrequency
	for k, v := range foundBytes {
		frequencies = append(frequencies, ByteFrequency{
			value:     k,
			frequency: v,
		})
	}

	// When running the automated tests, for some reason, sometimes the created slice
	// doesnt keep the order of the appended elements. To keep the tests consistent,
	// we must sort the result.
	sort.Slice(frequencies, func(a, b int) bool {
		return frequencies[a].frequency > frequencies[b].frequency
	})

	return frequencies
}
