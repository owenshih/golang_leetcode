package golang

func inorderTraversal(root *TreeNode) []int {

	var array []int
	if root == nil {
		return array
	}
	arrayLeft := inorderTraversal(root.Left)
	for _, v := range arrayLeft {
		array = append(array, v)
	}
	array = append(array, root.Val)
	arrayRight := inorderTraversal(root.Right)
	for _, v := range arrayRight {
		array = append(array, v)
	}
	return array
}
