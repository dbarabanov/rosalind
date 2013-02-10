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

func TestTranscribeRna(t *testing.T) {
	var in, out = "GATGGAACTTGACTACGTAAATT", "GAUGGAACUUGACUACGUAAAUU"
	if x := TranscribeRna(in); x != out {
		t.Errorf("TranscribeRna(%v) = %v, want %v", in, x, out)
	}
}

func TestReverseComplement(t *testing.T) {
	var in, out = "AAAACCCGGT", "ACCGGGTTTT"
	if x := ReverseComplement(in); x != out {
		t.Errorf("ReverseComplement(%v) = %v, want %v", in, x, out)
	}
}

