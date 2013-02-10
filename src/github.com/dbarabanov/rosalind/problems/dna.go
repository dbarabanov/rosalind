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
