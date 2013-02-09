package problems

import "testing"

func TestCount_nucleotides(t *testing.T) {
	var in, out = "GC", "0 1 1 0"
	if x := Count_nucleotides(in); x != out {
		t.Errorf("Count_nucleotides(%v) = %v, want %v", in, x, out)
	}

	in, out = "AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC", "20 12 17 21"
	if x := Count_nucleotides(in); x != out {
		t.Errorf("Count_nucleotides(%v) = %v, want %v", in, x, out)
	}

	in, out = "йываждло", "0 0 0 0"
	if x := Count_nucleotides(in); x != out {
		t.Errorf("Count_nucleotides(%v) = %v, want %v", in, x, out)
	}

}
