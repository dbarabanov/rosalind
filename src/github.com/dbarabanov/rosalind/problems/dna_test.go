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

//Given a DNA string t corresponding to a coding strand, its transcribed RNA string u is formed by replacing all occurrences of 'T' in t with 'U' in u.
//
//Given: A DNA string t having length at most 1000 nt.
//
//Return: The transcribed RNA string of t.
//
//Sample Dataset
//
//GATGGAACTTGACTACGTAAATT
//Sample Output
//
//GAUGGAACUUGACUACGUAAAUU
func TestTranscribeRna(t *testing.T) {
	var in, out = "GATGGAACTTGACTACGTAAATT", "GAUGGAACUUGACUACGUAAAUU"
	if x := TranscribeRna(in); x != out {
		t.Errorf("TranscribeRna(%v) = %v, want %v", in, x, out)
	}
}
