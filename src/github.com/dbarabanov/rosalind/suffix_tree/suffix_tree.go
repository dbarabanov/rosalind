package suffix_tree

import "fmt"

type SuffixTree struct {
	root *Node
}

//const Terminator = '$'
const Terminator = "$"

type Node struct {
	Edges []*Edge
}

type Edge struct {
	Value []int32
	Start *Node
	End   *Node
}

func ConstructSuffixTree(s string) (t *SuffixTree) {
	fmt.Println("constructing suffix tree from", s)
	root := &Node{nil}
	st := &SuffixTree{root}
	text := []int32(s + Terminator)
	for i := range text {
		addSuffix(root, text[i:])
	}
	Print(st)
	return st
}

func Print(st *SuffixTree) {
	fmt.Println("whole tree:")
	PrintNode(st.root, 0)

	//	DeprPrintNode(st.root)
}

func PrintNode(node *Node, depth int) {
	if node == nil {
		return
	}

	spaces := string(make([]int32, depth))
	for i := 0; i < depth; i++ {
		spaces += " "
	}

	for _, edge := range node.Edges {
		fmt.Printf("%v%v\n", spaces, string(edge.Value))
		PrintNode(edge.End, depth+2)
	}
}

func DeprPrintNode(node *Node) {
	if node == nil {
		return
	}

	for _, edge := range node.Edges {
		fmt.Printf("%v\n", string(edge.Value))
		DeprPrintNode(edge.End)
	}
}

func AppendEdge(node *Node, newEdge *Edge) {
	m := len(node.Edges)
	if m+1 > cap(node.Edges) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newEdges := make([]*Edge, (m+1)*2)
		copy(newEdges, node.Edges)
		node.Edges = newEdges
	}
	node.Edges = node.Edges[0 : m+1]
	node.Edges[m] = newEdge
}

func addSuffix(node *Node, suffix []int32) {
	fmt.Printf("adding suffix %v\n", string(suffix))
	for _, edge := range node.Edges {
		//		fmt.Printf("%v\n", edge)
		if edge.Value[0] == suffix[0] {
			fmt.Printf("matching edge for %v: %v\n", string(suffix), string(edge.Value))
			for i := range edge.Value {
				//				fmt.Printf("index: %v. edge.Value: %v\n", i, string(edge.Value[i]))
				if edge.Value[i] != suffix[i] {
					insertNode(edge, i, suffix[i:])
					return
				}
			}
			return
		}
	}
	AppendEdge(node, &Edge{suffix, node, nil})
}

func insertNode(edge *Edge, index int, suffix []int32) {
	fmt.Printf("inserting suffix %v at index %v of %v\n", string(suffix), index, string(edge.Value))
	firstPart := &Edge{edge.Value[:index], edge.Start, &Node{nil}}
	AppendEdge(firstPart.End, &Edge{suffix, firstPart.End, &Node{nil}})
	AppendEdge(firstPart.End, &Edge{edge.Value[index:], firstPart.End, edge.End})
	edge.End = firstPart.End
	edge.Value = firstPart.Value
	DeprPrintNode(edge.Start)
}

func FindSubstrings(st *SuffixTree, substring string) (indexes []int32) {
	indexes = make([]int32, 1)
	return indexes
}
