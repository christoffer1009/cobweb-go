package main

import (
	"fmt"

	"github.com/christoffer1009/cobweb-go/node"
	"github.com/christoffer1009/cobweb-go/ocurrence"
	"github.com/christoffer1009/cobweb-go/tree"
)

func main() {

	root := node.NewNode(0)
	cobweb := tree.NewTree(root)

	child1 := node.NewNode(1)
	child2 := node.NewNode(2)
	oc1 := ocurrence.NewOcurrence(1, "w", 1, 1)
	oc2 := ocurrence.NewOcurrence(2, "w", 2, 2)
	oc3 := ocurrence.NewOcurrence(3, "b", 2, 2)
	// oc4 := ocurrence.NewOcurrence(4, "b", 3, 1)

	cobweb.Root.AddChild(child1)

	child1.AddOcurrence(oc1)
	child1.AddOcurrence(oc3)
	child2.AddOcurrence(oc2)
	// child1.AddOcurrence(oc4)

	child1.PrintOcurrences()
	child1.CalcP()
	child2.CalcP()
	fmt.Println("FIM")

}
