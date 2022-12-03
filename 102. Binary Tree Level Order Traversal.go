package golang

func levelOrder(root *TreeNode) [][]int {
	var array [][]int
	if root == nil {
		return array
	}
	var end = false
	var nextLevelNode []*TreeNode
	var loopLevelNode []*TreeNode
	nextLevelNode = append(nextLevelNode, root)
	for !end {
		var currentLevelInt []int
		loopLevelNode = nextLevelNode
		nextLevelNode = nil
		for _, v := range loopLevelNode {
			if v == nil {
				continue
			}
			currentLevelInt = append(currentLevelInt, v.Val)
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
		array = append(array, currentLevelInt)
	}
	return array
}
