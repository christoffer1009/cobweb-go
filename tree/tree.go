package tree

import (
	"github.com/christoffer1009/cobweb-go/node"
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

func (t *Tree) CalcUC() float64 {
	return 0
}
