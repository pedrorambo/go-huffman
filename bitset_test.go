package main

import (
	"testing"
)

func TestCreateEmpty(t *testing.T) {
	bs := NewBitSet()
	if bs.Len() != 0 {
		t.Fatalf("The BitSet just created isn't empty")
	}
}

func TestAddValuesToBitSet(t *testing.T) {
	bs := NewBitSet()
	bs.Add(true)
	bs.Add(true)
	bs.Add(false)
	if bs.Len() != 3 {
		t.Fatalf("BitSet is not of size 3")
	}
	if bs.String() != "110" {
		t.Fatalf("BitSet as string is not 110")
	}
}

func TestAlteringACopyMustNotAffectTheOriginalBitSet(t *testing.T) {
	bs := NewBitSet()
	bs.Add(true)
	bs.Add(true)
	secondBs := bs.Copy()
	secondBs.Add(false)
	if bs.String() != "11" {
		t.Fatalf("Original BitSet is not correct")
	}
	if secondBs.String() != "110" {
		t.Fatalf("Copy BitSet is not correct")
	}
}
