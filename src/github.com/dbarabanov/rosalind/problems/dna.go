//Problem
//
//A string is simply an ordered collection of symbols selected from some alphabet and formed into a word; the length of a string is the number of symbols that it contains.
//
//An example of a length 21 DNA string (whose alphabet contains the symbols 'A', 'C', 'G', and 'T') is "ATGCTTCAGAAAGGTCTTACG."
//
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

package problems

import (
	"fmt"
	"strings"
)

func CountNucleotides(dna_string string) (counts string) {
	var seps = [...]string{"A", "C", "G", "T"}
	var out [4]int
	for i, sep := range seps {
		out[i] = strings.Count(dna_string, sep)
	}
	var result = fmt.Sprintf("%v", out)
	return result[1 : len(result)-1]
}

type StringIterator func() (letter byte, ok bool)

func MakeStringIterator(s string) StringIterator {
	i := -1
	var letters = []byte(s)
	return func() (byte, bool) {
		for i+1 < len(letters) {
			i++
			return s[i], true
		}
		return 0, false
	}
}
