package main

import (
	"fmt"
	"github.com/dbarabanov/rosalind/problems"
	"github.com/dbarabanov/rosalind/util"
)

func main() {
	fmt.Printf("in main\n")
	fmt.Printf("%v\n", problems.CountNucleotides("AT"))
	iter := util.MakeStringIterator("ATCGT")
	var counts = map[byte]int{}
	for {
		letter, ok := iter()
		if !ok {
			break
		}
		counts[letter]++
	}
	fmt.Println(counts)

	var rna = "AUGGCCAUGGCGCCCAGAACUGAGAUCAAUAGUACCCGUAUUAACGGGUGA"

	fmt.Println(problems.EncodeProtein(rna))
}
