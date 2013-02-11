package problems

import (
	"fmt"
	"github.com/dbarabanov/rosalind/util"
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
