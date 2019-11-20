package main

import (
	"fmt"
	"math"
)

type Node struct {
	Value float64 `json:"value"`
	Next  *Node   `json:"next"`
}

func (n *Node) Add(otherNode *Node) float64 {
	return n.NodeValue(0) + otherNode.NodeValue(0)
}

func (n *Node) NodeValue(pos float64) float64 {
	nodeValue := n.Value * math.Pow(10, pos)
	if n.hasNextNode() {
		return nodeValue + n.Next.NodeValue(pos+1)
	}
	return nodeValue
}

func (n *Node) hasNextNode() bool {
	return n.Next != nil
}

func main() {
	nodeOne := &Node{2, &Node{4, &Node{3, nil}}}
	nodeTwo := &Node{5, &Node{6, &Node{4, nil}}}

	fmt.Println(nodeOne.Add(nodeTwo))
}