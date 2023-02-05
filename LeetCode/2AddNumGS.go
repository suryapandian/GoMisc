package main

import (
	"math"
)

type Node struct {
	Value int `json:"value"`
	Next  *Node   `json:"next"`
}

func (n *Node) Add(otherNode *Node) int {
	return n.NodeValue(0) + otherNode.NodeValue(0)
}

func (n *Node) NodeValue(pos int) int {
	nodeValue := n.Value * int(math.Pow(10, float64(pos)))
	if n.hasNextNode() {
		return nodeValue + n.Next.NodeValue(pos+1)
	}
	return nodeValue
}

func (n *Node) hasNextNode() bool {
	return n.Next != nil
}

func createNode(num int) *Node{
	node := &Node{num%10, nil}
	for num/10 != 0{
		num = num/10
		node = &Node{num%10, node}
	}
	return node
}

func main() {
	nodeOne := &Node{2, &Node{4, &Node{3, nil}}}
	nodeTwo := &Node{5, &Node{6, &Node{4, nil}}}
	ans := nodeOne.Add(nodeTwo)
	createNode(ans)
}