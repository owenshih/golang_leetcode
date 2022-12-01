package golang

func preorderTraversal(root *TreeNode) []int {
	var array []int
	if root == nil {
		return array
	}
	array = append(array, root.Val)
	arrayLeft := preorderTraversal(root.Left)
	for _, v := range arrayLeft {
		array = append(array, v)
	}
	arrayRight := preorderTraversal(root.Right)
	for _, v := range arrayRight {
		array = append(array, v)
	}
	return array
}
