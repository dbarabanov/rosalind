package problems

import (
	//	"fmt"
	"testing"
)

func TestCountNucleotides(t *testing.T) {
	var in, out = "GC", "0 1 1 0"
	if x := CountNucleotides(in); x != out {
		t.Errorf("CountNucleotides(%v) = %v, want %v", in, x, out)
	}

	in, out = "AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC", "20 12 17 21"
	if x := CountNucleotides(in); x != out {
		t.Errorf("CountNucleotides(%v) = %v, want %v", in, x, out)
	}

	in, out = "йываждло", "0 0 0 0"
	if x := CountNucleotides(in); x != out {
		t.Errorf("CountNucleotides(%v) = %v, want %v", in, x, out)
	}

}

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
