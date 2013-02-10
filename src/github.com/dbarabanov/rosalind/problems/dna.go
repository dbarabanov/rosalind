package problems

import (
	"fmt"
	//	"github.com/dbarabanov/rosalind/util"
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
