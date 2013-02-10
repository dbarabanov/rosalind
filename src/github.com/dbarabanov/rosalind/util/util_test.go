package util

import (
	//	"fmt"
	"testing"
)

func TestStringIterator(t *testing.T) {
	var s = "ATTGCT"
	iter := MakeStringIterator(s)
	letters := map[byte]int{}
	for {
		letter, ok := iter()
		if !ok {
			break
		}
		letters[letter]++
	}
	expected := map[byte]int{65: 1, 71: 1, 67: 1, 84: 3}
	for key, value := range letters {
		if value != expected[key] {
			t.Errorf("StringIterator(%v) = %v, want %v", s, letters, expected)
		}
	}
}
