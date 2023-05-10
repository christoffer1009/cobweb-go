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

func (t *Tree) PrintTree() {
	if t.Root != nil {
		t.Root.PrintID()
	}
}

func CountNodes(n *node.Node) int {
	if n == nil {
		return 0
	}

	count := 0

	for _, child := range n.Children {
		count += 1 + CountNodes(child)
	}
	return count
}

func CalcUC(n *node.Node, occurrences []*occurrence.Occurrence) float64 {
	nodes_count := CountNodes(n)
	var uc float64 = 0
	if nodes_count > 0 {
		uc = node.SumP(n, occurrences) / float64(nodes_count)

	} else {
		uc = node.SumP(n, occurrences)
	}

	// t.UC = uc
	return uc
}

func CopyTree(original *Tree) *Tree {
	// Cria uma cópia do objeto Tree
	copy := &Tree{
		// UC:   original.UC,
		Root: nil,
	}

	// Realiza a cópia profunda do nó root
	copy.Root = node.CopyNode(original.Root)

	return copy
}

func (t *Tree) GetChildrenUC(n *node.Node, occ *occurrence.Occurrence) []map[string]interface{} {
	var childrenUC []map[string]interface{}

	temp := node.CopyNode(n)

	for i := 0; i < len(temp.Children); i++ {
		temp.Children[i].AddOccurrence(occ)
		uc := CalcUC(n, t.Root.Occurrences)
		item := map[string]interface{}{"index": i, "ID": temp.Children[i].ID, "UC": uc}
		childrenUC = append(childrenUC, item)
		temp = node.CopyNode(n)

		fmt.Printf("UC DO FILHO %d: %f\n", temp.Children[i].ID, uc)
	}

	sort.Slice(childrenUC, func(i, j int) bool {
		UC1 := childrenUC[i]["UC"].(float64)
		UC2 := childrenUC[j]["UC"].(float64)
		return UC1 > UC2
	})
	return childrenUC
}

func (t *Tree) GetNewChildUC(n *node.Node, occ *occurrence.Occurrence) float64 {
	temp := node.CopyNode(n)
  new := node.NewNode(len(n.Children))
	new.AddOccurrence(occ)
	temp.AddChild(new)
	newChildUC := CalcUC(n, t.Root.Occurrences)
	fmt.Printf("UC NOVO NÓ: %f \n", newChildUC)
	return newChildUC
}

func (tree *Tree) Cobweb(n *node.Node, occ *occurrence.Occurrence) {
	tree.Root.AddOccurrence(occ)
	fmt.Printf("\nOCORRENCIA: %s - %d - %d \n", occ.Color, occ.Nucleus, occ.Tail)

	if len(n.Children) == 0 {
    n.AddOccurrence(occ)
		new := node.NewNode(n.ID + 1)
		new.AddOccurrence(occ)
		n.AddChild(new)
		//tree.CalcUC(tree.Root.Occurrences)
		fmt.Println("FOLHA")

	} else {

		childrenUC := tree.GetChildrenUC(n, occ)
		newChildUC := tree.GetNewChildUC(n, occ)
		bestChildIndex := childrenUC[0]["index"].(int)

		if childrenUC[0]["UC"].(float64) > newChildUC {
			fmt.Printf("ADICIONA EM MELHOR FILHO %d \n", childrenUC[0]["ID"])
			tree.Cobweb(n.Children[bestChildIndex], occ)
      //n.Children[bestChildIndex].AddOccurrence(occ)
			//tree.CalcUC(n, tree.Root.Occurrences)
      
		} else {
			new := node.NewNode(len(n.Children) + 1)
			new.AddOccurrence(occ)
			n.AddChild(new)
			// tree.CalcUC(n, tree.Root.Occurrences)
			fmt.Printf("CRIAR NOVO NÓ %d \n", new.ID)
		}

	}

	fmt.Printf("\nNó : %d , FILHOS:\n", n.ID)
	for _, child := range n.Children {
		fmt.Printf("ID:%d OCC:%d\n", child.ID, len(child.Occurrences))
		child.PrintOccurrences()
	}

}
