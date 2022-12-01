package golang

func postorderTraversal(root *TreeNode) []int {
	var array []int
	if root == nil {
		return array
	}
	arrayLeft := postorderTraversal(root.Left)
	for _, v := range arrayLeft {
		array = append(array, v)
	}
	arrayRight := postorderTraversal(root.Right)
	for _, v := range arrayRight {
		array = append(array, v)
	}
	array = append(array, root.Val)
	return array
}
