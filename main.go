package main

import (
	"fmt"

	"github.com/christoffer1009/cobweb-go/node"
	"github.com/christoffer1009/cobweb-go/occurrence"
	"github.com/christoffer1009/cobweb-go/tree"
)

func main() {

	root := node.NewNode(0)
	cobweb := tree.NewTree(root)

	child1 := node.NewNode(1)
	child2 := node.NewNode(2)

	oc1 := occurrence.NewOcurrence(1, "w", 1, 1)
	oc2 := occurrence.NewOcurrence(2, "w", 2, 2)
	oc3 := occurrence.NewOcurrence(3, "b", 2, 2)
	// oc4 := occurrence.NewOcurrence(4, "b", 3, 1)

	cobweb.Root.AddChild(child1)
	cobweb.Root.AddChild(child2)

	cobweb.Root.AddOccurrence(oc1)
	cobweb.Root.AddOccurrence(oc2)
	cobweb.Root.AddOccurrence(oc3)
	// cobweb.Root.AddOccurrence(oc4)

	child1.AddOccurrence(oc1)
	child1.AddOccurrence(oc3)
	child2.AddOccurrence(oc2)

	root.PrintOccurrences()
	cobweb.CalcUC(root.Occurrences)

	fmt.Println("FIM")

}
