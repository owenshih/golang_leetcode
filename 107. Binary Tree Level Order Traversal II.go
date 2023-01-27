package golang

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var nextLevelNodes, currentLevelNodes []*TreeNode
	var resultLevelVal [][]int
	var currentLevelVal []int
	nextLevelNodes = append(nextLevelNodes, root)
	for len(nextLevelNodes) != 0 {
		currentLevelNodes, nextLevelNodes = nextLevelNodes, nil
		for _, v := range currentLevelNodes {
			if v == nil {
				continue
			}
			currentLevelVal = append(currentLevelVal, v.Val)
			nextLevelNodes = append(nextLevelNodes, v.Left)
			nextLevelNodes = append(nextLevelNodes, v.Right)
		}
		if len(currentLevelVal) > 0 {
			resultLevelVal = append(resultLevelVal, currentLevelVal)
			currentLevelVal = nil
		}
	}
	newResultLevelVal := make([][]int, 0, len(resultLevelVal))
	for i := len(resultLevelVal) - 1; i >= 0; i-- {
		newResultLevelVal = append(newResultLevelVal, resultLevelVal[i])
	}
	return newResultLevelVal
}
