package main

type bitSet struct {
	content []bool
}

type BitSet interface {
	Add(value bool)
	Len() int
	String() string
	Copy() BitSet
}

func NewBitSet() BitSet {
	return &bitSet{
		content: make([]bool, 0, 8),
	}

}

func (bs *bitSet) Copy() BitSet {
	newContent := make([]bool, 0, 8)
	newContent = append(newContent, bs.content...)
	return &bitSet{
		content: newContent,
	}
}

func (bs *bitSet) Add(value bool) {
	bs.content = append(bs.content, value)
}

func (bs *bitSet) Len() int {
	return len(bs.content)
}

func (bs *bitSet) String() string {
	output := ""
	for _, value := range bs.content {
		if value {
			output = output + "1"
		} else {
			output = output + "0"
		}
	}
	return output
}
