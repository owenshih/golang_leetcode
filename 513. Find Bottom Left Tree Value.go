package golang

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	var nextLevelNodes, currentLevelNodes []*TreeNode
	var result int
	findLeftest := false
	nextLevelNodes = append(nextLevelNodes, root)
	for len(nextLevelNodes) != 0 {
		currentLevelNodes, nextLevelNodes = nextLevelNodes, nil
		for _, v := range currentLevelNodes {
			if v == nil {
				continue
			}
			if !findLeftest {
				result = v.Val
				findLeftest = true
			}
			nextLevelNodes = append(nextLevelNodes, v.Left)
			nextLevelNodes = append(nextLevelNodes, v.Right)
		}
		findLeftest = false
	}
	return result
}
