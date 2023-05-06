package tree

import (
	"github.com/christoffer1009/cobweb-go/node"
	"github.com/christoffer1009/cobweb-go/occurrence"
)

type Tree struct {
	Root *node.Node
	UC   float64
}

func NewTree(node *node.Node) *Tree {
	return &Tree{
		Root: node,
		UC:   0,
	}
}

func (t *Tree) PrintTree() {
	if t.Root != nil {
		t.Root.PrintID()
	}
}

func (t *Tree) CountNodes(n *node.Node) int {
	if n == nil {
		return 0
	}

	count := 0

	for _, child := range n.Children {
		count += 1 + t.CountNodes(child)
	}
	return count
}

func (t *Tree) CalcUC(occurrences []*occurrence.Occurrence) float64 {
	nodes_count := t.CountNodes(t.Root)
	var uc float64 = 0
	if nodes_count > 0 {
		uc = node.SumP(t.Root, occurrences) / float64(nodes_count)

	} else {
		uc = node.SumP(t.Root, occurrences)
	}

	t.UC = uc
	return uc
}
