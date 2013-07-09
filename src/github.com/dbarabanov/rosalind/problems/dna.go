package problems

import (
	"fmt"
	"github.com/dbarabanov/rosalind/suffix_tree"
	"github.com/dbarabanov/rosalind/util"
	"io/ioutil"
	//"math"
	"sort"
	"strconv"
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

//Problem

//A sequence is an ordered collection of objects (usually numbers), which are allowed to repeat. Sequences can be finite or infinite. Two examples are the finite sequence (π,−2‾‾√,0,π) and the infinite sequence of odd numbers (1,3,5,7,9,…). We use the notation an to represent the n-th term of a sequence.

//A recurrence relation is a way of defining the terms of a sequence with respect to the values of previous terms. In the case of Fibonacci's rabbits from the introduction, any given month will contain the rabbits that were alive the previous month, plus any new offspring. A key observation is that the number of offspring in any month is equal to the number of rabbits that were alive two months prior. As a result, if Fn represents the number of rabbit pairs alive after the n-th month, then we obtain the Fibonacci sequence having terms Fn that are defined by the recurrence relation Fn=Fn−1+Fn−2 (with F1=F2=1 to initiate the sequence). Although the sequence bears Fibonacci's name, it was known to Indian mathematicians over two millennia ago.

//When finding the n-th term of a sequence defined by a recurrence relation, we can simply use the recurrence relation to generate terms for progressively larger values of n. This problem introduces us to the computational technique of dynamic programming, which successively builds up solutions by using the answers to smaller cases.

//Given: Positive integers n≤40 and k≤5.

//Return: The total number of rabbit pairs that will be present after n months if each pair of reproduction-age rabbits produces a litter of k rabbit pairs in each generation (instead of only 1 pair).

//Sample Dataset

//5 3
//Sample Output

//19
func rabbits(months int, litter int) (rabbits uint64) {
	//F(n) = F(n-1) + k*F(n-1)
	var last, previous uint64
	last, previous = 1, 1
	for now := 2; now < months; now++ {
		rabbits = uint64(uint64(litter)*last + previous)
		last = previous
		previous = rabbits
	}
	return rabbits
}

func proteinWeight(protein string) (weight float64) {
	weight = 0
	weights := map[byte]float64{'A': 71.03711,
		'C': 103.00919,
		'D': 115.02694,
		'E': 129.04259,
		'F': 147.06841,
		'G': 57.02146,
		'H': 137.05891,
		'I': 113.08406,
		'K': 128.09496,
		'L': 113.08406,
		'M': 131.04049,
		'N': 114.04293,
		'P': 97.05276,
		'Q': 128.05858,
		'R': 156.10111,
		'S': 87.03203,
		'T': 101.04768,
		'V': 99.06841,
		'W': 186.07931,
		'Y': 163.06333}
	for _, aa := range protein {
		weight += weights[byte(aa)]
	}

	return util.Round(weight, 3)
}

//Problem

//A graph whose nodes have all been labeled can be represented by an adjacency list, in which each row of the list contains the two node labels corresponding to a unique edge.

//A directed graph (or digraph) is a graph containing directed edges, each of which has an orientation. That is, a directed edge is represented by an arrow instead of a line segment; the starting and ending nodes of an edge form its tail and head, respectively. The directed edge with tail v and head w is represented by (v,w) (but not by (w,v)). A directed loop is a directed edge of the form (v,v).

//For a collection of strings and a positive integer k, the overlap graph for the strings is a directed graph Ok in which each string is represented by a node, and string s is connected to string t with a directed edge when there is a length k suffix of s that matches a length k prefix of t, as long as s≠t; we demand s≠t to prevent directed loops in the overlap graph (although directed cycles may be present).

//Given: A collection of DNA strings in FASTA format having total length at most 10 kbp.

//Return: The adjacency list corresponding to O3. You may return edges in any order.

//Sample Dataset

//>Rosalind_0498
//AAATAAA
//>Rosalind_2391
//AAATTTT
//>Rosalind_2323
//TTTTCCC
//>Rosalind_0442
//AAATCCC
//>Rosalind_5013
//GGGTGGG
//Sample Output

//Rosalind_0498 Rosalind_0442
//Rosalind_0498 Rosalind_2391
//Rosalind_2391 Rosalind_2323
func overlapGraph(filename string) (graph string) {
	k := 3
	reads := util.ReadFasta(filename)
	prefixes := make(map[string][]string)
	for id, seq := range reads {
		prefix := seq[:k]
		prefixes[prefix] = append(prefixes[prefix], id)
	}
	matches := make(map[string]struct{})
	for id, seq := range reads {
		suffix := seq[len(seq)-3:]
		for _, match := range prefixes[suffix] {
			if id != match {
				//matches[id] = match
				//matches += id + " " + match + "\n"
				matches[id+" "+match] = struct{}{}
			}
		}
	}
	var keys []string
	for k := range matches {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	out := ""
	for _, k := range keys {
		out += k + "\n"
	}
	return out
}

//An undirected graph is connected if there is a path connecting any two nodes. A tree is a connected (undirected) graph containing no cycles; this definition forces the tree to have a branching structure organized around a central core of nodes, just like its living counterpart. See Figure 2.

//We have already grown familiar with trees in “Mendel's First Law”, where we introduced the probability tree diagram to visualize the outcomes of a random variable.

//In the creation of a phylogeny, taxa are encoded by the tree's leaves, or nodes having degree 1. A node of a tree having degree larger than 1 is called an internal node.

//Given: A positive integer n (n≤1000) and an adjacency list corresponding to a graph on n nodes that contains no cycles.

//Return: The minimum number of edges that can be added to the graph to produce a tree.

//Sample Dataset

//10
//1 2
//2 8
//4 10
//5 9
//6 10
//7 9
//Sample Output

//3
func completeTree(filename string) (edge_count int) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("could not open file: " + filename)
	}
	type edge struct {
		from string
		to   string
	}
	var edges []edge
	var freeNodes int
	var deadClusters int
	lines := strings.Split(string(content), "\n")
	//fmt.Printf("lines: %v\n", lines)
	for i, line := range lines {
		if i == 0 {
			freeNodes, err = strconv.Atoi(line)
		}
		//fmt.Printf("line: %v\n", line)
		if len(strings.Split(line, " ")) == 2 {
			tokens := strings.Split(line, " ")
			edges = append(edges, edge{tokens[0], tokens[1]})
			//fmt.Printf("edges: %v\n", edges)
		}
	}
	//fmt.Printf("edges: %v\n", edges)

	type cluster struct {
		size    int
		members map[string]bool
	}
	var clusters []*cluster
	var firstCluster, secondCluster *cluster
	for _, edge := range edges {
		firstCluster, secondCluster = nil, nil
		for _, cluster := range clusters {
			if cluster.members[edge.from] {
				firstCluster = cluster
			}
			if cluster.members[edge.to] {
				secondCluster = cluster
			}
		}
		if firstCluster == nil && secondCluster == nil {
			freeNodes -= 2
			newCluster := cluster{2, map[string]bool{edge.from: true, edge.to: true}}
			clusters = append(clusters, &newCluster)
		} else if firstCluster == nil {
			freeNodes -= 1
			secondCluster.size += 1
			secondCluster.members[edge.from] = true
		} else if secondCluster == nil {
			freeNodes -= 1
			firstCluster.size += 1
			firstCluster.members[edge.to] = true
		} else {
			var mainCluster, smallCluster *cluster
			if firstCluster.size >= secondCluster.size {
				mainCluster = firstCluster
				smallCluster = secondCluster
			} else {
				mainCluster = secondCluster
				smallCluster = firstCluster
			}
			mainCluster.size += smallCluster.size
			for node := range smallCluster.members {
				mainCluster.members[node] = true
			}
			smallCluster = nil
			deadClusters++
		}

		//for _, cluster := range clusters {
		//fmt.Printf("%v", cluster)
		//}
		//fmt.Printf("\n")
		//fmt.Printf("freeNodes: %v\n", freeNodes)
		//fmt.Printf("deadClusters: %v\n", deadClusters)
		//fmt.Printf("len(clusters): %v\n", len(clusters))
	}

	return len(clusters) + freeNodes - deadClusters - 1
}

