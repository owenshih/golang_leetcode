package golang

func averageOfLevels(root *TreeNode) []float64 {
	var array []float64
	if root == nil {
		return array
	}
	var end = false
	var nextLevelNode []*TreeNode
	var loopLevelNode []*TreeNode
	nextLevelNode = append(nextLevelNode, root)
	for !end {
		var currentLevelInt float64
		var currentLevelIntCount float64
		loopLevelNode = nextLevelNode
		nextLevelNode = nil
		for _, v := range loopLevelNode {
			if v == nil {
				continue
			}
			currentLevelInt += float64(v.Val)
			currentLevelIntCount++
			if v.Left != nil {
				nextLevelNode = append(nextLevelNode, v.Left)
			}
			if v.Right != nil {
				nextLevelNode = append(nextLevelNode, v.Right)
			}
		}
		if len(loopLevelNode) == 0 {
			end = true
			return array
		}
		array = append(array, currentLevelInt/currentLevelIntCount)
	}
	return array
}
