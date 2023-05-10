package node

import (
	"fmt"
	"math"
	"strings"

	"github.com/christoffer1009/cobweb-go/occurrence"
)

type Node struct {
	ID          string
	Occurrences []*occurrence.Occurrence
	Parent      *Node
	Children    []*Node
	P           float64
	TotalP      float64
}

func NewNode(ID string) *Node {
	return &Node{
		ID:          ID,
		Occurrences: []*occurrence.Occurrence{},
		Children:    []*Node{},
	}
}

func PrintNodes(node *Node, indent int) {
	// Imprimir o nó atual com indentação
	fmt.Printf("%sNode ID: %s\n", strings.Repeat(" ", indent), node.ID)

	// Chamar a função recursivamente para cada filho
	for _, child := range node.Children {
		PrintNodes(child, indent+2)
	}
}

func (n *Node) AddChild(child *Node) {
	child.Parent = n
	n.Children = append(n.Children, child)
}

func (n *Node) RemoveChild(child *Node) {
	for i, c := range n.Children {
		if c == child {
			n.Children = append(n.Children[:i], n.Children[i+1:]...)
			break
		}
	}
}

func (n *Node) AddOccurrence(oc *occurrence.Occurrence) {
	n.Occurrences = append(n.Occurrences, oc)
}

func calcPColor(n *Node, occurrences []*occurrence.Occurrence) map[string]float64 {
	countMap := make(map[string]int)
	for _, oc := range occurrences {
		_, exists := countMap[oc.Color]
		if exists {
			countMap[oc.Color]++
		} else {
			countMap[oc.Color] = 1
		}
	}
	countMap["quantity"] = len(occurrences)
	pColor := calcHelper(countMap)

	return pColor
}

func calcPNucleus(n *Node, occurrences []*occurrence.Occurrence) map[string]float64 {

	countMap := make(map[string]int)
	for _, oc := range occurrences {
		_, exists := countMap[fmt.Sprint(oc.Nucleus)]
		if exists {
			countMap[fmt.Sprint(oc.Nucleus)]++
		} else {
			countMap[fmt.Sprint(oc.Nucleus)] = 1
		}
	}

	countMap["quantity"] = len(occurrences)
	pNucleus := calcHelper(countMap)
	return pNucleus
}

func calcPTail(n *Node, occurrences []*occurrence.Occurrence) map[string]float64 {
	countMap := make(map[string]int)
	for _, oc := range occurrences {
		_, exists := countMap[fmt.Sprint(oc.Tail)]
		if exists {
			countMap[fmt.Sprint(oc.Tail)]++
		} else {
			countMap[fmt.Sprint(oc.Tail)] = 1
		}
	}

	countMap["quantity"] = len(occurrences)
	pTail := calcHelper(countMap)
	return pTail
}

func calcHelper(m map[string]int) map[string]float64 {
	result := make(map[string]float64)
	for k, v := range m {
		if k != "quantity" {
			result[k] = math.Pow(float64(v)/float64(m["quantity"]), 2)
		}
	}
	return result
}

func calcP(n *Node, occurrences []*occurrence.Occurrence) float64 {
	pColor := calcPColor(n, n.Occurrences)
	pNucleus := calcPNucleus(n, n.Occurrences)
	pTail := calcPTail(n, n.Occurrences)
	var p float64 = 0
	for _, v := range pColor {
		p += v
	}
	for _, v := range pNucleus {
		p += v
	}
	for _, v := range pTail {
		p += v
	}
	pTotal := calcPTotal(n, occurrences)
	n.P = (p - pTotal) * float64(len(n.Occurrences)) / float64(len(occurrences))
	return n.P
}

func calcPTotal(n *Node, ocs []*occurrence.Occurrence) float64 {
	pColor := calcPColor(n, ocs)
	pNucleus := calcPNucleus(n, ocs)
	pTail := calcPTail(n, ocs)
	var p float64 = 0
	for _, v := range pColor {
		p += v
	}
	for _, v := range pNucleus {
		p += v
	}
	for _, v := range pTail {
		p += v
	}

	n.TotalP = p
	return p
}

func SumP(n *Node, occurences []*occurrence.Occurrence) float64 {
	if n == nil {
		return 0
	}

	calcP(n, occurences)
	sum := n.P

	for _, child := range n.Children {
		sum += SumP(child, occurences)
	}

	return sum
}

func CopyNode(original *Node) *Node {
	if original == nil {
		return nil
	}

	// Cria uma cópia do objeto Node
	copy := &Node{
		ID:          original.ID,
		Occurrences: original.Occurrences,
		Parent:      original.Parent,
		P:           original.P,
		TotalP:      original.TotalP,
	}

	// Copia os nós filhos recursivamente
	for _, child := range original.Children {
		copy.Children = append(copy.Children, CopyNode(child))
	}

	return copy
}