//Newick format is a way of representing trees even more concisely than using an adjacency list, especially when dealing with trees whose internal nodes have not been labeled.

//First, consider the case of a rooted tree T. A collection of leaves v1,v2,…,vn of T are neighbors if they are all adjacent to some internal node u. Newick format for T is obtained by iterating the following key step: delete all the edges {vi,u} from T and label u with (v1,v2,…,vn)u. This process is repeated all the way to the root, at which point a semicolon signals the end of the tree.

//A number of variations of Newick format exist. First, if a node is not labeled in T, then we simply leave blank the space occupied by the node. In the key step, we can write (v1,v2,…,vn) in place of (v1,v2,…,vn)u if the vi are labeled; if none of the nodes are labeled, we can write (,,…,).

//A second variation of Newick format occurs when T is unrooted, in which case we simply select any internal node to serve as the root of T. A particularly peculiar case of Newick format arises when we choose a leaf to serve as the root.

//Note that there will be a large number of different ways to represent T in Newick format; see Figure 1.

//Given: A collection of n trees (n≤40) in Newick format, with each tree containing at most 200 nodes; each tree Tk is followed by a pair of nodes xk and yk in Tk.

//Return: A collection of n positive integers, for which the kth integer represents the distance between xk and yk in Tk.

//Sample Dataset
//(cat)dog;
//dog cat

