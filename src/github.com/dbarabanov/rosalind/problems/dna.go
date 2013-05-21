package problems

import (
	"fmt"
	"github.com/dbarabanov/rosalind/suffix_tree"
	"github.com/dbarabanov/rosalind/util"
	"io/ioutil"
	//"math"
	"strings"
)

//Given: A DNA string s of length at most 1000 nt.
//
//Return: Four integers (separated by spaces) counting the respective number of times that the symbols 'A', 'C', 'G', and 'T' occur in s.
//
//Sample Dataset
//
//AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC
//Sample Output
//
//20 12 17 21
func CountNucleotides(dna_string string) (counts string) {
	var seps = [...]string{"A", "C", "G", "T"}
	var out [4]int
	for i, sep := range seps {
		out[i] = strings.Count(dna_string, sep)
	}
	var result = fmt.Sprintf("%v", out)
	return result[1 : len(result)-1]
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
func TranscribeRna(dna string) (rna string) {
	dna_bytes := []byte(dna)
	rna_bytes := make([]byte, len(dna_bytes))
	for i, dna_byte := range dna_bytes {
		if dna_byte == byte('T') {
			rna_bytes[i] = byte('U')
		} else {
			rna_bytes[i] = dna_byte
		}
	}
	return string(rna_bytes)
}

//The reverse complement of a DNA string s is the string sc formed by reversing the symbols of s, then taking the complement of each symbol (e.g., the reverse complement of "GTCA" is "TGAC").
//
//Given: A DNA string s of length at most 1000 bp.
//
//Return: The reverse complement sc of s.
//
//Sample Dataset
//
//AAAACCCGGT
//Sample Output
//
//ACCGGGTTTT
func ReverseComplement(dna string) (reverse_complement string) {

	f_reverse := func(dna []byte) (reversed []byte) {
		reversed = make([]byte, len(dna))
		for i, dna_byte := range dna {
			reversed[len(dna)-i-1] = dna_byte
		}
		return reversed
	}

	f_complement := func(dna []byte) (complement []byte) {
		complement_map := map[byte]byte{'A': 'T', 'C': 'G', 'G': 'C', 'T': 'A'}
		complement = make([]byte, len(dna))
		for i, dna_byte := range dna {
			complement[i] = complement_map[dna_byte]
		}
		return complement
	}
	return string(f_reverse(f_complement([]byte(dna))))
}

//The 20 commonly occurring amino acids are abbreviated by using 20 letters from the English alphabet (all letters except for B, J, O, U, X, and Z). Protein strings are constructed from these 20 symbols. Henceforth, the term genetic string will incorporate protein strings along with DNA strings and RNA strings.
//
//The RNA codon table dictates the details regarding the encoding of specific codons into the amino acid alphabet.
//
//Given: An RNA string s corresponding to a strand of mRNA (of length at most 10 kbp).
//
//Return: The protein string encoded by s.
//
//Sample Dataset
//
//AUGGCCAUGGCGCCCAGAACUGAGAUCAAUAGUACCCGUAUUAACGGGUGA
//Sample Output
//
//MAMAPRTEINSTRING
func EncodeProtein(rna_string string) (protein_string string) {
	CODON_LENGTH := 3
	protein_index := 0
	codon_index := 0
	rna := []byte(rna_string)
	protein := make([]byte, len(rna)/CODON_LENGTH)
	var codon = make([]byte, CODON_LENGTH, CODON_LENGTH)
	for _, letter := range rna {
		codon[codon_index] = letter
		codon_index++
		if codon_index >= 3 {
			rna_letter := util.CodonMap[string(codon)]
			if rna_letter == 0 {
				break
			}
			protein[protein_index] = rna_letter
			protein_index++
			codon_index = 0
		}
	}
	return string(protein[:protein_index])
}

//After identifying the exons and introns of an RNA string, we only need to delete the introns and concatenate the exons to form a new string ready for translation.
//
//Given: A DNA string s (of length at most 1 kbp) and a collection of substrings of s acting as introns. All strings are given in FASTA format.
//
//Return: A protein string resulting from transcribing and translating the exons of s. (Note: Only one solution will exist for the dataset provided.)
//
//Sample Dataset
//
//>Rosalind_10
//ATGGTCTACATAGCTGACAAACAGCACGTAGCAATCGGTCGAATCTCGAGAGGCATATGGTCACATGATCGGTCGAGCGTGTTTCAAAGTTTGCGCCTAG
//>Rosalind_12
//ATCGGTCGAA
//>Rosalind_15
//ATCGGTCGAGCGTGT
//Sample Output
//
//MVYIADKQHVASREAYGHMFKVCA
func SpliceRna(filename string) (protein string) {
	dna, introns, err := readSplcInput(filename)
	if err != nil {
		panic("failed to read input from " + filename)
	}
	st := suffix_tree.ConstructSuffixTree(dna)
	intronOffsets := make(map[int]int)
	for _, intron := range introns {
		for pos := range suffix_tree.FindSubstrings(st, intron) {
			intronOffsets[pos] = len(intron)
		}
	}
	exons := make([]rune, len(dna))
	intronEnd := 0
	inIntron := false
	l := 0
	for i, r := range dna {
		if length, present := intronOffsets[i]; present && i+length >= intronEnd {
			intronEnd = length + i
			inIntron = true
		}
		if i >= intronEnd {
			inIntron = false
		}
		if !inIntron {
			exons[l] = r
			l++
		}
	}
	//	fmt.Printf("dna  : %v\n", string(dna))
	//	fmt.Printf("exons: %v\n", string(exons))
	retVal := EncodeProtein(TranscribeRna(string(exons)))
	//	fmt.Printf("retVal: %v\n", string(retVal))
	return retVal
}

func RnaSplice(input string) (protein string) {
	var lineBreaks []int
	for i, r := range input {
		if r == '\n' {
			lineBreaks = util.AppendInt(lineBreaks, i)
		}
	}
	dna := input[lineBreaks[0]+1 : lineBreaks[1]]
	//	fmt.Println(dna)
	st := suffix_tree.ConstructSuffixTree(dna)
	intronOffsets := make(map[int]int)
	intronStart := lineBreaks[1] + 1
	for i, lineBreak := range lineBreaks[2:] {
		if i%2 == 1 {
			//			fmt.Println(input[intronStart:lineBreak])
			//			fmt.Println(suffix_tree.FindSubstrings(st, input[intronStart:lineBreak]))
			for k, _ := range suffix_tree.FindSubstrings(st, input[intronStart:lineBreak]) {
				//				intronOffsets = util.AppendInt(intronOffsets, k)
				intronOffsets[k] = lineBreak - intronStart
			}
			//fmt.Printf("intronOffsets: %v\n", intronOffsets)
		}
		intronStart = lineBreak + 1
	}

	exons := make([]rune, len(dna))
	intronEnd := 0
	inIntron := false
	l := 0
	for i, r := range dna {
		if length, present := intronOffsets[i]; present && i+length >= intronEnd {
			//intronEnd = intronOffsets[i]
			intronEnd = length + i
			//fmt.Printf("in new intron: %v\n", intronEnd)
			inIntron = true
		}
		if i >= intronEnd {
			inIntron = false
		}
		if !inIntron {
			exons[l] = r
			l++
		}
	}
	//fmt.Printf("dna  : %v\n", string(dna))
	//fmt.Printf("exons: %v\n", string(exons))
	//return TranscribeRna(string(exons))
	retVal := EncodeProtein(TranscribeRna(string(exons)))
	//	fmt.Printf("retVal: %v\n", string(retVal))
	return retVal
	//	return input
}

func readSplcInput(filename string) (dna string, introns []string, err error) {
	//	content, err := ioutil.ReadFile("test_data/SPLC_in_big.txt")
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return dna, nil, err
	}
	lines := strings.Split(string(content), "\n")
	dna = ""
	isDnaRead := false
	for _, line := range lines[1:] {
		if len(line) > 0 && line[0] == '>' {
			isDnaRead = true
			continue
		}
		if !isDnaRead {
			dna += line
		} else {
			introns = append(introns, line)
		}
	}
	//	fmt.Println(dna)
	//	fmt.Println(introns)
	return dna, introns, nil
}

//Given: Three positive integers k, m, and n, representing a population containing k+m+n organisms: k individuals are homozygous dominant for a factor, m are heterozygous, and n are homozygous recessive.

//Return: The probability that two randomly selected mating organisms will produce an individual possessing a dominant allele (and thus displaying the dominant phenotype). Assume that any two organisms can mate.

//Sample Dataset

//2 2 2
//Sample Output

//0.78333
func probDominant(k int, m int, n int) (prob float64) {
	prob = 1 - (float64(n*(n-1)+n*m)+float64(m*(m-1))/4)/float64((k+m+n)*(k+m+n-1))
	return util.Round(prob, 5)
}
