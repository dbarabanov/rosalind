package suffix_tree

import (
	//	"fmt"
	"testing"
)

func TestSuffixTree(t *testing.T) {
		text := "banana"
//	text := "ATGGTCTACATAGCTGACAAACAGCACGTAGCAATCGGTCGAATCTCGAGAGGCATATGGTCACATGATCGGTCGAGCGTGTTTCAAAGTTTGCGCCTAG"
	//	text := "ATGGTCTACATA"
//	pattern := "AC"
	    pattern := "an"
	//	pattern := "na"
	st := ConstructSuffixTree(text)
	//	out := map[int]struct{}{2:struct{}{}, 4:struct{}{}}
	out := map[int]struct{}{1: struct{}{}, 3: struct{}{}}
	x := FindSubstrings(st, pattern)
	if len(x) != len(out) {
		t.Errorf("FindSubstrings %v in %v failed. want %v. found %v",
			pattern, text, out, x)
		return
	}

	for i := range out {
		if x[i] != out[i] {
			t.Errorf("FindSubstrings %v in %v failed. want %v. found %v. diff in %v",
				pattern, text, out, x, i)
			return
		}
	}
}