//(dog,cat);
//dog cat
//Sample Output

//1 2
func distancesInTrees(filename string) (output string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("could not open file: " + filename)
	}
	lines := strings.Split(string(content), "\n")
	//fmt.Printf("lines: %v\n", lines)
	x, y := "", ""
	newick := ""
	readingNewick := true
	var tokens []string
	//var distances []string
	distances := make([]string, 0)
	for _, line := range lines {
		if line != "" {
			//fmt.Printf("line: %v\n", line)
			if readingNewick {
				newick += line
				if line[len(line)-1] == ';' {
					readingNewick = false
				}
			} else {
				tokens = strings.Split(line, " ")
				x, y = tokens[0], tokens[1]
				//fmt.Printf("x, y: %v %v\n", x, y)
				tokens = nil
				token := ""
				for _, symbol := range newick {
					if symbol == ')' || symbol == '(' || symbol == ',' {
						if token != "" {
							tokens = append(tokens, token)
							token = ""
						}
						tokens = append(tokens, string(symbol))
					} else if symbol == ';' {
						if token != "" {
							tokens = append(tokens, token)
							token = ""
						}
						//fmt.Printf("newick: %v\n", newick)
						//fmt.Printf("tokens: %v\n", tokens)
						break
					} else if symbol != ' ' {
						token += string(symbol)
					}
				}
				//fmt.Printf("tokens: %v\n", tokens)
				//fmt.Printf("x, y: %v %v\n", x, y)
				distances = append(distances, distanceInNewickTree(tokens, x, y))
				//fmt.Printf("distances: %v\n", distances)
				readingNewick = true
				newick = ""
			}
		}
	}
	return strings.Join(distances, " ")
}
func distanceInNewickTree(tokens []string, x string, y string) string {
	if x == y {
		return "0"
	}
	distance := 0
	depth := 0
	connectors := make([]int, 0)
	var peak int
	started := false
	var first int
    redundands := findredundands(tokens)
	//redundands := make(map[int]bool)
	for i, token := range tokens {
		if redundands[i] == true {
			continue
		} else if token == "(" {
			connectors = append(connectors, depth)
			depth++
		} else if token == ")" {
			depth = connectors[len(connectors)-1] + 1
			connectors = connectors[:len(connectors)-1]
			if peak >= depth {
				peak = depth - 1
			}
		} else if token == x || token == y {
			if !started {
				first = depth
				peak = depth
				started = true
				//fmt.Printf("*****starting. first: %v depth: %v, peak: %v\n", first, depth, peak)
			} else {
				//fmt.Printf("first: %v depth: %v, peak: %v\n", first, depth, peak)
				distance = first + depth - 2*peak
				break
			}
			depth++
		} else if token == "," {
			depth = connectors[len(connectors)-1] + 1
			if peak >= depth {
				peak = depth - 1
			}
		} else {
			depth++
			//fmt.Printf("increasing depth to : %v token '%v' y: '%v' \n", depth, token, y)
		}
		//fmt.Printf("token: %v depth: %v, connectors: %v\n", token, depth, connectors)
	}
	return strconv.Itoa(distance)
}

