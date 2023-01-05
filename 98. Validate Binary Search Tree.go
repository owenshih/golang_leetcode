package golang

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return check(root, nil, nil)
}

func check(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if (min != nil && root.Val <= min.Val) ||
		(max != nil && root.Val >= max.Val) {
		return false
	}
	return check(root.Left, min, root) && check(root.Right, root, max)
}
