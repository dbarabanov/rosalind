package problems

import (
	//	"fmt"
	"io/ioutil"
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

func TestEncodeProtein(t *testing.T) {
	var in, out = "AUGGCCAUGGCGCCCAGAACUGAGAUCAAUAGUACCCGUAUUAACGGGUGA", "MAMAPRTEINSTRING"
	if x := EncodeProtein(in); x != out {
		t.Errorf("EncodeProtein(%v) = %v, want %v", in, x, out)
	}
}

func TestRnaSplice(t *testing.T) {
	var in = `>Rosalind_10
ATGGTCTACATAGCTGACAAACAGCACGTAGCAATCGGTCGAATCTCGAGAGGCATATGGTCACATGATCGGTCGAGCGTGTTTCAAAGTTTGCGCCTAG
Rosalind_12
ATCGGTCGAA
Rosalind_15
ATCGGTCGAGCGTGT
`
	var out = "MVYIADKQHVASREAYGHMFKVCA"
	if x := RnaSplice(in); x != out {
		t.Errorf("RnaSplice(%v) = %v, want %v", in, x, out)
	}
}

func TestSplc(t *testing.T) {
	//	content, err := ioutil.ReadFile("test_data/SPLC_in_big.txt")
	content, err := ioutil.ReadFile("test_data/SPLC_in.txt")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		//	lines := strings.Split(string(content), "\n")
		//in := content
		in := string(content)
		//		t.Errorf(in)
		var out = "MVYIADKQHVASREAYGHMFKVCA"
		if x := RnaSplice(in); x != out {
			t.Errorf("RnaSplice(%v) = %v, want %v", in, x, out)
		}
	}
}

func TestSpliceRna(t *testing.T) {
	//	var out, in = "MVYIADKQHVASREAYGHMFKVCA", "test_data/SPLC_in.txt"
	//	var out, in = "MQCFTHPVEPKLGGVRLRWEEMSQRQVDSLKDCGTLFDGGALYLHATCARSLPEKHREDTILPVCRPGLAEQVIGFCPGANYAHFVSAAGCSQPGSGISTLTGRYGYRISGHNGTGQNTAMRFPLTIVGIHKNRYLSYRTLSVMQHTYTVFSSRGNAYLHLLVVHTSTGLNENSRRGRELSSGRP", "test_data/SPLC_in_big.txt"
	var out, in = "MAAGNLAVTRSKSVALPWHLSLPSIVSYENRKWRMVVIALSSLNAQVDLGVSIMDRGTVNVPLAVDNTILQLRCLLSSFSGPDHGFRPATAHISGLDIRPGLKHCYSSSHVKSGIGHQRLHSSLTQCSPGEKPVRVGRIPGVESDSWVLRIHVEMCTLLGYVMGRRGVHLEKIQRVNGYISVQR", "test_data/SPLC_in_big2.txt"
	if x := SpliceRna(in); x != out {
		t.Errorf("SpliceRna(%v) = %v, want %v", in, x, out)
	}
}

func TestProbDominant(t *testing.T) {
	var k, m, n = 2, 2, 2
	//var k, m, n = 26, 23, 29

	var out = 0.78333
	if x := probDominant(k, m, n); x != out {
		t.Errorf("probDominant(%v, %v, %v) = %v, want %v", k, m, n, x, out)
	}
}

func TestRabbits(t *testing.T) {
	var months, litter = 5, 3

	var out = uint64(19)
	if rabbits := rabbits(months, litter); rabbits != out {
		t.Errorf("rabbits(%v, %v) = %v, want %v", months, litter, rabbits, out)
	}
}
