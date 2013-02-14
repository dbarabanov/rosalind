package suffix_tree

import (
	//	"fmt"
	"testing"
)

func TestSuffixTree(t *testing.T) {
	text := "banana"
pattern := "an"
//	pattern := "na"
	st := ConstructSuffixTree(text)
	//	indexes := FindSubstrings(st, "an")
	//	out := []int{1, 3}
	//out := []int{2, 4}
//	out := map[int]struct{}{2:struct{}{}, 4:struct{}{}}
	out := map[int]struct{}{1:struct{}{}, 3:struct{}{}}
	x := FindSubstrings(st, pattern)
	for i := range out {
		if len(x) <= 0 || x[i] != out[i] {
			t.Errorf("FindSubstrings %v in %v failed. want %v. found %v. diff in %v",
				pattern, text, out, x, i)
			return
		}
	}
}
