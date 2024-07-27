package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestGetCorrectByteFrequency(t *testing.T) {
	stringReader := bufio.NewReader(strings.NewReader("aaabbc"))
	frequencies := GetByteFrequency(stringReader)

	expected_frequencies := []ByteFrequency{
		{
			value:     'a',
			frequency: 3,
		},
		{
			value:     'b',
			frequency: 2,
		},
		{
			value:     'c',
			frequency: 1,
		},
	}
	result := reflect.DeepEqual(frequencies, expected_frequencies)
	if !result {
		t.Fatal("Byte frequency is not as expected")
	}
}

func TestGetIncorrectByteFrequency(t *testing.T) {
	stringReader := bufio.NewReader(strings.NewReader("aaabbc"))
	frequencies := GetByteFrequency(stringReader)

	expected_frequencies := []ByteFrequency{
		{
			value:     'a',
			frequency: 4,
		},
		{
			value:     'b',
			frequency: 2,
		},
		{
			value:     'c',
			frequency: 1,
		},
	}
	result := reflect.DeepEqual(frequencies, expected_frequencies)
	if result {
		t.Fatal("Byte frequency is not as expected")
	}
}
