package tree

import (
	"fmt"
	"sort"

	"github.com/christoffer1009/cobweb-go/node"
	"github.com/christoffer1009/cobweb-go/occurrence"
)

type Tree struct {
	Root *node.Node
	// UC   float64
}

func NewTree(node *node.Node) *Tree {
	return &Tree{
		Root: node,
		// UC:   0,
	}
}

func countNodes(n *node.Node) int {
	if n == nil {
		return 0
	}

	count := 0

	for _, child := range n.Children {
		count += 1 + countNodes(child)
	}
	return count
}

func calcUC(n *node.Node, occurrences []*occurrence.Occurrence) float64 {
	nodes_count := countNodes(n)
	var uc float64 = 0
	if nodes_count > 0 {
		uc = node.SumP(n, occurrences) / float64(nodes_count)

	} else {
		uc = node.SumP(n, occurrences)
	}

	// t.UC = uc
	return uc
}

// func copyTree(original *Tree) *Tree {
// 	// Cria uma cópia do objeto Tree
// 	copy := &Tree{
// 		// UC:   original.UC,
// 		Root: nil,
// 	}

// 	// Realiza a cópia profunda do nó root
// 	copy.Root = node.CopyNode(original.Root)

// 	return copy
// }

func getChildrenUC(n *node.Node, occ *occurrence.Occurrence) []map[string]interface{} {
	var childrenUC []map[string]interface{}

	temp := node.CopyNode(n)

	for i := 0; i < len(temp.Children); i++ {
		temp.Children[i].AddOccurrence(occ)
		uc := calcUC(temp, n.Occurrences)
		item := map[string]interface{}{"index": i, "ID": temp.Children[i].ID, "UC": uc}
		childrenUC = append(childrenUC, item)
		temp = node.CopyNode(n)

		fmt.Printf("UC DO FILHO %s: %f\n", temp.Children[i].ID, uc)
	}

	sort.Slice(childrenUC, func(i, j int) bool {
		UC1 := childrenUC[i]["UC"].(float64)
		UC2 := childrenUC[j]["UC"].(float64)
		return UC1 > UC2
	})
	return childrenUC
}

func getNewChildUC(n *node.Node, occ *occurrence.Occurrence) float64 {
	temp := node.CopyNode(n)
	new := node.NewNode(fmt.Sprintf("%s.%d", n.ID, len(n.Children)))
	new.AddOccurrence(occ)
	temp.AddChild(new)
	newChildUC := calcUC(n, n.Occurrences)
	fmt.Printf("UC NOVO NÓ: %f \n", newChildUC)
	return newChildUC
}

func (tree *Tree) Cobweb(n *node.Node, occ *occurrence.Occurrence) {
	fmt.Printf("\nOCORRENCIA: %s - %d - %d \n", occ.Color, occ.Nucleus, occ.Tail)
	n.AddOccurrence(occ)

	if len(n.Children) == 0 {
		fmt.Println("FOLHA")
		new1 := node.NewNode(fmt.Sprintf("%s.%d", n.ID, len(n.Children)))
		new1.Occurrences = n.Occurrences
		n.AddChild(new1)
		new2 := node.NewNode(fmt.Sprintf("%s.%d", n.ID, len(n.Children)))
		new2.AddOccurrence(occ)
		n.AddChild(new2)

	} else {

		childrenUC := getChildrenUC(n, occ)
		newChildUC := getNewChildUC(n, occ)
		bestChildIndex := childrenUC[0]["index"].(int)

		if childrenUC[0]["UC"].(float64) > newChildUC {
			fmt.Printf("ADICIONA EM MELHOR FILHO %s \n", childrenUC[0]["ID"])
			tree.Cobweb(n.Children[bestChildIndex], occ)
			//n.Children[bestChildIndex].AddOccurrence(occ)
			//tree.calcUC(n, tree.Root.Occurrences)

		} else {
			new := node.NewNode(fmt.Sprintf("%s.%d", n.ID, len(n.Children)))
			new.AddOccurrence(occ)
			n.AddChild(new)
			// tree.calcUC(n, tree.Root.Occurrences)
			fmt.Printf("CRIAR NOVO NÓ %s \n", new.ID)
		}

	}

	// node.PrintNodes(n, 1)

	// fmt.Printf("\nNó : %d , FILHOS:\n", n.ID)
	// for _, child := range n.Children {
	// 	fmt.Printf("ID:%d OCC:%d\n", child.ID, len(child.Occurrences))
	// 	child.PrintOccurrences()
	// }

}
