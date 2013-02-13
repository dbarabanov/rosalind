package suffix_tree

import "fmt"

type SuffixTree struct {
	root *Node
}

const Terminator = '$'

type Node struct {
	Edges []*Edge
}

type Edge struct {
	Value []int32
	Start *Node
	End   *Node
}

func ConstructSuffixTree(s string) (t *SuffixTree) {
	fmt.Println("constructing suffix tree from %v", s)
	st := new(SuffixTree)
	root := new(Node)
	text := []int32(s)
	for i := range text {
		addSuffix(root, text[i:])
	}
	st.root = root
	fmt.Printf("constructed: %v\n", st)
	fmt.Printf("constructed: %p\n", st)
	Print(st)
	return st
}

func Print(st *SuffixTree) {
	PrintNode(st.root)
}

func PrintNode(node *Node) {
	if node == nil {
		return
	}

	for _, edge := range node.Edges {
		fmt.Printf("%v\n", string(edge.Value))
		PrintNode(edge.End)

	}
}

func AppendEdge(edges []*Edge, newEdge *Edge) []*Edge {
	m := len(edges)
	fmt.Println("edges: %v. m: %v", edges, m)
	if m+1 > cap(edges) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newEdges := make([]*Edge, (m+1)*2)
		copy(newEdges, edges)
		edges = newEdges
	}
	fmt.Println("edges: %v", edges)
	edges = edges[0 : m+1]
	//	copy(edges[m:m+1], newEdge)
	//	edges[m+1] = newEdge
	edges[m] = newEdge
	return edges
}

func addSuffix(node *Node, suffix []int32) {
	if node.Edges == nil {
		fmt.Println("no edges")
		node.Edges = make([]*Edge, 1)
		node.Edges[0] = &Edge{suffix, node, nil}
	} else {
		fmt.Println("has edges")
//		AppendEdge(node.Edges, &Edge{suffix, node, nil})
        node.Edges = AppendEdge(node.Edges, &Edge{suffix, node, nil})
 
		fmt.Println("added: %v", node.Edges)
	}

	fmt.Println(node.Edges)
}

func FindSubstrings(st *SuffixTree, substring string) (indexes []int32) {
	indexes = make([]int32, 1)
	return indexes
}
