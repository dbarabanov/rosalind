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

func TestProteinWeight(t *testing.T) {
	protein := "SKADYEK"

	var out = 821.392
	if x := proteinWeight(protein); x != out {
		t.Errorf("proteinWeight(%v) = %v, want %v", protein, x, out)
	}
}

func TestOverlapGraph(t *testing.T) {
	//input := `>Rosalind_0498
	//AAATAAA
	//>Rosalind_2391
	//AAATTTT
	//>Rosalind_2323
	//TTTTCCC
	//>Rosalind_0442
	//AAATCCC
	//>Rosalind_5013
	//GGGTGGG`
	filename := "test_data/overlapGraphs.input"
	var out = `Rosalind_0498 Rosalind_0442
Rosalind_0498 Rosalind_2391
Rosalind_2391 Rosalind_2323
`

	//var out = make(map[string]struct{})
	//out["Rosalind_0498 Rosalind_2391"] = struct{}{}
	//out["Rosalind_0498 Rosalind_0442"] = struct{}{}
	//out["Rosalind_2391 Rosalind_2323"] = struct{}{}
	if x := overlapGraph(filename); x != out {
		t.Errorf("overlapGraph(%v) = %v, want %v", filename, x, out)
	}
}

func TestCompleteTree(t *testing.T) {
	filename := "test_data/completeTree.input"
	var out = 3
	if x := completeTree(filename); x != out {
		t.Errorf("completeTree(%v) = %v, want %v", filename, x, out)
	}
}

func TestDistancesInTrees(t *testing.T) {
	//filename := "test_data/distancesInTrees.prod"
	//t.Fail()
	filename := "test_data/distancesInTrees.input"
	var out = "1 2 3 2 3 2 3 2"
	//var out = "15 7 53 1 25 9 6 2 2 10 9 2 2 43 21 2 18 19 14 30 15 20 11 15 23 10 13 27 26 8 2 2 17 12 22 37 2 1 22"
	//var out = "15 7 53 48 25 9 6 2 2 10 9 2 2 43 21 2 18 19 14 30 15 20 11 15 23 10 13 27 26 8 2 2 17 12 22 37 2 19 22"
	if x := distancesInTrees(filename); x != out {
		t.Errorf("distancesInTrees(%v) = %v, want %v.", filename, x, out)
	}
}

func TestConsensusAndProfile(t *testing.T) {
	filename := "test_data/ConsensusAndProfile.input"
	var out string
	content, err := ioutil.ReadFile("test_data/ConsensusAndProfile.out")
	if err != nil {
		t.Errorf(err.Error())
		panic(err.Error())
	} else {
		out = string(content)
	}

	if x := consensusAndProfile(filename); x != out {
		t.Errorf("consensusAndProfile(%v) = %v, want %v.", filename, x, out)
	}
}
