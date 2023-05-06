package node

import (
	"fmt"
	"math"

	"github.com/christoffer1009/cobweb-go/ocurrence"
)

type Node struct {
	ID          int
	Occurrences []*ocurrence.Occurrence
	Children    []*Node
	P           float64
	TotalP      float64
}

func NewNode(ID int) *Node {
	return &Node{
		ID:          ID,
		Occurrences: []*ocurrence.Occurrence{},
		Children:    []*Node{},
	}
}

func (n *Node) PrintID() {
	fmt.Println(n.ID)
	for _, child := range n.Children {
		child.PrintID()
	}
}

func (n *Node) PrintOcurrences() {
	for _, oc := range n.Occurrences {
		str := fmt.Sprintf("ID %d - Cor: %s - Núcleos: %d - Caudas: %d", oc.ID, oc.Color, oc.Nucleus, oc.Tail)
		fmt.Println(str)
	}
}

func (n *Node) AddChild(child *Node) {
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

func (n *Node) AddOcurrence(oc *ocurrence.Occurrence) {
	n.Occurrences = append(n.Occurrences, oc)
}

func calcPColor(n *Node) map[string]float64 {
	countMap := make(map[string]int)
	for _, oc := range n.Occurrences {
		_, exists := countMap[oc.Color]
		if exists {
			countMap[oc.Color]++
		} else {
			countMap[oc.Color] = 1
		}
	}
	countMap["quantity"] = len(n.Occurrences)
	pColor := CalcHelper(countMap)

	return pColor
}

func calcPNucleus(n *Node) map[string]float64 {

	countMap := make(map[string]int)
	for _, oc := range n.Occurrences {
		_, exists := countMap[fmt.Sprint(oc.Nucleus)]
		if exists {
			countMap[fmt.Sprint(oc.Nucleus)]++
		} else {
			countMap[fmt.Sprint(oc.Nucleus)] = 1
		}
	}

	countMap["quantity"] = len(n.Occurrences)
	pNucleus := CalcHelper(countMap)
	return pNucleus
}

func calcPTail(n *Node) map[string]float64 {
	countMap := make(map[string]int)
	for _, oc := range n.Occurrences {
		_, exists := countMap[fmt.Sprint(oc.Tail)]
		if exists {
			countMap[fmt.Sprint(oc.Tail)]++
		} else {
			countMap[fmt.Sprint(oc.Tail)] = 1
		}
	}

	countMap["quantity"] = len(n.Occurrences)
	pTail := CalcHelper(countMap)
	return pTail
}

func CalcHelper(m map[string]int) map[string]float64 {
	result := make(map[string]float64)
	for k, v := range m {
		if k != "quantity" {
			result[k] = math.Pow(float64(v)/float64(m["quantity"]), 2)
		}
	}
	return result
}

func (n *Node) CalcP() float64 {
	pColor := calcPColor(n)
	pNucleus := calcPNucleus(n)
	pTail := calcPTail(n)
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

	n.P = p
	return p
}
