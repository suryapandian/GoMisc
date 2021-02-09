/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }

    if !isMirror(root.Left, root.Right) {
        return false
    }
    return true

}

func isMirror(p *TreeNode, q *TreeNode) bool {
    if p == nil {
        if q == nil {
            return true
        }
        return false
    }

    if q == nil {
        return false
    }

    if p.Val != q.Val {
        return false
    }

    if !isMirror(p.Left, q.Right) {
        return false
    }
    if !isMirror(p.Right, q.Left) {
        return false
    }
    return true
}