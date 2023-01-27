package golang

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	var nextLevelNodes, currentLevelNodes []*TreeNode
	var result []int
	nextLevelNodes = append(nextLevelNodes, root)
	for len(nextLevelNodes) != 0 {
		currentLevelNodes, nextLevelNodes = nextLevelNodes, nil
		for i, v := range currentLevelNodes {
			if v == nil {
				continue
			}
			if i == len(currentLevelNodes)-1 {
				result = append(result, v.Val)
			}
			if v.Left != nil {
				nextLevelNodes = append(nextLevelNodes, v.Left)
			}
			if v.Right != nil {
				nextLevelNodes = append(nextLevelNodes, v.Right)
			}
		}
	}
	return result
}
