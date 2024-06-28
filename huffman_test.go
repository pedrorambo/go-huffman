package main

import (
	"fmt"
	"testing"
)

func TestConvertHuffmanTreeToString(t *testing.T) {
	left := *&huffmanNode{
		value:     byte('a'),
		frequency: 10,
		left:      nil,
		right:     nil,
	}
	rightRight := *&huffmanNode{
		value:     byte('b'),
		frequency: 20,
		left:      nil,
		right:     nil,
	}
	rightLeft := *&huffmanNode{
		value:     byte('c'),
		frequency: 30,
		left:      nil,
		right:     nil,
	}
	right := huffmanNode{
		value:     0,
		frequency: 50,
		left:      &rightLeft,
		right:     &rightRight,
	}
	root := huffmanNode{
		value:     0,
		frequency: 60,
		left:      &left,
		right:     &right,
	}
	formatted := HuffmanTreeToString(&root)
	if formatted != "(a)((c)(b))" {
		t.Fatal("Huffman tree to string doesnt match the expeted format")
	}
}

func TestCreateTheSimplestTree(t *testing.T) {
	var frequencies []ByteFrequency
	frequencies = append(frequencies, ByteFrequency{
		value:     byte('a'),
		frequency: 5,
	})
	frequencies = append(frequencies, ByteFrequency{
		value:     byte('b'),
		frequency: 9,
	})
	tree := CreateTreeForFrequencies(frequencies)
	format := HuffmanTreeToString(tree)
	if format != "(a)(b)" {
		t.Fatal("Tree doesnt match the expected format")
	}
}

func TestCreateTreeWith3Items(t *testing.T) {
	var frequencies []ByteFrequency
	frequencies = append(frequencies, ByteFrequency{
		value:     byte('a'),
		frequency: 5,
	})
	frequencies = append(frequencies, ByteFrequency{
		value:     byte('b'),
		frequency: 9,
	})
	frequencies = append(frequencies, ByteFrequency{
		value:     byte('c'),
		frequency: 10,
	})
	tree := CreateTreeForFrequencies(frequencies)
	format := HuffmanTreeToString(tree)
	if format != "(c)((a)(b))" {
		t.Fatal("Tree doesnt match the expected format")
	}
}

func TestCreateTreeWith4Items(t *testing.T) {
	var frequencies []ByteFrequency
	frequencies = append(frequencies, ByteFrequency{
		value:     byte('a'),
		frequency: 5,
	})
	frequencies = append(frequencies, ByteFrequency{
		value:     byte('b'),
		frequency: 9,
	})
	frequencies = append(frequencies, ByteFrequency{
		value:     byte('c'),
		frequency: 10,
	})
	frequencies = append(frequencies, ByteFrequency{
		value:     byte('d'),
		frequency: 13,
	})
	tree := CreateTreeForFrequencies(frequencies)
	format := HuffmanTreeToString(tree)
	if format != "((a)(b))((c)(d))" {
		t.Fatal("Tree doesnt match the expected format")
	}
}

func TestCreateTreeForFrequencies(t *testing.T) {
	var frequencies []ByteFrequency
	frequencies = append(frequencies, ByteFrequency{
		value:     byte('a'),
		frequency: 5,
	})
	frequencies = append(frequencies, ByteFrequency{
		value:     byte('b'),
		frequency: 9,
	})
	frequencies = append(frequencies, ByteFrequency{
		value:     byte('c'),
		frequency: 12,
	})
	frequencies = append(frequencies, ByteFrequency{
		value:     byte('d'),
		frequency: 13,
	})
	frequencies = append(frequencies, ByteFrequency{
		value:     byte('e'),
		frequency: 16,
	})
	tree := CreateTreeForFrequencies(frequencies)
	format := HuffmanTreeToString(tree)
	fmt.Println(format)
	if format != "((c)(d))(((e)(a))(b))" {
		t.Fatal("Tree doesnt match the expected format")
	}
}
