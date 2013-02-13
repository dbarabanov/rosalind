package suffix_tree

import (
	//	"fmt"
	"testing"
)

func TestSuffixTree(t *testing.T) {
	st := ConstructSuffixTree("banana")
	FindSubstrings(st, "an")
}
