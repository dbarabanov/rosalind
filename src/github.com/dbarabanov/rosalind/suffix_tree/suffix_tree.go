package suffix_tree

import "fmt"

type SuffixTree struct {
	root *Node
}

const Terminator = "$"

type NodeData struct {
	LeafIndex int
}

type Node struct {
	Edges    []*Edge
	NodeData *NodeData
}

type Edge struct {
	Label []rune
	Start *Node
	End   *Node
}

func ConstructSuffixTree(s string) (t *SuffixTree) {
//	fmt.Println("constructing suffix tree from", s)
	root := &Node{nil, nil}
	st := &SuffixTree{root}
	text := []rune(s + Terminator)
	for i := range text {
		addSuffix(root, text[i:], i)
	}
	//	Print(st)
	return st
}

func Print(st *SuffixTree) {
	fmt.Println("whole tree:")
	PrintNode(st.root, 0)
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
		fmt.Printf("%v%v\n", spaces, string(edge.Label))
		PrintNode(edge.End, depth+2)
	}
}

func DeprPrintNode(node *Node) {
	if node == nil {
		return
	}

	for _, edge := range node.Edges {
		fmt.Printf("%v\n", string(edge.Label))
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

func addSuffix(node *Node, suffix []rune, position int) {
	//	fmt.Printf("adding suffix %v\n", string(suffix))
	for _, edge := range node.Edges {
		//		fmt.Printf("%v\n", edge)
		if edge.Label[0] == suffix[0] {
			//			fmt.Printf("matching edge for %v: %v\n", string(suffix), string(edge.Label))
			for i := range edge.Label {
				//				fmt.Printf("index: %v. edge.Label: %v\n", i, string(edge.Label[i]))
				//				fmt.Printf("index: %v. edge.Label: %v\n", i, string(edge.Label))
				//				fmt.Printf("index: %v. suffix: %v\n", i, string(suffix))
				if edge.Label[i] != suffix[i] {
					insertNode(edge, i, suffix[i:], position)
					return
				}
			}
			addSuffix(edge.End, suffix[len(edge.Label):], position)
			return
		}
	}
	//	AppendEdge(node, &Edge{suffix, node, nil})
	AppendEdge(node, &Edge{suffix, node, &Node{nil, &NodeData{position}}})
}

func insertNode(edge *Edge, index int, suffix []rune, position int) {
	//	fmt.Printf("inserting suffix %v at index %v of %v\n", string(suffix), index, string(edge.Label))
	firstPart := &Edge{edge.Label[:index], edge.Start, &Node{nil, nil}}
	AppendEdge(firstPart.End, &Edge{suffix, firstPart.End, &Node{nil, &NodeData{position}}})
	AppendEdge(firstPart.End, &Edge{edge.Label[index:], firstPart.End, edge.End})
	edge.End = firstPart.End
	edge.Label = firstPart.Label
	//	DeprPrintNode(edge.Start)
}

func FindSubstrings(st *SuffixTree, substring string) (indexes map[int]struct{}) {
	if st == nil || st.root == nil || len(substring) <= 0{
		return indexes
	}
	return findSubstrings(st.root, []rune(substring))
}

func findSubstrings(node *Node, pattern []rune) (indexes map[int]struct{}) {
	index := 0
	var activeEdge *Edge = nil
	for _, edge := range node.Edges {
		//fmt.Println(string(edge.Label))
		if edge.Label[0] == pattern[index] {
			activeEdge = edge
			break
		}
	}
	if activeEdge != nil {
		//fmt.Printf("activeEdge: %v. index: %v\n", string(activeEdge.Label), index)
		for _, l := range activeEdge.Label {
			if index >= len(pattern) {
				//fmt.Println("A")
				//				TraverseNodes(activeEdge.End, indexes)
				return collectLeafsBelow(activeEdge.End)
			}
			if pattern[index] != l {
				//fmt.Println("B")
				return indexes
			}
			index++
		}
		if index >= len(pattern) {
			return collectLeafsBelow(activeEdge.End)
		}
		return findSubstrings(activeEdge.End, pattern[len(activeEdge.Label):])
	}
	return indexes
}

func collectLeafsBelow(node *Node) (indexes map[int]struct{}) {
	indexes = make(map[int]struct{})
	ch := Walker(node)
	for {
		index, ok := <-ch
		if !ok {
			return indexes
		}
		indexes[index] = struct{}{}
	}
	return indexes
}

func WalkLeafs(n *Node, ch chan int) {
	if n == nil {
		panic("attempting to walk nil node")
	}
	if len(n.Edges) <= 0 {
		ch <- n.NodeData.LeafIndex
	}
	for _, edge := range n.Edges {
		WalkLeafs(edge.End, ch)
	}
}

func Walker(n *Node) <-chan int {
	ch := make(chan int)
	go func() {
		WalkLeafs(n, ch)
		close(ch)
	}()
	return ch
}

func TraverseNodes(node *Node, indexes []int) {
	//	fmt.Printf("Traversing node %v\n", node.NodeData)
	if len(node.Edges) <= 0 {
		//fmt.Printf("Leaf %v\n", node.NodeData)
		AppendInt(indexes, node.NodeData.LeafIndex)
		//		return indexes
	} else {
		for _, edge := range node.Edges {
			TraverseNodes(edge.End, indexes)
		}
		//		return indexes
	}
}

func AppendInt(indexes []int, index int) []int {
	m := len(indexes)
	if m+1 > cap(indexes) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newIndexes := make([]int, (m+1)*2)
		copy(newIndexes, indexes)
		indexes = newIndexes
	}
	indexes = indexes[0 : m+1]
	indexes[m] = index
	return indexes
}
