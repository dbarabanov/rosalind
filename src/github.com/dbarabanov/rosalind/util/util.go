package util

import (
	"io/ioutil"
	"math"
	"strings"
)

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

var CodonMap = map[string]byte{"UUU": 'F', "CUU": 'L', "AUU": 'I', "GUU": 'V', "UUC": 'F', "CUC": 'L', "AUC": 'I', "GUC": 'V', "UUA": 'L', "CUA": 'L', "AUA": 'I', "GUA": 'V', "UUG": 'L', "CUG": 'L', "AUG": 'M', "GUG": 'V', "UCU": 'S', "CCU": 'P', "ACU": 'T', "GCU": 'A', "UCC": 'S', "CCC": 'P', "ACC": 'T', "GCC": 'A', "UCA": 'S', "CCA": 'P', "ACA": 'T', "GCA": 'A', "UCG": 'S', "CCG": 'P', "ACG": 'T', "GCG": 'A', "UAU": 'Y', "CAU": 'H', "AAU": 'N', "GAU": 'D', "UAC": 'Y', "CAC": 'H', "AAC": 'N', "GAC": 'D', "UAA": 0, "CAA": 'Q', "AAA": 'K', "GAA": 'E', "UAG": 0, "CAG": 'Q', "AAG": 'K', "GAG": 'E', "UGU": 'C', "CGU": 'R', "AGU": 'S', "GGU": 'G', "UGC": 'C', "CGC": 'R', "AGC": 'S', "GGC": 'G', "UGA": 0, "CGA": 'R', "AGA": 'R', "GGA": 'G', "UGG": 'W', "CGG": 'R', "AGG": 'R', "GGG": 'G'}

func AppendInt(array []int, value int) []int {
	m := len(array)
	if m+1 > cap(array) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newArray := make([]int, (m+1)*2)
		copy(newArray, array)
		array = newArray
	}
	array = array[0 : m+1]
	array[m] = value
	return array
}

func Round(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)

	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}

func ReadFasta(filename string) map[string]string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("could not open file: " + filename)
	}
	out := make(map[string]string)
	seq := ""
	id := ""
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if len(line) > 0 && line[0] == '>' {
			if seq != "" {
				out[id] = seq
				seq = ""
			}
			id = line[1:]
		} else {
			seq = seq + line
		}
	}
    out[id] = seq
	return out
}