func findredundands(tokens []string) map[int]bool {
	redundands := make(map[int]bool)
	type brace struct {
		position  int
		redundand bool
	}
	var braces []brace
	for i, token := range tokens {
		if token == "(" {
			braces = append(braces, brace{i, true})
		} else if token == ")" {
			if braces[len(braces)-1].redundand {
				redundands[braces[len(braces)-1].position] = true
				redundands[i] = true
			}
			braces = braces[:len(braces)-1]
		} else if token == "," {
			braces[len(braces)-1].redundand = false
		}
		//fmt.Printf("i: %v braces: %v redundands: %v token: '%v'\n", i, braces, redundands, token)
	}
	//fmt.Printf("redundands: %v\n", redundands)
	return redundands
}

//A matrix is a rectangular table of values divided into rows and columns. An m×n matrix has m rows and n columns. Given a matrix A, we write Ai,j to indicate the value found at the intersection of row i and column j.

//Say that we have a collection of DNA strings, all having the same length n. Their profile matrix is a 4×n matrix P in which P1,j represents the number of times that 'A' occurs in the jth position of one of the strings, P2,j represents the number of times that C occurs in the jth position, and so on (see below).

//A consensus string c is a string of length n formed from our collection by taking the most common symbol at each position; the jth symbol of c therefore corresponds to the symbol having the maximum value in the j-th column of the profile matrix. Of course, there may be more than one most common symbol, leading to multiple possible consensus strings.

//A T C C A G C T
//G G G C A A C T
//A T G G A T C T
//DNA Strings A A G C A A C C
//T T G G A A C T
//A T G C C A T T
//A T G G C A C T
//A   5 1 0 0 5 5 0 0
//Profile C   0 0 1 4 2 0 6 1
//G   1 1 6 3 0 1 0 0
//T   1 5 0 0 0 1 1 6
//Consensus   A T G C A A C T
//Given: A collection of at most 10 DNA strings of equal length (at most 1 kbp) in FASTA format.

//Return: A consensus string and profile matrix for the collection. (If several possible consensus strings exist, then you may return any one of them.)

//Sample Dataset

//>Rosalind_1
//ATCCAGCT
//>Rosalind_2
//GGGCAACT
//>Rosalind_3
//ATGGATCT
//>Rosalind_4
//AAGCAACC
//>Rosalind_5
//TTGGAACT
//>Rosalind_6
//ATGCCATT
//>Rosalind_7
//ATGGCACT
//Sample Output

//ATGCAACT
//A: 5 1 0 0 5 5 0 0
//C: 0 0 1 4 2 0 6 1
//G: 1 1 6 3 0 1 0 0
//T: 1 5 0 0 0 1 1 6
func consensusAndProfile(filename string) (result string) {
	reads := util.ReadFasta(filename)
	//fmt.Printf("reads: %v\n", reads)
	letters := []rune{'A', 'C', 'G', 'T'}

	var profile map[rune][]int
	var rowlen int
	for _, seq := range reads {
		rowlen = len(seq)
		if profile == nil {
			profile = map[rune][]int{'A': make([]int, rowlen), 'T': make([]int, rowlen), 'G': make([]int, rowlen), 'C': make([]int, rowlen)}
		}
		for i, l := range seq {
			profile[l][i]++
		}
	}

	var consensus []rune
	for i := range profile['A'] {
		var winner rune
		maxScore := 0
		for _, k := range letters {
			if profile[k][i] >= maxScore {
				winner = k
				maxScore = profile[k][i]
			}
		}
		consensus = append(consensus, winner)
	}

	result = string(consensus) + "\n"
	for _, l := range letters {
		result += string(l) + ": " + strings.Trim(fmt.Sprintf("%v", profile[l]), "[]") + "\n"
	}
	//ioutil.WriteFile("test_data/consensusAndProfile.out", []byte(result), 0644)

	return result
}
