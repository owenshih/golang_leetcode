package golang

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return dfsSum(root, targetSum)
}

func dfsSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}
	if dfsSum(root.Left, targetSum) || dfsSum(root.Right, targetSum) {
		return true
	} else {
		return false
	}
}
