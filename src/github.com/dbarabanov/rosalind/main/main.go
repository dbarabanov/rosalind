package main

import (
	"fmt"
	"github.com/dbarabanov/rosalind/problems"
)

func main() {
	fmt.Printf("in main\n")
	fmt.Printf("%v\n", problems.CountNucleotides("AT"))
	iter := problems.MakeStringIterator("ATCGT")
	var counts = map[byte]int{}
	for {
		letter, ok := iter()
		if !ok {
			break
		}
		counts[letter]++
	}
	fmt.Println(counts)
}
