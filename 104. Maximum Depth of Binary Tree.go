package golang

func maxDepth(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, count int) int {
	if root == nil {
		return count
	}

	count++

	leftCount := dfs(root.Left, count)
	rightCount := dfs(root.Right, count)
	if leftCount > rightCount {
		return leftCount
	} else {
		return rightCount
	}
}
