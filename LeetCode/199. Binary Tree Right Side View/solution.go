/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    maxlevel := 0
    return rightSideViewRecursive(root, 1, &maxlevel, []int{})

}

func rightSideViewRecursive(root *TreeNode, level int, maxlevel *int, values []int) []int {
    if root == nil {
        return values
    }
    if *maxlevel < level {
        values = append(values, root.Val)
        *maxlevel = level
    }

    level = level + 1
    values = rightSideViewRecursive(root.Right, level, maxlevel, values)
    values = rightSideViewRecursive(root.Left, level, maxlevel, values)
    return values
}