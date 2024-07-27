package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
)

type huffmanNode struct {
	value     byte
	frequency int
	left      *huffmanNode
	right     *huffmanNode
}

type treeLookupItem struct {
	value  byte
	bitSet BitSet
}

func generateTreeLookup(node *huffmanNode, bs BitSet) []treeLookupItem {
	if node == nil {
		return make([]treeLookupItem, 0)
	}
	var items []treeLookupItem
	if node.left == nil && node.right == nil {
		items = append(items,
			treeLookupItem{
				value:  node.value,
				bitSet: bs,
			},
		)
		return items
	}
	leftBs := bs.Copy()
	leftBs.Add(false)
	rightBs := bs.Copy()
	rightBs.Add(true)
	items = append(items, generateTreeLookup(node.left, leftBs)...)
	items = append(items, generateTreeLookup(node.right, rightBs)...)
	return items
}

func getBitSetForLookupItem(lookup []treeLookupItem, value byte) BitSet {
	for _, item := range lookup {
		if item.value == value {
			return item.bitSet
		}
	}
	return nil
}

func HuffmanTreeToString(node *huffmanNode) string {
	if node == nil {
		return ""
	}
	if node.left == nil && node.right == nil {
		return string(node.value)
	}
	return "(" + HuffmanTreeToString(node.left) + ")" + "(" + HuffmanTreeToString(node.right) + ")"
}

func CreateTreeForFrequencies(frequencies []ByteFrequency) *huffmanNode {
	var lastNode *huffmanNode

	for {
		sort.Slice(frequencies, func(i, j int) bool {
			return frequencies[i].frequency > frequencies[j].frequency
		})

		if len(frequencies) == 1 {
			break
		}

		lastFrequency := frequencies[len(frequencies)-1]
		preLastFrequency := frequencies[len(frequencies)-2]

		if lastFrequency.node == nil && preLastFrequency.node == nil {
			totalFrequency := lastFrequency.frequency + preLastFrequency.frequency
			newNode := huffmanNode{
				value:     0,
				frequency: totalFrequency,
				left: &huffmanNode{
					value:     lastFrequency.value,
					frequency: lastFrequency.frequency,
				},
				right: &huffmanNode{
					value:     preLastFrequency.value,
					frequency: preLastFrequency.frequency,
				},
			}
			lastNode = &newNode
			frequencies = frequencies[:len(frequencies)-2]
			frequencies = append(frequencies, ByteFrequency{
				value:     0,
				frequency: totalFrequency,
				node:      &newNode,
			})
		} else {
			if lastFrequency.node == nil {
				totalFrequency := preLastFrequency.frequency + lastFrequency.frequency
				newNode := huffmanNode{
					value:     0,
					frequency: totalFrequency,
					left: &huffmanNode{
						value:     lastFrequency.value,
						frequency: lastFrequency.frequency,
					},
					right: preLastFrequency.node,
				}
				lastNode = &newNode
				frequencies = frequencies[:len(frequencies)-2]
				frequencies = append(frequencies, ByteFrequency{
					value:     0,
					frequency: totalFrequency,
					node:      &newNode,
				})
			} else if preLastFrequency.node == nil {
				totalFrequency := lastFrequency.frequency + preLastFrequency.frequency
				newNode := huffmanNode{
					value:     0,
					frequency: totalFrequency,
					left:      lastFrequency.node,
					right: &huffmanNode{
						value:     preLastFrequency.value,
						frequency: preLastFrequency.frequency,
					},
				}
				lastNode = &newNode
				frequencies = frequencies[:len(frequencies)-2]
				frequencies = append(frequencies, ByteFrequency{
					value:     0,
					frequency: totalFrequency,
					node:      &newNode,
				})
			} else {
				totalFrequency := lastFrequency.frequency + preLastFrequency.frequency
				newNode := huffmanNode{
					value:     0,
					frequency: totalFrequency,
					left:      lastFrequency.node,
					right:     preLastFrequency.node,
				}
				lastNode = &newNode
				frequencies = frequencies[:len(frequencies)-2]
				frequencies = append(frequencies, ByteFrequency{
					value:     0,
					frequency: totalFrequency,
					node:      &newNode,
				})
			}
		}
	}

	return lastNode
}

func main() {
	f, err := os.Open("./file.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	var frequencies []ByteFrequency = GetByteFrequency(reader)

	lastNode := CreateTreeForFrequencies(frequencies)

	lookupTable := generateTreeLookup(lastNode, NewBitSet())

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

		foundBitSet := getBitSetForLookupItem(lookupTable, b)

		if foundBitSet == nil {
			panic("Error")
		}

		size := uint64(foundBitSet.Len())
		newTotalSizeInBits += size
	}

	fmt.Println(newTotalSizeInBits, "bits")
}
