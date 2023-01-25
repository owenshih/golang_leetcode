package golang

func isSymmetric(root *TreeNode) bool {
	return check(root.Left, root.Right)
}

func check(left *TreeNode, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	if left.Val != right.Val {
		return false
	}
	return check(left.Left, right.Right) && check(left.Right, right.Left)
}
