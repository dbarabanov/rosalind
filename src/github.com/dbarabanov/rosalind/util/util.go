package util

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

var CodonMap = map[string]byte{"UUU":'F',"CUU":'L',"AUU":'I',"GUU":'V',"UUC":'F',"CUC":'L',"AUC":'I',"GUC":'V',"UUA":'L',"CUA":'L',"AUA":'I',"GUA":'V',"UUG":'L',"CUG":'L',"AUG":'M',"GUG":'V',"UCU":'S',"CCU":'P',"ACU":'T',"GCU":'A',"UCC":'S',"CCC":'P',"ACC":'T',"GCC":'A',"UCA":'S',"CCA":'P',"ACA":'T',"GCA":'A',"UCG":'S',"CCG":'P',"ACG":'T',"GCG":'A',"UAU":'Y',"CAU":'H',"AAU":'N',"GAU":'D',"UAC":'Y',"CAC":'H',"AAC":'N',"GAC":'D',"UAA":0,"CAA":'Q',"AAA":'K',"GAA":'E',"UAG":0,"CAG":'Q',"AAG":'K',"GAG":'E',"UGU":'C',"CGU":'R',"AGU":'S',"GGU":'G',"UGC":'C',"CGC":'R',"AGC":'S',"GGC":'G',"UGA":0,"CGA":'R',"AGA":'R',"GGA":'G',"UGG":'W',"CGG":'R',"AGG":'R',"GGG":'G'}

func AppendInt(array []int, value int) []int {
	m := len(array)
	if m+1 > cap(array) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newArray:= make([]int, (m+1)*2)
		copy(newArray, array)
		array = newArray
	}
	array = array[0 : m+1]
	array[m] = value 
	return array
}
