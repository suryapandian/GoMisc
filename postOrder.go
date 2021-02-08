package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

/* Example tree from https://en.wikipedia.org/wiki/Binary_search_tree
        8
      /   \
    3       10
  /   \        \
1       6        14
      /   \    /
    4      7  13
    https://us02web.zoom.us/j/87688071120?pwd=ejlDQWJINTV3T0R1L1lyYUF2RlpIdz09
*/

func PostOrder(n *Node) (r []int) {
	if n == nil {
		return
	}
	if n.Left != nil {
		r = append(r, PostOrder(n.Left)...)
	}
	if n.Right != nil {
		r = append(r, PostOrder(n.Right)...)
	}

	r = append(r, n.Value)

	return
}

func isValid(r []int) (b bool) {
	// [3,14,10,8]
	n := len(r)
	if n == 0 {
		return
	}
	//var arr []int

	root := r[n-1]

	for i := 0; i < (n - 1); i++ {
		if r[i] > root {

		}
		if r[i] < root {
			return
		}

	}

	// mid := (n - 1) / 2
	// for i := 0; i < mid; i++ {
	// 	fmt.Println("left", r[i])
	// 	if r[i] > r[n-1] {
	// 		return false
	// 	}
	// }

	// for i := mid + 1; i < n; i++ {
	// 	fmt.Println("Right", r[i])
	// 	if r[i] < r[n-1] {
	// 		return false
	// 	}

	// }
	return true
	// for _, v := range r {

	// }
}

func main() {
	input := []int{3, 10, 8}
	var tree = &Node{
		Value: 8,
		Left: &Node{
			Value: 3,
			Left:  &Node{Value: 1},
			Right: &Node{
				Value: 6,
				Left:  &Node{Value: 4},
				Right: &Node{Value: 7},
			},
		},
		Right: &Node{
			Value: 10,
			Right: &Node{
				Value: 14,
				Left:  &Node{Value: 13},
			},
		},
	}
	fmt.Println("array", PostOrder(tree))
	fmt.Println("isValid", isValid(input))
	fmt.Println("big array", isValid(PostOrder(tree)))
	//fmt.Println("invalid tree", isValid([]int{}))

}
