package tree

import (
	"fmt"
	"sort"

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

func CopyTree(original *Tree) *Tree {
	// Cria uma cópia do objeto Tree
	copy := &Tree{
		UC:   original.UC,
		Root: nil,
	}

	// Realiza a cópia profunda do nó root
	copy.Root = node.CopyNode(original.Root)

	return copy
}

func (t *Tree) GetChildrenUC(occ *occurrence.Occurrence) []map[string]interface{} {
	var childrenUC []map[string]interface{}

	temp := CopyTree(t)

	for i := 0; i < len(temp.Root.Children); i++ {
		temp.Root.Children[i].AddOccurrence(occ)
		uc := temp.CalcUC(temp.Root.Occurrences)
		item := map[string]interface{}{"index": i, "ID": temp.Root.Children[i].ID, "UC": uc}
		childrenUC = append(childrenUC, item)
		temp = CopyTree(t)

		fmt.Printf("UC DO FILHO %d: %f\n", temp.Root.Children[i].ID, uc)
	}

	sort.Slice(childrenUC, func(i, j int) bool {
		UC1 := childrenUC[i]["UC"].(float64)
		UC2 := childrenUC[j]["UC"].(float64)
		return UC1 > UC2
	})
	return childrenUC
}

func (t *Tree) GetNewChildUC(occ *occurrence.Occurrence) float64 {
	temp := CopyTree(t)
	new := node.NewNode(len(t.Root.Children))
	new.AddOccurrence(occ)
	temp.Root.AddChild(new)
	newChildUC := temp.CalcUC(temp.Root.Occurrences)

	return newChildUC
}

func (tree *Tree) Cobweb(n *node.Node, occ *occurrence.Occurrence) {
	tree.Root.AddOccurrence(occ)
	fmt.Printf("OCORRENCIA: %s - %d - %d \n", occ.Color, occ.Nucleus, occ.Tail)

	if len(n.Children) == 0 {
		new := node.NewNode(n.ID + 1)
		new.AddOccurrence(occ)
		n.AddChild(new)
		tree.CalcUC(tree.Root.Occurrences)
		fmt.Println("FOLHA")

	} else {

		childrenUC := tree.GetChildrenUC(occ)
		newChildUC := tree.GetNewChildUC(occ)
		bestChildIndex := childrenUC[0]["index"].(int)

		if childrenUC[0]["UC"].(float64) > newChildUC {
			fmt.Printf("ADICIONA EM MELHOR FILHO %d \n", childrenUC[0]["ID"])
			n.Children[bestChildIndex].AddOccurrence(occ)
			tree.CalcUC(tree.Root.Occurrences)
		} else {
			fmt.Printf("CRIAR NOVO NÓ %f \n", newChildUC)
			new := node.NewNode(len(n.Children) + 1)
			new.AddOccurrence(occ)
			n.AddChild(new)
			tree.CalcUC(tree.Root.Occurrences)

		}

	}

}
